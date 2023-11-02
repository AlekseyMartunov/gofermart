package mockloginhandlers

import "AlekseyMartunov/internal/users"

type UserDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (dto *UserDTO) ToEntity() users.User {
	return users.User{
		Login:    dto.Login,
		Password: dto.Password,
	}
}
