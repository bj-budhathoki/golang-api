package dtos

type UserUpdateDTOS struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"name" form:"email" binding:"required" validate:"email" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}
type UserCreateDTOS struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"name" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
}
