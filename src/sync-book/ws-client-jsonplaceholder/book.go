package ph_client

type Book struct {
	Id     int `json:"userid"`
	UserId int `json:"id"`
	Title  string
	Ok     bool `json:"completed"`
}
