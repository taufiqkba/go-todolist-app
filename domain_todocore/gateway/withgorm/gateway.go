package withgorm

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/taufiqkba/go-todolist-app/domain_todocore/model/entity"
	"github.com/taufiqkba/go-todolist-app/domain_todocore/model/vo"
	"github.com/taufiqkba/go-todolist-app/shared/config"
	"github.com/taufiqkba/go-todolist-app/shared/gogen"
	"github.com/taufiqkba/go-todolist-app/shared/infrastructure/logger"
)

type gateway struct {
	appData gogen.ApplicationData
	config  *config.Config
	log     logger.Logger
	db      *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	dsn := "root:root@tcp(127.0.0.1:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	err = db.AutoMigrate(entity.Todo{})
	if err != nil {
		panic(err)
	}

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		db:      db,
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called-FindAllTodo")

	var todoObjs []*entity.Todo
	var count int64

	err := r.db.
		Model(entity.Todo{}).
		Count(&count).
		Limit(size).
		Offset((page - 1) * size).
		Find(&todoObjs).Error
	if err != nil {
		return nil, 0, err
	}
	return todoObjs, count, nil
}

func (r *gateway) FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called-FindOneTodoByID")

	var todoObj entity.Todo

	err := r.db.First(&todoObj, "id = ?", todoID).Error
	if err != nil {
		return nil, err
	}
	return &todoObj, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called-SaveTodo")
	err := r.db.Save(obj).Error
	if err != nil {
		return err
	}
	return nil
}
