
package collector

import (
	"time"

	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
)

type SystemStats struct {
	CPUUsage    float64
	MemoryUsed  float64
	MemoryTotal float64
	DiskUsed    float64
	DiskTotal   float64
	Uptime      time.Duration
}


func GetSystemStats() (SystemStats,error){
	cpuPercentage , err := cpu.Percent(time.Second,false)
	if err != nil {
		return SystemStats{}, err
	}
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, err
	}
	diskStats , err  := disk.Usage("/")
	if err != nil {
		return SystemStats{}, err
	}
	uptime , err := host.Uptime()
	if err != nil {
		return SystemStats{}, err
	}


	   sysStats := SystemStats{
			   CPUUsage:    cpuPercentage[0],
			   MemoryUsed:  float64(memStats.Used) / 1e9,
			   MemoryTotal: float64(memStats.Total) / 1e9,
			   DiskUsed:    float64(diskStats.Used) / 1e9,
			   DiskTotal:   float64(diskStats.Total) / 1e9,
			   Uptime:      time.Duration(uptime) * time.Second,
	   }
	return sysStats , nil
}