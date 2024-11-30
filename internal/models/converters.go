package models

func (u *UserRes) ToDto(user User) {
	u.ID = user.ID
	u.Name = user.Name
	u.Email = user.Email
	u.Password = string(user.Password)
	u.CreatedAt = user.CreatedAt.String()
	u.UpdatedAt = user.UpdatedAt.String()
}

func (u *User) ToEntity(user UserReq) {
	u.Name = user.Name
	u.Email = user.Email
	u.Password = []byte(user.Password)
}
