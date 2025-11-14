package metrics

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetMetrics() *Metrics {
	metrics := new(Metrics)
	metrics.Cpu = *getCpu()
	metrics.Ram = *getRam()
	return metrics
}

func getCpu() *CPU {
	// Get CPU Info
	cpuModel := new(CPU)
	infos, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	info := infos[0]
	cpuModel.ModelName = info.ModelName
	cpuModel.Cores = info.Cores
	cpuModel.Mhz = info.Mhz

	// Get CPU Percentage
	time.Sleep(time.Second)
	percentages, err := cpu.Percent(0, true)
	if err != nil {
		log.Fatal(err)
	} else {
		cpuModel.Percent = percentages
	}

	return cpuModel
}

func getRam() *RAM {
	// Get Memory Info
	ramModel := new(RAM)
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	ramModel.Total = virtualMemory.Total
	ramModel.Free = virtualMemory.Free
	ramModel.Used = virtualMemory.Used
	ramModel.UsedPercent = virtualMemory.UsedPercent
	return ramModel
}

type Metrics struct {
	Ram RAM `json:"ram"`
	Cpu CPU `json:"cpu"`
}

type CPU struct {
	ModelName string    `json:"modelName"`
	Cores     int32     `json:"cores"`
	Mhz       float64   `json:"mhz"`
	Percent   []float64 `json:"percentage"`
}

type RAM struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}
