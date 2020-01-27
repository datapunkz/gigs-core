package talks

//Talks Object
type Talks struct {
	Ownerid     string `json:"ownerid"`
	Talkid      string `json:"talkid"`
	Title       string `json:"title"`
	Abstract    string `json:"abstract"`
	Description string `json:"description"`
	Publishdate string `json:"publishdate"`
	Authors     []struct {
		Userid    string `json:"userid"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"authors"`
	Resources []struct {
		Media string `json:"media"`
		URL   string `json:"url"`
	} `json:"resources"`
}
