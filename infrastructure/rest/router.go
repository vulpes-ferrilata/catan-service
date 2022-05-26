package rest

import (
	"net/http"

	infrastructure_context "github.com/VulpesFerrilata/catan-service/infrastructure/context"
	"github.com/VulpesFerrilata/catan-service/infrastructure/middlewares"
	"github.com/VulpesFerrilata/catan-service/presentation/rest/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
)

type Router interface {
	Init(app *iris.Application)
}

func NewRouter(translatorMiddleware *middlewares.TranslatorMiddleware,
	authenticationMiddleware *middlewares.AuthenticationMiddleware,
	errorHandlerMiddleware *middlewares.ErrorHandlerMiddleware,
	catanController controllers.CatanController,
	catanWebsocketController controllers.CatanWebsocketController) Router {
	return &router{
		translatorMiddleware:     translatorMiddleware,
		authenticationMiddleware: authenticationMiddleware,
		errorHandlerMiddleware:   errorHandlerMiddleware,
		catanController:          catanController,
		catanWebsocketController: catanWebsocketController,
	}
}

type router struct {
	translatorMiddleware     *middlewares.TranslatorMiddleware
	authenticationMiddleware *middlewares.AuthenticationMiddleware
	errorHandlerMiddleware   *middlewares.ErrorHandlerMiddleware
	catanController          controllers.CatanController
	catanWebsocketController controllers.CatanWebsocketController
}

func (r router) Init(app *iris.Application) {
	api := app.Party("/api")

	catanApi := api.Party("/catan")
	catanApi.Use(r.translatorMiddleware.Serve)
	catanApi.Use(r.authenticationMiddleware.Serve)
	catanMvc := mvc.New(catanApi)
	catanMvc.HandleError(r.errorHandlerMiddleware.Handle)
	catanMvc.Handle(r.catanController)
	catanMvc.HandleWebsocket(r.catanWebsocketController)
	catanWs := websocket.New(websocket.DefaultGorillaUpgrader, catanMvc)
	catanWs.IDGenerator = func(w http.ResponseWriter, r *http.Request) string {
		ctx := r.Context()
		userID, err := infrastructure_context.GetUserID(ctx)
		if err != nil {
			return neffos.DefaultIDGenerator(w, r)
		}

		return userID
	}
	catanApi.Get("/ws", websocket.Handler(catanWs))
}
