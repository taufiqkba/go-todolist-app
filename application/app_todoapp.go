package application

import (
	"github.com/taufiqkba/go-todolist-app/domain_todocore/controller/restapi"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/gateway/withgorm"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/usecase/getalltodo"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/usecase/runtodocheck"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/usecase/runtodocreate"
	"github.com/taufiqkba/go-todolist-app/shared/config"
	"github.com/taufiqkba/go-todolist-app/shared/gogen"
	"github.com/taufiqkba/go-todolist-app/shared/infrastructure/logger"
	"github.com/taufiqkba/go-todolist-app/shared/infrastructure/token"
)

type todoapp struct{}

func NewTodoapp() gogen.Runner {
	return &todoapp{}
}

func (todoapp) Run() error {

	const appName = "todoapp"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	primaryDriver := restapi.NewController(appData, log, cfg, jwtToken)

	primaryDriver.AddUsecase(
		//
		getalltodo.NewUsecase(datasource),
		runtodocheck.NewUsecase(datasource),
		runtodocreate.NewUsecase(datasource),
	)

	primaryDriver.RegisterRouter()

	primaryDriver.Start()

	return nil
}
