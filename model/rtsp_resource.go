package model

type RTSPResource struct {
	ID           int    `gorm:"id"`
	URL          string `gorm:"url"`
	Position     string `gorm:"position"`
	AlgModel     int    `gorm:"model"`
	EncodeNeeded int    `gorm:"encodeneeded"`
}

func (r RTSPResource) TableName() string {
	return "rtsp_resource"
}
