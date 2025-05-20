package user

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}
