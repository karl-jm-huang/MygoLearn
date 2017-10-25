package entity

// Date ：
type Date struct {
	Year, Month, Day, Hour, Minute int
}

// Meeting :
type Meeting struct {
	Sponsor            string
	Participators      []string
	StartDate, EndDate Date
	Tittle             string
}
