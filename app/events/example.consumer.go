package eventconsumer

import (
	"fmt"
	"framework/internal/app/logger"
)

func ExampleConsumer(payload string, l logger.Logger) {
	l.Info(fmt.Sprintf("Event %s consumed", payload))
}
