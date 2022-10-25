package entity

type User struct {
	ID       uint64 `json:"firstname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	BaseModel
}
