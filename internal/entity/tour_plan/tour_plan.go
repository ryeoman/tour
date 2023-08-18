package entity

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
