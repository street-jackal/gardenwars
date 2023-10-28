package gardenwars

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/street-jackal/gardenwars/service/types/responses"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (svc *Service) LoginUser(c *gin.Context) {
	type userLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	userLoginReq := &userLoginRequest{}

	if err := c.ShouldBindJSON(userLoginReq); err != nil {
		c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Status:  http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})

		return
	}

	// check that the user exists
	user, err := svc.UsersRepo.GetByEmail(c.Request.Context(), userLoginReq.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, responses.BaseResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "User does not exist",
			})

			return
		}

		c.JSON(http.StatusBadRequest, responses.BaseResponse{
			Status:  http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})

		return
	}

	// check that the password is correct
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(userLoginReq.Password),
	); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusBadRequest, responses.BaseResponse{
				Status:  http.StatusBadRequest,
				Message: "Incorrect password",
				Success: false,
			})

			return
		}

		c.JSON(http.StatusInternalServerError, responses.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Success: false,
		})

		return
	}

	c.JSONP(http.StatusOK, user)
}
