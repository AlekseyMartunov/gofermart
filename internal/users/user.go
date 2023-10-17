package users

type User struct {
	ID        string
	UUID      string
	Login     string
	Password  string
	Withdrawn float64
	Bonuses   float64
}
