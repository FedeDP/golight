package signals

import (
	"os"
	"os/signal"
	"syscall"
)

func Subscribe() <-chan os.Signal {
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	return sigC
}
