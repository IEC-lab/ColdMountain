package model

type RTSPResource struct {
	ID           int64  `gorm:"id"`
	URL          string `gorm:"url"`
	Position     string `gorm:"position"`
	AlgModel     int64  `gorm:"model"`
	EncodeNeeded int8   `gorm:"encodeneeded"`
}

func (r RTSPResource) TableName() string {
	return "rtsp_resource"
}
