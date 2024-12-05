package userctl

import (
	"net/http"

	createtradusercmd "github.com/calango-productions/api/internal/core/use-cases/user/createTraditional"
	"github.com/calango-productions/api/internal/types"
	"github.com/calango-productions/api/pkg/dtohdls"
	"github.com/calango-productions/api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type CreateBodyDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type ActorContextDto struct {
	ActorToken string `json:"actorToken"`
}

// @Summary Create user using traditional login.
// @Description Create user using traditional login.
// @Tags users
// @Accept  json
// @Produce  json
// @Param Email query int true "email"
// @Param Password query int true "password"
// @Success 200 {object} map[string]int "Result"
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /users/traditional [post]
func (uc *UserController) CreateTradionalAction(ctx *gin.Context) {
	userData, err := dtohdls.GetUserData(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "invalid actor token type"})
		return
	}

	err = validator.ValidateBody[CreateBodyDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	body, err := dtohdls.GetBody[CreateBodyDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	executeParams := uc.parseCreateParams(body, userData)
	command := createtradusercmd.New(uc.userRepository, uc.credentialRepository)

	result, err := command.Execute(executeParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}

func (uc *UserController) parseCreateParams(body *CreateBodyDto, userData *types.UserData) createtradusercmd.Params {
	return createtradusercmd.Params{
		ActorToken: userData.UserToken,
		Email:      body.Email,
		Password:   body.Password,
	}
}
