package data

type Employee struct {
	ID          string `json:"ID"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Designation string `json:"designation"`
	Location    string `json:"location"`
	Company     string `json:"company"`
}
