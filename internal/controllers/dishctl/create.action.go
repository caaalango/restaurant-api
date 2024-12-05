package dishctl

import (
	"net/http"

	createdishcmd "github.com/calango-productions/api/internal/core/use-cases/dish/create"
	"github.com/calango-productions/api/internal/types"
	"github.com/calango-productions/api/pkg/dtohdls"
	"github.com/calango-productions/api/pkg/validator"
	"github.com/gin-gonic/gin"
)

type CreateBodyDto struct {
	Title        string   `json:"title" validate:"required,min=1,max=40"`
	Description  string   `json:"description" validate:"required,min=10,max=100"`
	Restrictions []string `json:"restrictions"`
	Price        int      `json:"price" validate:"required,gt=0"`
}

func (dc *DishController) CreateAction(ctx *gin.Context) {
	userData, err := dtohdls.GetUserData(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "invalid token type"})
		return
	}

	err = validator.ValidateBody[CreateBodyDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	body, err := dtohdls.GetBody[CreateBodyDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid request body"})
		return
	}

	executeParams := dc.parseCreateParams(body, userData)

	command := createdishcmd.New(dc.DishRepository)
	result, err := command.Execute(executeParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}

func (dc *DishController) parseCreateParams(body *CreateBodyDto, context *types.UserData) createdishcmd.Params {
	return createdishcmd.Params{
		Middle: createdishcmd.MiddleParams{
			ClientToken: context.ClientToken,
		},
		Core: createdishcmd.CoreParams{
			Title:        body.Title,
			Description:  body.Description,
			Restrictions: body.Restrictions,
			Price:        body.Price,
		}}
}
