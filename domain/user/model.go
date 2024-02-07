package user

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewMockIDUser(firstName string, lastName string) User {
	return User{
		ID:        1,
		FirstName: firstName,
		LastName:  lastName,
	}
}
