package domain

type CacheConnection struct {
	Conn interface{}
	Err  error
}

type CacheResponse struct {
	Response ResponseOperation
	Status   StatusOperation
}
type ResponseOperation struct {
	Data string
}

type StatusOperation struct {
	Status string
	Err    error
}

func (cr *CacheResponse) SetStatus(Err error) {

	if Err != nil {
		cr.Status.Status = "FAIL"
	} else {
		cr.Status.Status = "SUCCESS"
	}
	cr.Status.Err = Err
}

func (cr *CacheResponse) SetData(data string) {
	cr.Response.Data = data
}

func (cr *CacheResponse) Data() string {
	return cr.Response.Data
}

func (cr *CacheResponse) Error() string {
	if cr.Status.Err != nil {
		return cr.Status.Err.Error()
	}
	return ""
}
