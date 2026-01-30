package handlers

import (
	"net/http"

	"github.com/TheTeemka/GoProjects/hw_6/models"
	"github.com/TheTeemka/GoProjects/hw_6/utils"
	"github.com/labstack/echo/v4"
)

type IUserService interface {
	CreateUser(dto *models.CreateUserRequest) error
	GetUserByEmail(email string) (*models.UserDTO, error)
	Login(email, plainPassword string) (accessToken string, refreshToken string, err error)
	RefreshAccessToken(refreshToken string) (string, error)
}

type ITokenService interface {
	ValitateToken(userID int, token string) (bool, error)
}

type AuthHandler struct {
	userService IUserService
}

func NewUserHandler(userService IUserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "Create user request"
// @Success 201 {string} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func (uh *AuthHandler) CreateUser(c echo.Context) error {
	var req models.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := uh.userService.CreateUser(&req); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.NoContent(201)
}

// GetMe godoc
// @Summary Get current authenticated user
// @Tags User
// @Produce json
// @Success 200 {object} models.UserDTO
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /auth/users/me [get]
func (uh *AuthHandler) GetMe(c echo.Context) error {
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

// Login godoc
// @Summary Login user and return access token
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.LoginUserRequest true "Login request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/user/login [get]
func (uh *AuthHandler) Login(c echo.Context) error {
	var req models.LoginUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	accessToken, refreshToken, err := uh.userService.Login(req.Email, req.PlainPassword)

	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
	})

	return c.JSON(200, map[string]string{"access_token": accessToken})
}

func (uh *AuthHandler) RefreshAccessToken(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(401, map[string]string{"error": "unauthorized"})
	}
	refreshToken := cookie.Value

	accessToken, err := uh.userService.RefreshAccessToken(refreshToken)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{"access_token": accessToken})
}
