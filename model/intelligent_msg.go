package model

import "time"

type IntelligentMsg struct {
	ID            int       `gorm:"column:id"`
	VehicleImgURL string    `gorm:"column:vehicle_img_url"`
	VehicleLP     string    `gorm:"column:vehicle_lp"`
	VehicleType   string    `gorm:"column:vehicle_type"`
	VehicleColor  string    `gorm:"column:vehicle_color"`
	TaskID        string    `gorm:"column:task_id"`
	TimeStamp     time.Time `gorm:"column:ts"`
}

func (r IntelligentMsg) TableName() string {
	return "intelligent_info"
}
