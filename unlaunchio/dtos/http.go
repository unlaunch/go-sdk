package dtos

// HTTPError ...
type HTTPError struct {
	Code    int
	Msg string
}

// Error to implement Go error interface
func (e HTTPError) Error() string {
	return e.Msg
}