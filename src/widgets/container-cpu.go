package widgets

import (
	"fmt"
	"time"

	ui "github.com/cjbassi/termui"
	"github.com/phpor/ctools/cpu"
)

type ContainerCPU struct {
	*ui.LineGraph
	interval time.Duration
}

func NewContainerCPU(interval time.Duration, zoom int) *ContainerCPU {
	self := &ContainerCPU{
		LineGraph: ui.NewLineGraph(),
		interval:  interval,
	}
	self.Label = "CPU Usage"
	self.Zoom = zoom
	self.Data["Average"] = []float64{0}

	// update asynchronously because of 1 second blocking period
	go self.update()

	ticker := time.NewTicker(self.interval)
	go func() {
		for range ticker.C {
			self.update()
		}
	}()

	return self
}

// calculates the CPU usage over a 1 second interval and blocks for the duration
func (self *ContainerCPU) update() {
		percent := cpu.GetCpuUsage() * 100
		self.Data["Average"] = append(self.Data["Average"], percent)
		self.Labels["Average"] = fmt.Sprintf("%3.0f%%", percent)
}
