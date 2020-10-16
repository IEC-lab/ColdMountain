package pkg

import (
	"ColdMountain/connection"
	"ColdMountain/model"
	"log"
)

func DiscoverStreams() []string {
	if RTSPResources, err := GetRTSPResources(); err != nil{
		return []string{}
	}else {
		var ret = []string{}
		for _, RTSPResource := range RTSPResources{
			ret = append(ret, RTSPResource.URL)
		}
		return ret
	}
}

func GetRTSPResources() ([]model.RTSPResource, error) {
	RTSPResourceDBClient := connection.GetRTSPResourceDBClient()
	RTSPResources := []model.RTSPResource{}
	if dbErr := RTSPResourceDBClient.Find(&RTSPResources).Error; dbErr != nil{
		log.Printf("dberr: %+v", dbErr)
		return nil, dbErr
	}
	return RTSPResources, nil
}