package repository

import (
	"encoding/json"
	"go_ecommerce/models"
)

func checkDataRequest(dataRequst models.CheckDataRequest) map[string]bool {
	datajson, err := json.Marshal(dataRequst)
	result := map[string]bool{}
	if err != nil {
		return result
	}
	json.Unmarshal([]byte(datajson), &result)
	return result
}
