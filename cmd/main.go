package main

import (
	"context"
	"go-commerce-api/infrastructure/config"
	"go-commerce-api/infrastructure/database"
	"go-commerce-api/pkg/middleware"
	ur "go-commerce-api/internal/user/router"
	pdr "go-commerce-api/internal/product/router"
	pyr "go-commerce-api/internal/payment/router"

	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	user := e.Group("/users")
	ur.UserRouter(user, db)

	product := e.Group("/products")
	pdr.ProductRouter(product, db)

	payment := e.Group("/payments")
	pyr.PaymentRouter(payment, db)

}

func main() {

	middleware.InitLogger()

	godotenv.Load()
	config, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("[ERROR] failed to load configuration: %v", err)
	}

	mdb := database.ConnectMySQL()
	// psql := database.ConnectPostgreSQL()

	e := echo.New()

	middleware.RemoveTrailingSlash(e)
	e.Use(middleware.Logger)
	middleware.RateLimiter(e)
	middleware.Recover(e)
	middleware.CORS(e)

	SetupRoutes(e, mdb)

	host := config.SERVER.SERVER_HOST
	port := config.SERVER.SERVER_PORT
	address := host + ":" + port

	errChan := make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		logrus.Info("[INFO] server is running on address ", address)
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		<-quit
		logrus.Warn("[WARN] shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			logrus.Errorf("[ERROR] error shutting down server: %v", err)
		}
		close(errChan)
	}()

	select {
	case err := <-errChan:
		logrus.Fatalf("[CRITICAL] server error: %v", err)
	case <-time.After(1 * time.Second):
		logrus.Info("[INFO] server is running smoothly...")
	}

	wg.Wait()
	logrus.Info("[INFO] server has been shut down gracefully.")
}
