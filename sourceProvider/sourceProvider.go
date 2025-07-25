package sourceProvider

import (
	"context"
	"fmt"
	"time"
)

type SourceProvider interface {
	Name() string
	ValueChan() <-chan float64
}



func NewSourceProvider(name string, ctx context.Context, interval time.Duration) (SourceProvider, error) {
	switch name {
	case "cpu":
		return NewCPUPercentProvider(ctx, interval), nil
	case "fakecpu":
		return NewFakeCpuPercentProvider(ctx), nil
	case "mem":
		return NewMemProvider(ctx, interval), nil
	default:
		return nil, fmt.Errorf("unknown source provider: %s", name)
	}
}


