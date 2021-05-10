package pkg

import (
	"ColdMountain/connection"
	model2 "ColdMountain/graphql/graph/model"
	"ColdMountain/model"
	"log"
	"time"
)

func GetIntelligentMsgs(timeStampStart *string, timeStampEnd *string, vehicleLp *string, vehicleColor *string, taskID *string) []*model2.IntelligentMsg {
	if intelligentMsgs, err := GetIntelligentMsgsFromDB(timeStampStart, timeStampEnd, vehicleLp, vehicleColor, taskID); err != nil {
		return []*model2.IntelligentMsg{}
	} else {
		var ret = make([]*model2.IntelligentMsg, len(intelligentMsgs))
		for i := 0; i < len(intelligentMsgs); i++ {
			ts := time.Unix(intelligentMsgs[i].TimeStamp.Unix(), 0).Format("2006-01-02 15:04:05")
			ret[i] = &model2.IntelligentMsg{
				ID:            &intelligentMsgs[i].ID,
				VehicleImgURL: &intelligentMsgs[i].VehicleImgURL,
				VehicleLp:     &intelligentMsgs[i].VehicleLP,
				VehicleType:   &intelligentMsgs[i].VehicleType,
				VehicleColor:  &intelligentMsgs[i].VehicleColor,
				TaskID:        &intelligentMsgs[i].TaskID,
				TimeStamp:     &ts,
			}
		}
		return ret
	}
}

func GetIntelligentMsgsFromDB(timeStampStart *string, timeStampEnd *string, vehicleLp *string, vehicleColor *string, taskID *string) ([]model.IntelligentMsg, error) {
	cl := connection.GetRTSPResourceDBClient()
	intelligentMsgs := []model.IntelligentMsg{}
	whereStat := ""
	if *timeStampStart != "" {
		whereStat += "ts >= '" + *timeStampStart + "' "
	}
	if *timeStampEnd != "" {
		if whereStat != "" {
			whereStat += "and "
		}
		whereStat += "ts <= '" + *timeStampEnd + "' "
	}
	if *vehicleLp != "" {
		if whereStat != "" {
			whereStat += "and "
		}
		whereStat += "vehicle_lp = '" + *vehicleLp + "' "
	}
	if *vehicleColor != "" {
		if whereStat != "" {
			whereStat += "and "
		}
		whereStat += "vehicle_color = '" + *vehicleColor + "' "
	}
	if *taskID != "" {
		if whereStat != "" {
			whereStat += "and "
		}
		whereStat += "task_id = '" + *taskID + "' "
	}
	if dbErr := cl.Where(whereStat).Order("ts ASC").Find(&intelligentMsgs).Error; dbErr != nil {
		log.Printf("dberr: %+v", dbErr)
		return nil, dbErr
	}
	return intelligentMsgs, nil
}
