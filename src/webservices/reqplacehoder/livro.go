package reqplacehoder

type Book struct {
	Code       int    `json:"id"`
	Id         int    `json:"userId"`
	Nome       string `json:"title"`
	Finalizado bool   `json:"completed"`
}
