package registry

import (
	"github.com/sarulabs/di"
	"user_tdd/user/repository"
	"user_tdd/user/service"
)

type Container struct {
	di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}
	err = builder.Add(di.Def{
		Name:  "user-service",
		Build: buildUserService,
	})
	if err != nil {
		return nil, err
	}

	return &Container{builder.Build()}, nil
}

func buildUserService(ctn di.Container) (interface{}, error) {
	//config := common.DbConfig{DbHost: config.MONGODB_HOST, DbName: config.MONGODB_DATABASE_NAME}
	//mongoConnection, _ := common.GetMongoDbConnection(config)
	//mongoRepo := repository.NewMongoUserRepo(mongoConnection)
	inMemoryRepo := repository.NewInMemoryRepository()
	service := service.NewUserService(inMemoryRepo)
	return service, nil
}
