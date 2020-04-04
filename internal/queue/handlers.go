package queue

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type netHandler int

const (
	CLOSED = netHandler(iota)
	RECONNECT
	DISCOVERED
)

func (n netHandler) String() string {
	return []string{"CLOSED", "RECONNECT", "DISCOVERED"}[n]
}

func (q *NATSQueue) connHandler(t netHandler) nats.ConnHandler {
	return func(conn *nats.Conn) {
		id, _ := conn.GetClientID()
		q.log.Info("conn handler",
			zap.String("type", t.String()),
			zap.Uint64("clientID", id),
			zap.String("connectedURL", conn.ConnectedUrl()),
			zap.String("serverID", conn.ConnectedServerId()),
			zap.Strings("discoveredServers", conn.DiscoveredServers()),
			zap.String("status", getStatusString(int(conn.Status()))),
		)
	}
}

func (q *NATSQueue) disconnectErrHandler(conn *nats.Conn, err error) {
	id, _ := conn.GetClientID()
	q.log.Info("disconnecting error",
		zap.Error(err),
		zap.Uint64("clientID", id),
		zap.String("connectedURL", conn.ConnectedUrl()),
		zap.String("serverID", conn.ConnectedServerId()),
		zap.Strings("discoveredServers", conn.DiscoveredServers()),
		zap.String("status", getStatusString(int(conn.Status()))),
	)
}

func (q *NATSQueue) errorHandler(conn *nats.Conn, sub *nats.Subscription, err error) {
	id, _ := conn.GetClientID()
	q.log.Info("nats received error",
		zap.Error(err),
		zap.String("subject", sub.Subject),
		zap.String("queue", sub.Queue),
		zap.Uint64("clientID", id),
		zap.String("connectedURL", conn.ConnectedUrl()),
		zap.String("serverID", conn.ConnectedServerId()),
		zap.Strings("discoveredServers", conn.DiscoveredServers()),
		zap.String("status", getStatusString(int(conn.Status()))),
	)
}

// getStatusString returns the string version of nats.Status
func getStatusString(s int) string {
	return []string{"DISCONNECTED", "CONNECTED", "CLOSED", "RECONNECTING", "CONNECTING", "DRAINING_SUBS", "DRAINING_PUBS"}[s]
}
