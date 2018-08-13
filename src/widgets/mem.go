package widgets

import (
	"fmt"
	"time"

	"github.com/cjbassi/gotop/src/utils"
	ui "github.com/cjbassi/termui"
	"github.com/phpor/ctools/mem"
)

type Mem struct {
	*ui.LineGraph
	interval time.Duration
}

func NewMem(interval time.Duration, zoom int) *Mem {
	self := &Mem{
		LineGraph: ui.NewLineGraph(),
		interval:  interval,
	}
	self.Label = "Memory Usage"
	self.Zoom = zoom
	self.Data["Main"] = []float64{0}

	self.update()

	ticker := time.NewTicker(self.interval)
	go func() {
		for range ticker.C {
			self.update()
		}
	}()

	return self
}

func (self *Mem) update() {
	memstat,_ := mem.Usage()
	memUsedPercent := float64(memstat.Used)/float64(memstat.Total) * 100
	self.Data["Main"] = append(self.Data["Main"], memUsedPercent)

	mainTotalBytes, mainTotalMagnitude := utils.ConvertBytes(memstat.Total)
	mainUsedBytes, mainUsedMagnitude := utils.ConvertBytes(memstat.Used)
	self.Labels["Main"] = fmt.Sprintf("%3.0f%% %.0f%s/%.0f%s", memUsedPercent, mainUsedBytes, mainUsedMagnitude, mainTotalBytes, mainTotalMagnitude)
}
