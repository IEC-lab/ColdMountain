package pkg

import (
	"ColdMountain/connection"
	model2 "ColdMountain/graphql/graph/model"
	"ColdMountain/model"
	"log"
)

func DiscoverStreams() []*model2.FrameStream {
	var ret = []*model2.FrameStream{}
	if RTSPResources, err := GetRTSPResources(); err != nil{
		return ret
	}else {
		for _, RTSPResource := range RTSPResources{
			temp := model2.FrameStream{
				ID:           &RTSPResource.ID,
				URL:          &RTSPResource.URL,
				Position:     &RTSPResource.Position,
				AlgModel:     &RTSPResource.AlgModel,
				EncodeNeeded: &RTSPResource.EncodeNeeded,
			}
			ret = append(ret, &temp)
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