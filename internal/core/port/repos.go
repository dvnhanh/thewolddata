package port

type TheworlddataMysqlRepoS interface {
	Register(email, password string) error
}
