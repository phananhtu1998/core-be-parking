package initialize

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

// GetServerInfo returns basic system information
func GetServerInfo() {
	// Get CPU info
	cpuInfo, _ := cpu.Info()
	cpuPercent, _ := cpu.Percent(time.Second, false)

	// Get memory info
	memInfo, _ := mem.VirtualMemory()

	// Get disk info
	diskInfo, _ := disk.Usage("/")

	// Get host info
	hostInfo, _ := host.Info()

	fmt.Printf("CPU Model: %v\n", cpuInfo[0].ModelName)
	fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])
	fmt.Printf("Memory Total: %v GB\n", memInfo.Total/1024/1024/1024)
	fmt.Printf("Memory Used: %.2f%%\n", memInfo.UsedPercent)
	fmt.Printf("Disk Total: %v GB\n", diskInfo.Total/1024/1024/1024)
	fmt.Printf("Disk Used: %.2f%%\n", diskInfo.UsedPercent)
	fmt.Printf("OS: %v\n", hostInfo.OS)
	fmt.Printf("Platform: %v\n", hostInfo.Platform)
	fmt.Printf("Hostname: %v\n", hostInfo.Hostname)
}
