package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseRobotCheckAvailability struct {
	models.RobotCheckAvailabilityResponse
	APIResponseBase
}

func (a APIResponseRobotCheckAvailability) Base() APIResponseBase {
	return a.APIResponseBase
}
