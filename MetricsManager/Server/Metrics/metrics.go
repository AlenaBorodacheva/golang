package metrics

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetMetrics() *Metrics {
	getCpu()
	getRam()
	return nil
}

func getCpu() {
	// Get CPU Info
	info, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}
	for _, ci := range info {
		fmt.Printf("CPU Info: %+v\n", ci)
	}

	// Get CPU Percentage
	percentages, err := cpu.Percent(0, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CPU Percentages: %+v\n", percentages)
}

func getRam() {
	// Get Memory Info
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total RAM: %v bytes\n", virtualMemory.Total)
	fmt.Printf("Free RAM: %v bytes\n", virtualMemory.Free)
	fmt.Printf("Used RAM: %v bytes\n", virtualMemory.Used)
	fmt.Printf("Used percent: %f%%\n", virtualMemory.UsedPercent)
}

type Metrics struct {
	Ram RAM
}

type CPU struct {
}

type RAM struct {
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}
