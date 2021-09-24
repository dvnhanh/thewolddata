package port

type TheworlddataService interface {
	Register(email, password string) error
}
