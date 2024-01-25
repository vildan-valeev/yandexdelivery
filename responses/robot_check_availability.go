package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseRobotCheckAvailability struct {
	models.RobotCheckAvailabilityResponse
	APIResponseBase
}

func (a APIResponseRobotCheckAvailability) Base() APIResponseBase {
	return a.APIResponseBase
}
