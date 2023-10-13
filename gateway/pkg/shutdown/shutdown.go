package shutdown

import (
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/maxim12233/crypto-app-server/gateway/pkg/logger"
)

func Graceful(signals []os.Signal, closeItems ...io.Closer) {
	logger := logger.GetLogger()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, signals...)
	sig := <-sigc
	logger.Info(fmt.Sprintf("Caught signal %s. Shutting down...", sig))

	// Here we can do graceful shutdown (close connections and etc)
	for _, closer := range closeItems {
		if err := closer.Close(); err != nil {
			logger.Error(fmt.Sprintf("failed to close %v: %v", closer, err))
		}
	}
}
