package handler

// This is a request register account.
type ReqRegister struct {
	Email    string `JSON:"email"`
	Password string `JSON:"password"`
}
