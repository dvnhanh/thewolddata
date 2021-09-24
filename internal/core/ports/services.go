package ports

type ThewolddataService interface {
	Register(email, password string) error
}
