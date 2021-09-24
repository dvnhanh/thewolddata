package ports

type ThewolddataMysqlRepoS interface {
	Register(email, password string) error
}
