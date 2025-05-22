package model

import (
	"fmt"
	"time"
)

type Stats struct {
	MemoryUsed  uint64
	MemoryTotal uint64
	DiskUsed    uint64
	DiskTotal   uint64
	CPUUsage    []float64
	Uptime      uint64
}

func (s Stats) MemoryUsedGB() string {
	return fmt.Sprintf("%.2fGB", float64(s.MemoryUsed)/(1024*1024*1024))
}

func (s Stats) MemoryTotalGB() string {
	return fmt.Sprintf("%.2fGB", float64(s.MemoryTotal)/(1024*1024*1024))
}

func (s Stats) DiskUsedGB() string {
	return fmt.Sprintf("%.2fGB", float64(s.DiskUsed)/1e9)
}

func (s Stats) DiskTotalGB() string {
	return fmt.Sprintf("%.2fGB", float64(s.DiskTotal)/1e9)
}

func (s Stats) CPUUsagePercent() string {
	return fmt.Sprintf("%.2f%%", s.CPUUsage[0])
}

func (s Stats) UptimeString() string {
	t := time.Duration(s.Uptime) * time.Second
	return t.String()
}
