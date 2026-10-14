package repository

type User struct {
}

type UserRepository interface {
	GetUserByPubKey(pubKey string) (User, error)
}
