package user

import (
	"clean-api/database/user"
	"clean-api/model"
	"clean-api/pkg/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	repository user.Repository
}

func NewController(repository user.Repository) *Controller {
	return &Controller{repository: repository}
}

func NewHandlers(instance *echo.Echo, r *Controller) {

	grp := instance.Group("/auth/")
	grp.POST("login", r.Login)
	grp.POST("register", r.Register)
	grp.GET("me", r.Me)
}

func (r Controller) Login(c echo.Context) error {
	Email := c.FormValue("email")
	Password := c.FormValue("password")

	usrLogin := model.User{
		Email:    Email,
		Password: Password,
	}

	err := bcrypt.CompareHashAndPassword([]byte(usrLogin.Password), []byte(usrLogin.Password))
	if err != nil {
		return c.JSON(401, map[string]string{"message": "Invalid credentials"})
	}

	token, err := jwt.MakeToken(usrLogin.ID)
	if err != nil {
		return c.JSON(500, "Server Error")
	}
	return c.JSON(200, map[string]string{"token": token})
}

func (r Controller) Register(c echo.Context) error {
	Email := c.FormValue("email")
	Password := c.FormValue("password")

	reg := new(model.User)
	reg.Email = Email
	reg.Password = Password
	reg.ID = uuid.NewString()

	// Return user as json
	token, err := jwt.MakeToken(reg.ID)
	if err != nil {
		return c.JSON(500, "Server Error")
	}
	return c.JSON(200, map[string]string{"token": token})

}

// Me godoc
// @Summary Me
// @Description User Me
// @Tags Auth
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router /auth/me [get]
func (r Controller) Me(c echo.Context) error {
	// Get user id from jwt and check it
	userId, _ := jwt.JWTGetUserFromToken("")
	if len(userId.ID) == 0 {
		return c.JSON(401, map[string]string{"message": "Invalid credentials"})
	}

	return c.JSON(200, nil)
}
