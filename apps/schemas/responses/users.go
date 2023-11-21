package responses

import (
	"time"

	"github.com/JngMkk/golang-fiber/apps/models"
)

type UserResp struct {
	Email     string    `json:"email"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUserResponse(u *models.User) *UserResp {
	res := new(UserResp)

	res.CreatedAt = u.CreatedAt
	res.Email = u.Email
	res.IsActive = u.IsActive
	return res
}
