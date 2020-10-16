package pkg

import (
	"ColdMountain/connection"
	model2 "ColdMountain/graphql/graph/model"
	"ColdMountain/model"
	"log"
)

func DiscoverStreams() []*model2.FrameStream {
	if RTSPResources, err := GetRTSPResources(); err != nil {
		return []*model2.FrameStream{}
	} else {
		var ret = make([]*model2.FrameStream, len(RTSPResources))
		for i := 0; i < len(RTSPResources); i++ {
			ret[i] = &model2.FrameStream{
				ID:           &RTSPResources[i].ID,
				URL:          &RTSPResources[i].URL,
				Position:     &RTSPResources[i].Position,
				AlgModel:     &RTSPResources[i].AlgModel,
				EncodeNeeded: &RTSPResources[i].EncodeNeeded,
			}
		}
		return ret
	}
}

func GetRTSPResources() ([]model.RTSPResource, error) {
	RTSPResourceDBClient := connection.GetRTSPResourceDBClient()
	RTSPResources := []model.RTSPResource{}
	if dbErr := RTSPResourceDBClient.Find(&RTSPResources).Error; dbErr != nil {
		log.Printf("dberr: %+v", dbErr)
		return nil, dbErr
	}
	return RTSPResources, nil
}
