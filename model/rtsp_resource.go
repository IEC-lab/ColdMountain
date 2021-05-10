package model

type RTSPResource struct {
	ID           int    `gorm:"column:id"`
	URL          string `gorm:"column:url"`
	Position     string `gorm:"column:position"`
	AlgModel     int    `gorm:"column:model"`
	EncodeNeeded int    `gorm:"column:encodeneeded"`
}

func (r RTSPResource) TableName() string {
	return "rtsp_resource"
}
