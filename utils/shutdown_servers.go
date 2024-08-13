package utils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ShutdownServers(httpShutdownChan chan struct{}) {
	// Wait for an interrupt signal in the main function
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Received shutdown signal, shutting down servers...")

	// Signal both servers to shut down in separate goroutines
	//go close(grpcShutdownChan)
	go close(httpShutdownChan)

	// Channels to confirm servers have shut down
	//grpcStopped := make(chan struct{})
	httpStopped := make(chan struct{})

	//go func() {
	//	<-grpcShutdownChan // Wait until gRPC server confirms it's stopped
	//	close(grpcStopped)
	//}()

	go func() {
		<-httpShutdownChan // Wait until HTTP server confirms it's stopped
		close(httpStopped)
	}()

	// Block until both servers have confirmed they've shut down
	//<-grpcStopped
	<-httpStopped

	//log.Println("Both servers have been shut down.")
}
