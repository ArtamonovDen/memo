package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Start server...")

	// Graceful shutdown based on context instead of explicit signal handling
	// https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Set up router
	router := gin.Default()

	router.GET("/api/account", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "folks",
		})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server listening error %s\n", err)
		}
	}()

	// Handling interrupt signal
	<-ctx.Done()

	stop()
	log.Println("Waiting server to shut down gracefully, press Ctrl+C again to force")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Ask server to shutdown with timeout
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error when server shutdown: ", err)
	}
	log.Println("Exiting...")

}
