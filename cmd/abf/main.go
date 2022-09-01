package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/WHGLD/antibruteforce/internal/app"
	"github.com/WHGLD/antibruteforce/internal/cache"
	"github.com/WHGLD/antibruteforce/internal/server"
	"github.com/jonboulle/clockwork"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	dbConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Name,
	)
	dbRepository, err := app.NewRepository(dbConnection)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbRepository.Close()

	redisCache, err := cache.New(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer redisCache.Close()

	limiter := app.NewRateLimiter(
		*redisCache,
		clockwork.NewRealClock(),
		cfg.Limits.Login,
		cfg.Limits.Password,
		cfg.Limits.IP,
	)

	grpcServer := server.NewServer(
		*app.NewApp(*limiter, *dbRepository),
		cfg.GRPC.Host+":"+strconv.Itoa(cfg.GRPC.Port),
	)

	go func() {
		<-ctx.Done()

		ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), time.Second*3)
		defer cancelTimeout()

		if errServer := grpcServer.Stop(ctxTimeout); errServer != nil {
			fmt.Println(errServer.Error())
		}
	}()

	if errServer := grpcServer.Start(ctx); errServer != nil {
		fmt.Println(errServer.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
