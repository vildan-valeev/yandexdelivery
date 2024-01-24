package models

type Status string

const (
	StatusNew                          Status = "new"
	StatusEstimating                   Status = "estimating"
	StatusEstimatingFailed             Status = "estimating_failed"
	StatusReadyForApproval             Status = "ready_for_approval"
	StatusAccepted                     Status = "accepted"
	StatusPerformerLookup              Status = "performer_lookup"
	StatusPerformerDraft               Status = "performer_draft"
	StatusPerformerFound               Status = "performer_found"
	StatusPerformerNotFound            Status = "performer_not_found"
	StatusPickupArrived                Status = "pickup_arrived"
	StatusReadyForPickupConfirmation   Status = "ready_for_pickup_confirmation"
	StatusPickUped                     Status = "pickuped"
	StatusDeliveryArrived              Status = "delivery_arrived"
	StatusReadyForDeliveryConfirmation Status = "ready_for_delivery_confirmation"
	StatusPayWaiting                   Status = "pay_waiting"
	StatusDelivered                    Status = "delivered"
	StatusDeliveredFinish              Status = "delivered_finish"
	StatusReturning                    Status = "returning"
	StatusReturnArrived                Status = "return_arrived"
	StatusReadyForReturnConfirmation   Status = "ready_for_return_confirmation"
	StatusReturned                     Status = "returned"
	StatusReturnedFinish               Status = "returned_finish"
	StatusFailed                       Status = "failed"
	StatusCancelled                    Status = "cancelled"
	StatusCancelledWithPayment         Status = "cancelled_with_payment"
	StatusCancelledByTaxi              Status = "cancelled_by_taxi"
	StatusCancelledWithItemsOnHands    Status = "cancelled_with_items_on_hands"
)
