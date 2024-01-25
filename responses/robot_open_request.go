package responses

import "github.com/vildan-valeev/yandexlogistic/models"

type APIResponseRobotOpen struct {
	models.RobotOpenResponse
	APIResponseBase
}

func (a APIResponseRobotOpen) Base() APIResponseBase {
	return a.APIResponseBase
}
