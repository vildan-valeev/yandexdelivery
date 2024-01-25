package models

type RobotOpenRequest struct {
	Delay   int64 `json:"delay"`
	PointID int64 `json:"point_id"`
}

type RobotOpenResponse struct {
}
