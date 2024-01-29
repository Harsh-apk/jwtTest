package api

import (
	"github.com/Harsh-apk/jwtTest/db"
	"github.com/Harsh-apk/jwtTest/types"
	"github.com/Harsh-apk/jwtTest/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		UserStore: userStore,
	}
}

func (u *UserHandler) HandleLogin(c *fiber.Ctx) error {
	var inUser types.IncomingUser
	err := c.BodyParser(&inUser)
	if err != nil {
		return err
	}
	user, err := u.UserStore.LoginUser(&inUser, c.Context())
	if err != nil {
		return err
	}
	tokenStr, err := utils.CreateToken(user)
	if err != nil {
		return err
	}
	cookie := fiber.Cookie{
		Path:     "/",
		Name:     "jwtTest",
		Value:    *tokenStr,
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Secure:   true,
	}
	c.Cookie(&cookie)
	return c.JSON(user)

}
func (u *UserHandler) HandleCreateAccount(c *fiber.Ctx) error {
	var inUser types.IncomingUser
	err := c.BodyParser(&inUser)
	if err != nil {
		return err
	}
	User, err := utils.CreateUserFromIncomingUser(&inUser)
	if err != nil {
		return err
	}
	err = u.UserStore.CreateUser(User, c.Context())
	if err != nil {
		return err
	}
	return c.JSON(User)
}
func (u *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Context().UserValue("id")
	oid, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}
	user, err := u.UserStore.GetUserById(&oid, c.Context())
	if err != nil {
		return err
	}
	return c.JSON(&user)
}
