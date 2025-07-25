package sourceProvider

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// CPUProvider provides CPU usage information.

type CPUProvider struct {
	valueChan chan float64
	percentBuff []float64
}

func (c *CPUProvider) Name() string {
	return "CPU"
}
func (p *CPUProvider) ValueChan() <-chan float64 {
	return p.valueChan
}

func (c *CPUProvider) addPercent(percent float64) {
	c.percentBuff = append(c.percentBuff, percent)
	if len(c.percentBuff) > 10 {
		c.percentBuff = c.percentBuff[1:]
	}
}

func (c *CPUProvider) getPercent() float64 {
	if len(c.percentBuff) == 0 {
		return 0.0
	}
	sum := 0.0
	for _, percent := range c.percentBuff {
		sum += percent
	}
	return sum / float64(len(c.percentBuff))
}

// NewCPUProvider creates a new CPUProvider instance.
func NewCPUPercentProvider(ctx context.Context, interval time.Duration) *CPUProvider {

	c := &CPUProvider{
		valueChan: make(chan float64),
	}

	go func() {
		ticker := time.NewTicker(interval)
		update := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		defer update.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-update.C:
				percent, err := cpu.Percent(100 * time.Millisecond, false)				
				if err != nil {
					continue
				}
				c.addPercent(percent[0])

			case <-ticker.C:
				// 不準確 先這樣 有空再重寫
				usage := c.getPercent()
				c.valueChan <- usage
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