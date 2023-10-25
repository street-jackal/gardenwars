package gardenwars

import (
	"log/slog"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/street-jackal/gardenwars/repository/models"
)

func (svc *Service) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// return error if user already exists
	if _, err := svc.UsersRepo.GetByEmail(c.Request.Context(), user.Email); err == nil {
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}
	}

	// create a secure hash for the password so we're not saving it in plain text
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	newUser := &models.User{
		ID:       uuid.NewString(),
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	if err := svc.UsersRepo.Insert(c.Request.Context(), newUser); err != nil {
		slog.Error("Failed to insert a User", err)
	}

	c.JSONP(http.StatusOK, user)
}

// hashPassword generates a secure hash from a given string with a default cost factor
func hashPassword(password string) (string, error) {
	// Generate a salt with a cost factor (work factor) to determine the computational complexity.
	cost := 12 // Recommended value (higher is slower but more secure)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
