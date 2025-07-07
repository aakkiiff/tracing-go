package routes

import (
	"context"
	"fmt"
	"net/http"

	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

func hello(c *gin.Context) {
	tracer := otel.Tracer("hellofunc")
	ctx, span := tracer.Start(c.Request.Context(), "HTTP GET /hello")
	defer span.End()

	step1(ctx)
	step2(ctx)

	c.JSON(http.StatusOK, gin.H{"message": "hello world", "status": "200"})
}

func step1(ctx context.Context) {
	tracer := otel.Tracer("step1 func")
	_, span := tracer.Start(ctx, "step1")
	defer span.End()

	fmt.Println("sleeping 1 sec")
	time.Sleep(300 * time.Millisecond)
}

func step2(ctx context.Context) {
	tracer := otel.Tracer("step2 func")
	_, span := tracer.Start(ctx, "step2")
	defer span.End()

	step2pre(ctx)

	// Generate random sleep duration between 1 and 9 seconds
	rand.Seed(time.Now().UnixNano())
	sleepDuration := time.Duration(rand.Intn(9)+1) * time.Second

	fmt.Printf("sleeping %v\n", sleepDuration)
	time.Sleep(sleepDuration)

	step2a(ctx)
}

func step2a(ctx context.Context) {
	tracer := otel.Tracer("step2a func")
	_, span := tracer.Start(ctx, "step2a")
	defer span.End()

	fmt.Println("sleeping 1 sec")
	time.Sleep(100 * time.Millisecond)
}

func step2pre(ctx context.Context) {
	tracer := otel.Tracer("step2pre func")
	_, span := tracer.Start(ctx, "step2pre")
	defer span.End()

	fmt.Println("sleeping 1 sec")
	time.Sleep(100 * time.Millisecond)
}
