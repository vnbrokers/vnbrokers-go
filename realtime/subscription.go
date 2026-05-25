package realtime

import (
	"sync"
)

type ConnectionStatus string

const (
	StatusIdle           ConnectionStatus = "idle"
	StatusConnecting     ConnectionStatus = "connecting"
	StatusConnected      ConnectionStatus = "connected"
	StatusAuthenticating ConnectionStatus = "authenticating"
	StatusSubscribed     ConnectionStatus = "subscribed"
	StatusReconnecting   ConnectionStatus = "reconnecting"
	StatusClosing        ConnectionStatus = "closing"
	StatusClosed         ConnectionStatus = "closed"
	StatusFailed         ConnectionStatus = "failed"
)

type Subscription[T any] interface {
	Events() <-chan T
	Errors() <-chan error
	Status() <-chan ConnectionStatus
	Close() error
}

type QueueSubscription[T any] struct {
	events chan T
	errors chan error
	status chan ConnectionStatus
	close  func() error
	once   sync.Once
}

func NewQueueSubscription[T any](buffer int, closeFn func() error) *QueueSubscription[T] {
	if buffer <= 0 {
		buffer = 64
	}
	return &QueueSubscription[T]{
		events: make(chan T, buffer),
		errors: make(chan error, buffer),
		status: make(chan ConnectionStatus, buffer),
		close:  closeFn,
	}
}

func (s *QueueSubscription[T]) Events() <-chan T {
	return s.events
}

func (s *QueueSubscription[T]) Errors() <-chan error {
	return s.errors
}

func (s *QueueSubscription[T]) Status() <-chan ConnectionStatus {
	return s.status
}

func (s *QueueSubscription[T]) Close() error {
	var err error
	s.once.Do(func() {
		s.PublishStatus(StatusClosed)
		if s.close != nil {
			err = s.close()
		}
		close(s.events)
		close(s.errors)
		close(s.status)
	})
	return err
}

func (s *QueueSubscription[T]) PublishEvent(event T) {
	s.events <- event
}

func (s *QueueSubscription[T]) PublishError(err error) {
	s.errors <- err
}

func (s *QueueSubscription[T]) PublishStatus(status ConnectionStatus) {
	s.status <- status
}
