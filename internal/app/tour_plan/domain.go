package domain

// DDD concept => An "entity" is a specific type of domain object.

// TourPlan tour plan domain entity.
type TourPlan struct {
	ID                   int
	CustomerName         string
	FromLocation         string
	TourPackage          string
	VisitPlan            string
	NumberOfParticipants int
	TimestampUnix        int64
}
