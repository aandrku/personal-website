package model

type Stats struct {
	MemoryUsed  uint64
	MemoryTotal uint64
	DiskUsed    uint64
	DiskTotal   uint64
	CPUUsage    []float64
	Uptime      uint64
}
