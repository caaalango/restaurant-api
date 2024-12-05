package dishctl

import (
	"net/http"

	updatedishcmd "github.com/calango-productions/api/internal/core/use-cases/dish/update"
	"github.com/calango-productions/api/internal/types"
	"github.com/calango-productions/api/pkg/dtohdls"
	"github.com/calango-productions/api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type UpdateBodyDto struct {
	Description  string   `json:"description" validate:"required,min=10,max=100"`
	Restrictions []string `json:"restrictions"`
	Price        int      `json:"price" validate:"required,gt=0"`
}

type UpdateParamsDto struct {
	DishToken string
}

func (dc *DishController) UpdateAction(ctx *gin.Context) {
	userData, err := dtohdls.GetUserData(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "invalid user data"})
		return
	}

	err = validator.ValidateBody[CreateBodyDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	body, err := dtohdls.GetBody[UpdateBodyDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}

	params, err := dtohdls.GetParams[UpdateParamsDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}

	executeParams := dc.parseUpdateParams(params, body, userData)

	command := updatedishcmd.New(dc.DishRepository)
	result, err := command.Execute(executeParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}

func (dc *DishController) parseUpdateParams(params *UpdateParamsDto, body *UpdateBodyDto, context *types.UserData) updatedishcmd.Params {
	return updatedishcmd.Params{
		Middle: updatedishcmd.MiddleParams{
			ClientToken: context.ClientToken,
		},
		Core: updatedishcmd.CoreParams{
			DishToken:    params.DishToken,
			Description:  body.Description,
			Restrictions: body.Restrictions,
			Price:        body.Price,
		}}
}
