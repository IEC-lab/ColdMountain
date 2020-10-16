package connection

import (
	"ColdMountain/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var aRTSPResourceDBClient *gorm.DB

func init() {
	coldRepo := conf.GetGlobalConfig().ColdRepo
	var dbErr error
	aRTSPResourceDBClient, dbErr = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/colddb?charset=utf8mb4&parseTime=True&loc=Local",
		coldRepo.User, coldRepo.Password, coldRepo.Address, coldRepo.Port))
	if dbErr != nil {
		log.Panicf("dberr: %+v", dbErr)
	}
}

func GetRTSPResourceDBClient() *gorm.DB {
	return aRTSPResourceDBClient
}
