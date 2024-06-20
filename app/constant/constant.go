package constant

type ResponseStatus int

const (
	Success ResponseStatus = iota
	NotFound
	UnknownError
	InvalidRequest
	Unauthorized
)

func (r ResponseStatus) ResponseStatus() string {
	return [...]string{"Success", "NotFound", "UnknownError", "InvalidRequest", "Unauthorized"}[r]
}
