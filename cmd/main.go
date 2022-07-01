package main

import (
	"context"
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/delivery"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config := config.NewConfig()
	handler := http.NewServeMux()
	delivery.RegisterHTTPEndPoints(handler, config)
	server := &http.Server{
		Addr:    ":3030",
		Handler: handler,
	}
	errChan := make(chan error, 1)
	fmt.Printf("Starting server at port :3030\nhttp://localhost:3030")
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
			return
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-errChan:
		log.Println("error during server work")
	case <-c:
		log.Println("shutting down server due to: cancellatiion signal")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println(err.Error())
		return
	}
	// context,gorutines, channels izuchi
	// cancel()
}
