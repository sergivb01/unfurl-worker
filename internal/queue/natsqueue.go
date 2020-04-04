package queue

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"

	"github.com/sergivb01/unfurl-worker/internal/encoder"
	"github.com/sergivb01/unfurl-worker/internal/httpclient"
	"github.com/sergivb01/unfurl-worker/internal/metadata"
)

type NATSQueue struct {
	nc      *nats.EncodedConn
	ncMutex sync.Mutex

	pool *ants.PoolWithFunc

	log *zap.Logger

	Topic string
}

// TODO: use config for queue
func NewQueue(debug bool) (*NATSQueue, error) {
	logConfig := zap.NewDevelopmentConfig()
	if !debug {
		logConfig = zap.NewProductionConfig()
	}
	logConfig.DisableStacktrace = true
	// logConfig.DisableCaller = true

	logger, err := logConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("error creating logger: %w", err)
	}

	q := &NATSQueue{
		Topic: "requests",
		log:   logger,
	}

	rawConn, err := nats.Connect(nats.DefaultURL,
		nats.ClosedHandler(q.connHandler(CLOSED)),
		nats.ReconnectHandler(q.connHandler(RECONNECT)),
		nats.DiscoveredServersHandler(q.connHandler(DISCOVERED)),
		nats.DisconnectErrHandler(q.disconnectErrHandler),
		nats.ErrorHandler(q.errorHandler),
	)
	if err != nil {
		return nil, fmt.Errorf("error connecting to nats: %w", err)
	}

	nats.RegisterEncoder("msgpack", &encoder.MsgPackEncoder{})
	logger.Info("registered msgpack encoder to nats")

	encodedConn, err := nats.NewEncodedConn(rawConn, "msgpack")
	if err != nil {
		return nil, fmt.Errorf("failed to create encoded connection: %w", err)
	}

	q.nc = encodedConn

	return q, nil
}

func (q *NATSQueue) Subscribe() error {
	pool, err := ants.NewPoolWithFunc(15, func(i interface{}) {
		m, ok := i.(*nats.Msg)
		if !ok {
			return
		}
		q.handleMessage(m)
	}, ants.WithPreAlloc(true))
	if err != nil {
		return fmt.Errorf("failed to create ants pool: %q", err)
	}

	q.pool = pool

	_, err = q.nc.QueueSubscribe(q.Topic, "group", func(msg *nats.Msg) {
		if err := q.pool.Invoke(msg); err != nil {
			q.log.Error("error invoking ants pool", zap.Error(err))
		}
	})
	return err
}

func (q *NATSQueue) Start() {
	if err := q.nc.Flush(); err != nil {
		q.log.Error("error flushing", zap.Error(err))
	}

	if err := q.nc.LastError(); err != nil {
		q.log.Error("lastError is not nil!", zap.Error(err))
		return
	}

	q.log.Info("started nats connection", zap.String("topic", q.Topic))

	// Setup the interrupt handler to draining, so we don't miss
	// requests when scaling down.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	q.log.Info("draining connection")
	if err := q.nc.Drain(); err != nil {
		q.log.Error("error draining", zap.Error(err))
	}
}

func (q *NATSQueue) Queue(url string) (*metadata.PageInfo, error) {
	defer q.Track("Queue("+url+")", time.Now())

	q.ncMutex.Lock()
	nc := q.nc
	q.ncMutex.Unlock()

	var info metadata.PageInfo
	if err := nc.Request(q.Topic, &Request{URL: url}, &info, time.Second*3); err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	return &info, nil
}

func (q *NATSQueue) handleMessage(m *nats.Msg) {
	defer q.Track("handleMessage("+m.Reply+")", time.Now())

	var r Request
	if err := q.nc.Enc.Decode(m.Subject, m.Data, &r); err != nil {
		q.log.Error("error decoding message", zap.Error(err))
		return
	}

	// TODO: use a proper context
	reader, err := httpclient.GetReaderFromURL(context.TODO(), r.URL, true)
	if err != nil {
		q.log.Error("error getting reader from URL", zap.Error(err), zap.String("url", r.URL))
		return
	}

	info, err := metadata.ExtractInfoFromReader(reader)
	if err != nil {
		q.log.Error("error getting metadata from reader", zap.Error(err), zap.String("url", r.URL))
		return
	}

	if err := q.nc.Publish(m.Reply, info); err != nil {
		q.log.Error("error replying to queue", zap.Error(err), zap.String("url", r.URL))
	}
}
