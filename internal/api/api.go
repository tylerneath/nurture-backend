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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		claims, err := service.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}

}

func Run(ctx context.Context, cfg config.Config, log *zap.Logger) {
	// load the env vairable for the key
	r := gin.New()

	db := store.MustCreateNewDB(cfg)

	// create the repositories
	userRepo := repo.NewUserRepository(db, log)
	messageRepo := repo.NewMessageRepository(db, log)

	// create the services
	userService := service.NewUserService(ctx, userRepo, messageRepo, log)
	messageService := service.NewMessageService(ctx, messageRepo, log)

	// create the handlers
	userHandler := handler.NewUserHandler(userService, log)
	messageHandler := handler.NewMessageHandler(messageService, log)

	// create user management routes 
	r.POST("/register", userHandler.RegisterUser)
	r.POST("/login", userHandler.LoginUser)

	protected := r.Group("/auth")
	protected.Use(AuthMiddleware())

	apiV1 := r.Group("/api/v1")
	apiV1.POST("/users", userHandler.CreateUser)
	apiV1.GET("/users/:id", userHandler.GetUserByID)
	apiV1.PUT("/users/:id", userHandler.UpdateUser)
	apiV1.DELETE("/users/:id", userHandler.DeleteUser)

	apiV1.POST("/messages", messageHandler.CreateMessage)
	apiV1.GET("/messages/:id", messageHandler.GetMessageByID)
	apiV1.PUT("/messages/:id", messageHandler.UpdateMessage)
	apiV1.DELETE("/messages/:id", messageHandler.DeleteMessage)

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

	select {
	case <-ctx.Done():
		log.Info("timeout of 5 seconds.")
	}
	log.Info("Server exiting")

}
