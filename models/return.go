package models

type ReturnRequest struct {
	Comment         string `json:"comment"`
	NeedReturnItems bool   `json:"need_return_items"`
	PointID         int64  `json:"point_id"`
}

type ReturnResponse struct {
}
