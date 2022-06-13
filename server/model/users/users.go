package users

type User struct {
	Id           int    `json:"-" db:"id"`
	Name         string `json:"name" db:"name"`
	Surname      string `json:"surname" db:"surname"`
	FatherName   string `json:"father_name" db:"father_name"`
	Nickname     string `json:"nickname" db:"nickname"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
}
