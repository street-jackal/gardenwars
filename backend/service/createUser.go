package gardenwars

import (
	"log/slog"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/street-jackal/gardenwars/repository/models"
	"github.com/street-jackal/gardenwars/service/types/responses"
)

func (svc *Service) CreateUser(c *gin.Context) {
	type createUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	createUserReq := &createUserRequest{}

	if err := c.ShouldBindJSON(createUserReq); err != nil {
		slog.Error("Failed to bind the request body to a struct", err)

		c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		})

		return
	}

	// return error if user already exists
	if _, err := svc.UsersRepo.GetByEmail(c.Request.Context(), createUserReq.Email); err == nil {
		if err == nil {
			slog.Error("User already exists")

			c.JSON(http.StatusBadRequest, responses.BaseResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "User already exists",
			})

			return
		}
	}

	// create a secure hash for the password so we're not saving it in plain text
	hashedPassword, err := hashPassword(createUserReq.Password)
	if err != nil {
		slog.Error("Failed to hash the password", err)

		c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Success: false,
		})

		return
	}

	newUser := &models.User{
		ID:       uuid.NewString(),
		Email:    createUserReq.Email,
		Password: hashedPassword,
	}

	if err := svc.UsersRepo.Insert(c.Request.Context(), newUser); err != nil {
		slog.Error("Failed to insert a User", err)

		c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Success: false,
		})

		return
	}

	c.JSON(
		http.StatusOK,
		responses.BaseResponse{
			Status:  http.StatusOK,
			Message: "User created successfully",
			Success: true,
		},
	)
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
