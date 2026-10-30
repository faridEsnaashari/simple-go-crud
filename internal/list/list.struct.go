package list

import "time"

type ListStatus string

const (
	ListStatusCreated    ListStatus = "created"
	ListStatusInProgress ListStatus = "inProgress"
	ListStatusDone       ListStatus = "done"
)

type List struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Status    ListStatus `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
