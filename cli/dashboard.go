package cli

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"github.com/AdityaAnandCodes/GoStats/collector"
	"log"
	"github.com/fatih/color"
)

func clearScreen() {
	cmd := exec.Command("clear")
	if _, exists := os.LookupEnv("OS"); exists && os.Getenv("OS") == "Windows_NT" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func formatUptime(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60

	if days > 0 {
		return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
	}
	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}
	return fmt.Sprintf("%dm", minutes)
}

func RenderDashboard() {
	stats, err := collector.GetSystemStats()
	if err != nil {
		log.Println("Error Fetching the Statistics:", err)
		return
	}
	color.Green("======== GoStats CLI Dashboard ========")
	fmt.Printf("CPU Usage:       %.2f%%\n", stats.CPUUsage)
	fmt.Printf("Memory Used:     %.2f GB / %.2f GB\n", stats.MemoryUsed, stats.MemoryTotal)
	fmt.Printf("Disk Used:       %.2f GB / %.2f GB\n", stats.DiskUsed, stats.DiskTotal)
	fmt.Printf("System Uptime:   %s\n", formatUptime(stats.Uptime))
	fmt.Println("========================================")
	fmt.Println(`
	 ____       _____           _     
	/ ___| ___ | ____|_ __   __| |___ 
	| |  _ / _ \|  _| | '_ \ / _  / __|
	| |_| | (_) | |___| | | | (_| \__ \
	\____|\___/|_____|_| |_|\__,_|___/
	`)
}

func LiveDashboard() {
	for {
		clearScreen()
		RenderDashboard()
		time.Sleep(10 * time.Second)
	}
}