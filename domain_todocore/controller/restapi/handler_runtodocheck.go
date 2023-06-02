package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/model/entity"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/model/vo"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/usecase/runtodocheck"
	"github.com/taufiqkba/go-todolist-app/shared/gogen"
	"github.com/taufiqkba/go-todolist-app/shared/infrastructure/logger"
	"github.com/taufiqkba/go-todolist-app/shared/model/payload"
	"github.com/taufiqkba/go-todolist-app/shared/util"
)

func (r *controller) runTodoCheckHandler() gin.HandlerFunc {

	type InportRequest = runtodocheck.InportRequest
	type InportResponse = runtodocheck.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		TodoID vo.TodoID `uri:"todo_id"`
	}

	type response struct {
		Todo *entity.Todo `json:"todo"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		err := c.BindUri(&jsonReq)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.TodoID = jsonReq.TodoID

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Todo = res.Todo

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
