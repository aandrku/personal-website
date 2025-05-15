package analytics

var visits int

type Service struct {
}

func (Service) IncrementVisits() {
	visits++
}

func (Service) TotalVisits() int {
	return visits
}
