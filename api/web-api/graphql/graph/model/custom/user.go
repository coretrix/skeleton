package custom

type User struct {
	ID       string
	Username string
	Age      int64
}

func (u *User) Pet() *Animal {
	return &Animal{Nickname: "kiopek"}
}
