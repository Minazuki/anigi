package sourceProvider

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/mem"
)

// MemProvider provides memory usage information.
type MemProvider struct {
	valueChan chan float64
}

func (m *MemProvider) Name() string {
	return "Memory"
}

func (m *MemProvider) ValueChan() <-chan float64 {
	return m.valueChan
}

// NewMemProvider creates a new MemProvider instance.
func NewMemProvider(ctx context.Context, interval time.Duration) *MemProvider {

	m := &MemProvider{
		valueChan: make(chan float64),
	}

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				v, err := mem.VirtualMemory()
				if err != nil {
					continue
				}
				m.valueChan <- v.UsedPercent
			}
		}
	}()

	return m
}
