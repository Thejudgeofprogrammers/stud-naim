package domain

import "time"

type ResponseStatus string

const (
	ResponseNew      ResponseStatus = "new"
	ResponseAccepted ResponseStatus = "accepted"
	ResponseRejected ResponseStatus = "rejected"
	ResponseReserve  ResponseStatus = "reserve"
)

type Response struct {
	ID            string         `json:"id"`
	UserID        string         `json:"user_id"`
	OpportunityID string         `json:"opportunity_id"`
	Status        ResponseStatus `json:"status"`
	AppliedAt     time.Time      `json:"applied_at"`
	Message       string         `json:"message,omitempty"`
}