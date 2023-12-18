//go:build wireinject
// +build wireinject

package wire

import (
	"feature/internal/handler"
	"feature/internal/repository"
	"feature/internal/server"
	"feature/internal/service"
	"feature/pkg/app"
	"feature/pkg/helper/sid"
	"feature/pkg/jwt"
	"feature/pkg/log"
	"feature/pkg/server/http"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewPatientRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewPatientService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewPatientHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
	server.NewTask,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
