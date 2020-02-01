package users

// User Object
type User struct {
	Ownerid   string `json:"ownerid"`
	Userid    string `json:"userid"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Accounts  []struct {
		Provider   string `json:"provider"`
		Username   string `json:"username"`
		Userid     string `json:"userid"`
		ProfileURL string `json:"profile_url"`
	} `json:"accounts"`
}
