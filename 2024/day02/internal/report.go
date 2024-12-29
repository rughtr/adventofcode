package internal

type Report struct {
	Levels []int
}

func NewReport(levels []int) Report {
	return Report{
		Levels: levels,
	}
}
