package bidtracker

import (
	"github.com/gofrs/uuid"
)

// Bid struct stores a bid for a given item
type Bid struct {
	ItemUUID  uuid.UUID `json:"itemuuid"`
	UserUUID  uuid.UUID `json:"useruuid"`
	Timestamp int64     `json:"timestamp"`
	Amount    float64   `json:"amount"`
}
