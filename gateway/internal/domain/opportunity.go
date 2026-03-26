package domain

import "time"

type OpportunityType string

const (
	TypeInternship OpportunityType = "internship"
	TypeJob        OpportunityType = "job"
	TypeEvent      OpportunityType = "event"
	TypeMentorship OpportunityType = "mentorship"
)

type WorkFormat string

const (
	FormatOffice  WorkFormat = "office"
	FormatRemote  WorkFormat = "remote"
	FormatHybrid  WorkFormat = "hybrid"
)

type Opportunity struct {
	ID          string
	Title       string
	Description string

	CompanyID string

	Type   OpportunityType
	Format WorkFormat

	Location string

	Latitude  float64
	Longitude float64

	Salary *int

	Tags []string

	CreatedAt time.Time
	ExpiredAt time.Time
}
