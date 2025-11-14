package metrics

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func GetCpu() *CPU {
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

	cpuModel.Date = time.Now()

	return cpuModel
}

func GetRam() *RAM {
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
	ramModel.Date = time.Now()
	return ramModel
}

type CPU struct {
	ModelName string    `json:"modelName"`
	Cores     int32     `json:"cores"`
	Mhz       float64   `json:"mhz"`
	Percent   []float64 `json:"percentage"`
	Date      time.Time `json:"datetime"`
}

type RAM struct {
	Total       uint64    `json:"total"`
	Free        uint64    `json:"free"`
	Used        uint64    `json:"used"`
	UsedPercent float64   `json:"usedPercent"`
	Date        time.Time `json:"datetime"`
}
