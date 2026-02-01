package handlers

import (
	"net/http"

	"github.com/TheTeemka/GoProjects/hw/models"
	"github.com/TheTeemka/GoProjects/hw/utils"
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
func (h *AuthHandler) CreateUser(c echo.Context) error {
	var req models.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	if err := h.userService.CreateUser(&req); err != nil {
		return err
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
func (h *AuthHandler) GetMe(c echo.Context) error {
	userClaims, ok := utils.GetUserClaims(c)
	if !ok {
		return c.JSON(401, map[string]string{"error": "unauthorized"})
	}
	email := userClaims.Email

	userDTO, err := h.userService.GetUserByEmail(email)
	if err != nil {
		return err
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
func (h *AuthHandler) Login(c echo.Context) error {
	var req models.LoginUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	accessToken, refreshToken, err := h.userService.Login(req.Email, req.PlainPassword)

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

func (h *AuthHandler) RefreshAccessToken(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(401, map[string]string{"error": "does not have cookie"})
	}
	refreshToken := cookie.Value

	accessToken, err := h.userService.RefreshAccessToken(refreshToken)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{"access_token": accessToken})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   -1,
	})

	return c.NoContent(200)
}
