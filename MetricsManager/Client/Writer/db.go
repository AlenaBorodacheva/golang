package writer

import (
	"time"

	sqlite "github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("DB connection error: " + err.Error())
	}

	if err := db.AutoMigrate(&CpuModel{}, &RamModel{}); err != nil {
		panic("Migration error: " + err.Error())
	}

	return db
}

type CpuModel struct {
	ID        int `gorm:"primaryKey"`
	Port      int
	ModelName string    `json:"modelName"`
	Cores     int       `json:"cores"`
	Mhz       float64   `json:"mhz"`
	Percent   string    `json:"percentage"`
	Date      time.Time `json:"datetime"`
}

type RamModel struct {
	ID          int `gorm:"primaryKey"`
	Port        int
	Total       uint64    `json:"total"`
	Free        uint64    `json:"free"`
	Used        uint64    `json:"used"`
	UsedPercent float64   `json:"usedPercent"`
	Date        time.Time `json:"datetime"`
}
