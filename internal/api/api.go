package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tylerneath/nuture-backend/internal/config"
	"github.com/tylerneath/nuture-backend/internal/handler"
	"github.com/tylerneath/nuture-backend/internal/repo"
	"github.com/tylerneath/nuture-backend/internal/service"
	"github.com/tylerneath/nuture-backend/internal/store"
	"go.uber.org/zap"
)

func Run(ctx context.Context, cfg config.Config, log *zap.Logger) {
	r := gin.Default()

	db := store.MustCreateNewDB(cfg)
	userRepo := repo.NewUserRepository(db, log)

	if userRepo == nil {
		log.Fatal("userRepo is nil")
	}

	userService := service.NewUserService(ctx, userRepo, log)
	userHandler := handler.NewUserHandler(userService, log)

	r.POST(("/user"), userHandler.CreateUser)
	r.POST("/register", userHandler.RegisterUser)
	r.DELETE("/user", userHandler.DeleteUser)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("listen: ", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", zap.Error(err))
	}

	// catching ctx.Done(). timeout of 5 seconds.

	select {
	case <-ctx.Done():
		log.Info("timeout of 5 seconds.")
	}
	log.Info("Server exiting")

}
