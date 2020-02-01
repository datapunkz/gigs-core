package events

// Events Object
type Events struct {
	Ownerid       string `json:"ownerid"`
	Eventid       string `json:"eventid"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	EventURL      string `json:"event_url"`
	EventContacts []struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	} `json:"event_contacts"`
	Venues []struct {
		Name     string `json:"name"`
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City     string `json:"city"`
		State    string `json:"state"`
		Zip      string `json:"zip"`
	} `json:"venues"`
	Submissions []struct {
		Talkid       string `json:"talkid"`
		Name         string `json:"name"`
		SubmitterID  string `json:"submitter_id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		SubmitDate   string `json:"submit_date"`
		AcceptedDate string `json:"accepted_date"`
		Status       string `json:"status"`
	} `json:"submissions"`
}
