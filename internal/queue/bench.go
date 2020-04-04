package queue

import (
	"time"

	"go.uber.org/zap"
)

func (q *NATSQueue) Track(name string, start time.Time) {
	q.log.Debug("timer finished", zap.String("name", name), zap.Duration("took", time.Since(start)))
}
