package gardenwars

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check that the user exists
	user, err := svc.UsersRepo.GetByEmail(c.Request.Context(), userLoginReq.Email)
	if err == nil {
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusBadRequest, gin.H{"error": "User does not exist"})
				return
			}

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// check that the password is correct
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(userLoginReq.Password),
	); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, user)
}
