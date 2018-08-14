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
	self.Data["Mem"] = []float64{0}

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
	self.Data["Mem"] = append(self.Data["Mem"], memUsedPercent)

	memTotalBytes, memTotalMagnitude := utils.ConvertBytes(memstat.Total)
	memUsedBytes, memUsedMagnitude := utils.ConvertBytes(memstat.Used)
	memCachedBytes, memCachedMagnitude := utils.ConvertBytes(memstat.Cached)
	memCachedPercent := float64(memstat.Cached)/float64(memstat.Total) * 100

	self.Labels["Mem"] = fmt.Sprintf("Total: %.0f%s   Used: %3.1f%%/%.0f%s  cache: %3.1f%%/%.0f%s", memTotalBytes, memTotalMagnitude, memUsedPercent, memUsedBytes, memUsedMagnitude, memCachedPercent,memCachedBytes, memCachedMagnitude)
}
