package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryeoman/tour/internal/infra/http/handler"
)

// HTTPServer represents the HTTP server for the application.
type HTTPServer struct {
	router          *gin.Engine
	tourPlanHandler handler.TourPlanHandler
}

// NewHTTPServer creates a new HTTPServer instance.
func NewHTTPServer(tourPlanHandler handler.TourPlanHandler) *HTTPServer {
	return &HTTPServer{
		tourPlanHandler: tourPlanHandler,
	}
}

// Start starts the HTTP server.
func (s *HTTPServer) Start(addr string) error {
	s.router = gin.Default()
	s.router.GET("/tour-plans", s.tourPlanHandler.Get)
	s.router.GET("/tour-plans/:id", s.tourPlanHandler.GetByID)
	s.router.POST("/tour-plans", s.tourPlanHandler.Create)

	srv := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Println("Running http server at port: ", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Failed to run http server: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return nil
}
