package dishctl

import (
	"net/http"

	inactivatedishcmd "github.com/calango-productions/api/internal/core/use-cases/dish/inactivate"
	"github.com/calango-productions/api/internal/types"
	"github.com/calango-productions/api/pkg/dtohdls"
	"github.com/gin-gonic/gin"
)

type InactivateParamsDto struct {
	DishToken string `json:"dishToken"`
}

func (dc *DishController) InactivateAction(ctx *gin.Context) {
	userData, err := dtohdls.GetUserData(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "invalid token type"})
		return
	}

	params, err := dtohdls.GetParams[InactivateParamsDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	executeParams := dc.parseInactivateParams(params, userData)
	command := inactivatedishcmd.New(dc.DishRepository)
	result, err := command.Execute(executeParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}
func (dc *DishController) parseInactivateParams(params *InactivateParamsDto, context *types.UserData) inactivatedishcmd.Params {
	return inactivatedishcmd.Params{
		Middle: inactivatedishcmd.MiddleParams{
			ClientToken: context.ClientToken,
		},
		Core: inactivatedishcmd.CoreParams{
			DishToken: params.DishToken,
		}}
}
