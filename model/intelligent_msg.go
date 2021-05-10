package model

type IntelligentMsg struct {
	ID            int    `gorm:"id"`
	VehicleImgURL string `gorm:"vehicle_img_url"`
	VehicleLP     string `gorm:"vehicle_lp"`
	VehicleType   string `gorm:"vehicle_type"`
	VehicleColor  string `gorm:"vehicle_color"`
	TaskID        string `gorm:"task_id"`
	TimeStamp     string `gorm:"ts"`
}

func (r IntelligentMsg) TableName() string {
	return "intelligent_info"
}
