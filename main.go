package engine

import (
	"context"
	"github.com/yunnet/engine"
	"time"
)

var config struct {
	EnableAudio    bool
	EnableVideo    bool
	PublishTimeout time.Duration
	MaxRingSize    int
}

func init() {
	engine.InstallPlugin(&engine.PluginConfig{
		Name:   "Core",
		Config: &config,
		Run:    run,
	})
}

func run(ctx context.Context) {
	engine.Print("::::core::::")
}
