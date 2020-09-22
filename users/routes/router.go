package routes

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"
	"user/users/controllers"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// Prevent abnormal shutdown while panic
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				log.Print(string(debug.Stack()))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func wrapHandler(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

//RouterConfig config routes.
func RouterConfig() (router *httprouter.Router) {
	indexHandlers := alice.New(recoverHandler)
	router = httprouter.New()

	router.POST("/user", wrapHandler(indexHandlers.ThenFunc(controllers.SaveUser)))
	router.PUT("/user/:id", wrapHandler(indexHandlers.ThenFunc(controllers.UpdateUser)))
	router.GET("/user/:id", wrapHandler(indexHandlers.ThenFunc(controllers.GetUser)))
	router.DELETE("/user/:id", wrapHandler(indexHandlers.ThenFunc(controllers.DeleteUser)))
	router.GET("/user/:id/search", wrapHandler(indexHandlers.ThenFunc(controllers.GetUserDetails)))

	return
}
