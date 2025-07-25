package sourceProvider

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// CPUProvider provides CPU usage information.

type CPUProvider struct {
	valueChan chan float64
}

func (c *CPUProvider) Name() string {
	return "CPU"
}
func (p *CPUProvider) ValueChan() <-chan float64 {
	return p.valueChan
}

// NewCPUProvider creates a new CPUProvider instance.
func NewCPUPercentProvider(ctx context.Context, interval time.Duration) *CPUProvider {

	c := &CPUProvider{
		valueChan: make(chan float64),
	}

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				// 不準確 先這樣 有空再重寫
				percent, err := cpu.Percent(time.Second, false)
				if err != nil {
					continue
				}
				c.valueChan <- percent[0]
			}
		}
	}()

	return c
}

// NewFakeCpuPercentProvider 會從0%開始，每5秒提升10%，到100%回到0%循環直到 ctx.Done()
func NewFakeCpuPercentProvider(ctx context.Context) *CPUProvider {
	c := &CPUProvider{
		valueChan: make(chan float64),
	}
	go func() {
		percent := 0.0
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				close(c.valueChan)
				return
			case <-ticker.C:
				c.valueChan <- percent
				percent += 10.0
				if percent > 100.0 {
					percent = 0.0
				}
			}
		}
	}()
	return c
}