package models

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int16  `json:"age"`
}

func NewUser(id int64, name string, email string, age int16) *User {
	return &User{Id: id, Name: name, Email: email, Age: age}
}
