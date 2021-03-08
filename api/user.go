package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/omaciel/gobutsu/db/sqlc"
)

type createUserRequest struct {
	Username       string `json:"username" binding:"required,alphanum"`
	HashedPassword string `json:"hashed_password" binding:"required,min=6"`
	Email          string `json:"email" binding:"required"`
}

type userResponse struct {
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	CreatedOn      time.Time `json:"created_on"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Email:          user.Email,
		CreatedOn:      user.CreatedOn,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.HashedPassword,
		Email:          req.Email,
	}

	user, err := server.app.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}
