package menuctl

import (
	"encoding/json"
	"fmt"
	"net/http"

	showmenucmd "github.com/calango-productions/api/internal/core/use-cases/menu/show"
	"github.com/calango-productions/api/pkg/dtohdls"
	"github.com/gin-gonic/gin"
)

type GetParamsDto struct {
	Name string
}

func (mc *MenuController) ShowAction(ctx *gin.Context) {
	params, err := dtohdls.GetParams[GetParamsDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	executeParams := mc.parseGetParams(params)
	command := showmenucmd.New(
		mc.RestaurantRepository,
		mc.DishRepository,
		mc.CommentRepository,
		mc.RedisRepository,
	)

	result, err := command.Execute(executeParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = json.Marshal(result.Data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to marshal JSON: %+v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}
func (mc *MenuController) parseGetParams(params *GetParamsDto) showmenucmd.Params {
	return showmenucmd.Params{
		Core: showmenucmd.CoreParams{
			Name: params.Name,
		}}
}
