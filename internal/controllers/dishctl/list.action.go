package dishctl

import (
	"net/http"

	listdishcmd "github.com/calango-productions/api/internal/core/use-cases/dish/list"
	"github.com/calango-productions/api/internal/types"
	"github.com/calango-productions/api/pkg/dtohdls"
	"github.com/gin-gonic/gin"
)

type ListQueryDto struct {
	Page int
	Size int
}

func (dc *DishController) ListAction(ctx *gin.Context) {
	userData, err := dtohdls.GetUserData(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "invalid token type"})
		return
	}

	queries, err := dtohdls.GetQueries[ListQueryDto](ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid queries"})
		return
	}

	executeParams := dc.parseListParams(queries, userData)
	command := listdishcmd.New(dc.DishRepository)

	result, err := command.Execute(executeParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"result": result})
}
func (dc *DishController) parseListParams(queries *ListQueryDto, context *types.UserData) listdishcmd.Params {
	return listdishcmd.Params{
		Middle: listdishcmd.MiddleParams{
			ClientToken: context.ClientToken,
		},
		Core: listdishcmd.CoreParams{
			Page: queries.Page,
			Size: queries.Size,
		},
	}
}
