package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/VadimBoganov/fulgur/internal/domain"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func (h *Handler) Register(c *gin.Context) {
	var data map[string]interface{}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		logger.Errorf("Error while read user request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), 14)

	user := domain.User{
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Password: password,
	}

	id, err := h.service.User.Add(&user)
	if err != nil {
		logger.Errorf("Error while add user: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}
	user.Id = uint(id)

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c *gin.Context) {
	var data map[string]interface{}
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		logger.Errorf("Error while read user request: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	user, err := h.service.GetByEmail(data["email"].(string))
	if err != nil {
		logger.Errorf("Error while get user by email: %s", err.Error())
		_ = c.AbortWithError(400, err)
		return
	}

	if user.Id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, "user not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"].(string))); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "incorect password")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: &jwt.Time{
			Time: time.Now().Add(time.Hour * 24),
		},
	})

	err = godotenv.Load(".env")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "could not read env file")
		return
	}

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "could not login")
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusAccepted, "success")
}

func (h *Handler) Logout(c *gin.Context) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusAccepted, "success")
}
