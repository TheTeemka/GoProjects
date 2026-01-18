package handlers

import (
	"github.com/TheTeemka/GoProjects/hw_5/models"
	"github.com/TheTeemka/GoProjects/hw_5/utils"
	"github.com/labstack/echo"
)

type IUserService interface {
	CreateUser(dto *models.CreateUserRequest) error
	GetUserByEmail(email string) (*models.UserDTO, error)
}

type IJWTService interface {
	CreateToken(userID int, email string) (string, error)
	ParseToken(tokenStr string) (*models.UserClaims, error)
}

type UserHandler struct {
	userService IUserService
	jwtService  IJWTService
}

func NewUserHandler(userService IUserService, jwtService IJWTService) *UserHandler {
	return &UserHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	var req models.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := uh.userService.CreateUser(&req); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.NoContent(201)
}

func (uh *UserHandler) GetMe(c echo.Context) error {
	userClaims, ok := utils.GetUserClaims(c)
	if !ok {
		return c.JSON(401, map[string]string{"error": "unauthorized"})
	}
	email := userClaims.Email

	userDTO, err := uh.userService.GetUserByEmail(email)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	if userDTO == nil {
		return c.JSON(404, map[string]string{"error": "User not found"})
	}

	return c.JSON(200, userDTO)
}

func (uh *UserHandler) Login(c echo.Context) error {
	var req models.LoginUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	userDTO, err := uh.userService.GetUserByEmail(req.Email)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	if userDTO == nil {
		return c.JSON(401, map[string]string{"error": "invalid email or password"})
	}

	if err := utils.ComparePassword(userDTO.PasswordHash, []byte(req.PlainPassword)); err != nil {
		return c.JSON(401, map[string]string{"error": "invalid email or password"})
	}

	token, err := uh.jwtService.CreateToken(userDTO.ID, userDTO.Email)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{"access_token": token})
}
