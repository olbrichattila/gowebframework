// Package boostrap is called when the application starts, regardless of that it is cli or http
package Bootstrap

import (
	eventconsumer "framework/app/events"
	"framework/internal/app/event"
	"framework/internal/app/logger"
)

// Bootstrap called when application starts, handles autowire dependencies in the function signature
func Bootstrap(l logger.Logger, e event.Eventer) {
	// Example event subscriber, consumer
	e.Subscribe("topic", "e4", eventconsumer.ExampleConsumer)
	e.Dispatch("topic", "event1")
	l.Info("bootstrap.called")
}
