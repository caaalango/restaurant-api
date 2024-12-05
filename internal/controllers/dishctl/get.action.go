package dishctl

import (
	"net/http"

	getdishcmd "github.com/calango-productions/api/internal/core/use-cases/dish/get"
	"github.com/calango-productions/api/internal/types"
	"github.com/calango-productions/api/pkg/dtohdls"
	"github.com/gin-gonic/gin"
)

type GetParamsDto struct {
	DishToken string `json:"dishToken"`
}

func (dc *DishController) GetAction(ctx *gin.Context) {
	userData, err := dtohdls.GetUserData(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "invalid token type"})
		return
	}

	params, err := dtohdls.GetParams[GetParamsDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	executeParams := dc.parseGetParams(params, userData)
	command := getdishcmd.New(dc.DishRepository, dc.RedisRepository)

	result, err := command.Execute(executeParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}
func (dc *DishController) parseGetParams(params *GetParamsDto, context *types.UserData) getdishcmd.Params {
	return getdishcmd.Params{
		Middle: getdishcmd.MiddleParams{
			ClientToken: context.ClientToken,
		},
		Core: getdishcmd.CoreParams{
			DishToken: params.DishToken,
		}}
}
