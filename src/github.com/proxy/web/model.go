package web

import (
	"bytes"
	"net/http"
)

//struct para concatenar string sem gerar um novo objeto em memoria
type ConcatString struct {
	buffer bytes.Buffer
}

func (c *ConcatString) Add(String string) *ConcatString {
	c.buffer.WriteString(String)
	return c
}
func (c *ConcatString) Build() string {
	return c.buffer.String()
}

type ResponseModel struct {
	DataContent    string
	Error          []error
	HttpStatusCode int
	HttpResponse   *http.Response
}
