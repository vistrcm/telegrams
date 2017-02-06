package telegrams

// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type update struct {
	UpdateId      int      `json:"update_id"`
	Password      string   `json:"password"`
	FilterUsers   []string `json:"filterUsers"`
	LoginUrl      string   `json:"loginUrl"`
	PeopleListUrl string   `json:"peopleListUrl"`
}
