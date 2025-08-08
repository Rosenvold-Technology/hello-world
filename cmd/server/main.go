package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/rosenvold-technologies/hello-world/internal/config"
	"github.com/rosenvold-technologies/hello-world/internal/middleware"
	"github.com/rosenvold-technologies/hello-world/internal/router"
)

// initTracer wires OpenTelemetry to a simple stdout exporter.
// The returned shutdown func flushes everything on exit.
func initTracer() func() {
	exp, _ := stdouttrace.New(stdouttrace.WithPrettyPrint())

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			// Service name:
			"hello.website",
			// Extra resource attributes:
			attribute.String("env", os.Getenv("ENVIRONMENT")),
		)),
	)

	otel.SetTracerProvider(tp)
	return func() { _ = tp.Shutdown(context.Background()) }
}

func main() {
	// ─────────────────── App config ───────────────────
	cfg := config.Load() // loads PORT, GIN_MODE, etc.

	// ─────────────────── Tracing ──────────────────────
	shutdown := initTracer()
	defer shutdown()

	// ─────────────────── Gin bootstrap ────────────────
	gin.SetMode(cfg.GinMode)

	// router.New() builds the engine, registers routes
	// (/, /healthz, static, templates), and adds Gin’s
	// own logger+recovery middleware.
	r := router.New()

	// Your structured-log middleware (adds request-ID, etc.)
	r.Use(middleware.Logger())

	// Expose Prometheus metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// ─────────────────── HTTP server ──────────────────
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MiB
	}

	// Start in a goroutine so we can wait for shutdown signals
	go func() {
		log.Printf("🚀 Starting server on port %d", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// ─────────────────── Graceful shutdown ────────────
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("💡 Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}

	log.Println("✅ Server exited cleanly")
}
