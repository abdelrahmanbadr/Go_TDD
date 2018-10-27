package registry

import (
	"github.com/sarulabs/di"
	"user_tdd/common"
	"user_tdd/config"
	"user_tdd/user/repository"
	"user_tdd/user/service"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "user-service",
			Build: buildUserUsecase,
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	config := common.DbConfig{DbHost: config.MONGODB_HOST, DbName: config.MONGODB_DATABASE_NAME}
	mongoConnection, _ := common.GetMongoDbConnection(config)
	repo := repository.NewMongoUserRepo(mongoConnection)
	service := service.NewUserService(repo)
	return service, nil
}
