package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/usecase/getalltodo"
	"github.com/taufiqkba/go-todolist-app/shared/gogen"
	"github.com/taufiqkba/go-todolist-app/shared/infrastructure/logger"
	"github.com/taufiqkba/go-todolist-app/shared/model/payload"
	"github.com/taufiqkba/go-todolist-app/shared/util"
)

func (r *controller) getAllTodoHandler() gin.HandlerFunc {

	type InportRequest = getalltodo.InportRequest
	type InportResponse = getalltodo.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		Page int `form:"page,omitempty,default=0"`
		Size int `form:"size,omitempty,default=0"`
	}

	type response struct {
		Count int64 `json:"count"`
		Items []any `json:"items"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.Bind(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.Page = jsonReq.Page
		req.Size = jsonReq.Size

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Count = res.Count
		jsonRes.Items = res.Items

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
