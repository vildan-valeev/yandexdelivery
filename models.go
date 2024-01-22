package yandexlogistic

//
////type ResponseInfo struct {
////	ID                 string        `json:"id,omitempty"`
////	RoutePoints        []RoutePoints `json:"route_points,omitempty"`
////	Pricing            Pricing       `json:"pricing,omitempty"`
////	ClientRequirements Requirements  `json:"client_requirements,omitempty"`
////	Status             string        `json:"status,omitempty"`
////	Version            int32         `json:"version,omitempty"`
////	ErrorMessages      []*APIError   `json:"error_messages,omitempty"`
////	Due                string        `json:"due,omitempty"`
////}
//
////// CreateTripRequest запрос метода создания заявки на поездку API Yandex.Taxi
////type CreateTripRequest struct {
////	RoutePoints        []RoutePoints `json:"route_points,omitempty"`
////	ClientRequirements Requirements  `json:"client_requirements,omitempty"`
////	Items              []Items       `json:"items,omitempty"`
////	EmergencyContact   Contact       `json:"emergency_contact,omitempty"`
////	Due                *string       `json:"due,omitempty"`
////}
//
//// Requirements требования к поездке
//type Requirements struct {
//	TaxiClass string `json:"taxi_class,omitempty"`
//}
//
//// Coordinates координаты поездки
//type Coordinates struct {
//	Coordinates []float64 `json:"coordinates,omitempty"`
//}
//
//// Items данные товаров
//type Items struct {
//	Quantity     float64 `json:"quantity,omitempty"`
//	PickupPoint  int64   `json:"pickup_point,omitempty"`
//	DroppofPoint int64   `json:"droppof_point,omitempty"`
//	Title        string  `json:"title,omitempty"`
//	CostValue    string  `json:"cost_value,omitempty"`
//	CostCurrency string  `json:"cost_currency,omitempty"`
//}
//
//// ServiceLevels условия поездки
//type ServiceLevels struct {
//	PriceRaw float64 `json:"price_raw,omitempty"`
//}
//
//// RoutePoints точки маршрута поездки
//type RoutePoints struct {
//	Address         Address `json:"address,omitempty"`
//	Type            string  `json:"type,omitempty"`
//	PointID         int64   `json:"point_id,omitempty"`
//	VisitOrder      int64   `json:"visit_order,omitempty"`
//	Contact         Contact `json:"contact,omitempty"`
//	ExternalOrderID string  `json:"external_order_id,omitempty"`
//}
//
//// Contact контакты точек маршрута
//type Contact struct {
//	Name  string `json:"name,omitempty"`
//	Phone string `json:"phone,omitempty"`
//	Email string `json:"email,omitempty"`
//}
//
//// Address координаты поездки
//type Address struct {
//	Coordinates []float64 `json:"coordinates,omitempty"`
//	FullName    string    `json:"fullname,omitempty"`
//	Comment     string    `json:"comment,omitempty"`
//}
//
//// CreateTripResponse ответ метода создания заявки на поездку API Yandex.Taxi
//type CreateTripResponse struct {
//	ID                 string        `json:"id,omitempty"`
//	RoutePoints        []RoutePoints `json:"route_points,omitempty"`
//	Pricing            Pricing       `json:"pricing,omitempty"`
//	ClientRequirements Requirements  `json:"client_requirements,omitempty"`
//	Status             string        `json:"status,omitempty"`
//	Version            int32         `json:"version,omitempty"`
//}
//
//// Pricing стоимость поездки
//type Pricing struct {
//	Currency   string `json:"currency,omitempty"`
//	FinalPrice string `json:"final_price,omitempty"`
//	Offer      Offer  `json:"offer,omitempty"`
//}
//
//// Offer предложение
//type Offer struct {
//	OfferID string `json:"offer_id,omitempty"`
//	Price   string `json:"price,omitempty"`
//}
