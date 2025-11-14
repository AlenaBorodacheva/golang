package writer

import (
	"encoding/json"

	"gorm.io/gorm"
)

func WriteRamToSql(db *gorm.DB, data []byte, port int) {
	var rams []RamModel
	json.Unmarshal(data, &rams)
	for i := range rams {
		rams[i].Port = port
		db.Create(&rams[i])
	}
}

func WriteCpuToSql(db *gorm.DB, data []byte, port int) {
	var cpus []CpuModel
	json.Unmarshal(data, &cpus)
	for i := range cpus {
		cpus[i].Port = port
		db.Create(&cpus[i])
	}
}
