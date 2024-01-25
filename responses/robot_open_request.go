package responses

import "github.com/vildan-valeev/yandexdelivery/models"

type APIResponseRobotOpen struct {
	models.RobotOpenResponse
	APIResponseBase
}

func (a APIResponseRobotOpen) Base() APIResponseBase {
	return a.APIResponseBase
}
