package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"star-golang-migrations/configs"
	"star-golang-migrations/pkg/controller"
	"star-golang-migrations/pkg/di"
	"syscall"
)

func Execute() {
	ctx, cancel := NewCtx()
	defer cancel()

	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
		return
	}

	ghs := di.InitializeGitHubController()
	if err = controller.Run(ctx, config.GithubToken, ghs); err != nil {
		log.Fatal(err)
		return
	}
}

func NewCtx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		trap := make(chan os.Signal, 1)
		signal.Notify(trap, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT)
		<-trap
	}()

	return ctx, cancel
}
