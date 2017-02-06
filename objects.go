package telegrams

type abstractAPIResponse struct {
	Ok          bool               `json:"ok"`
	Description string             `json:"description"`
	ErrorCode   int                `json:"error_code"`
	Parameters  ResponseParameters `json:"ResponseParameters"`
}

type UserAPIResponse struct {
	abstractAPIResponse
	Result User `json:"result"`
}

// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateId      int      `json:"update_id"`
	Password      string   `json:"password"`
	FilterUsers   []string `json:"filterUsers"`
	LoginUrl      string   `json:"loginUrl"`
	PeopleListUrl string   `json:"peopleListUrl"`
}
