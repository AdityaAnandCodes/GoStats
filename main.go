package main

import (
	"fmt"
	"log"
	"github.com/AdityaAnandCodes/GoStats/collector"
)


func main(){

	   stats, err := collector.GetSystemStats()
	   if err != nil {
			   log.Fatal("Failed to get system stats: ", err)
	   }

	   fmt.Println("System Stats:")
	   fmt.Printf("CPU Usage: %.2f%%\n", stats.CPUUsage)
	   fmt.Printf("Memory: %.2f GB / %.2f GB\n", stats.MemoryUsed, stats.MemoryTotal)
	   fmt.Printf("Disk: %.2f GB / %.2f GB\n", stats.DiskUsed, stats.DiskTotal)
	   fmt.Printf("Uptime: %s\n", stats.Uptime)

}