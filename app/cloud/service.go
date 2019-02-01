package cloud

import (
	"context"
	app2 "github.com/gofunct/stencil/app"

	"log"

	"github.com/gorilla/mux"
)

func main() {

	ctx := context.Background()
	var app *app2.App
	app = app2.Initialize(app)
	var api = new(Api)
	var cleanup func()
	var err error
	switch app.Values.Env {
	case "gcp":
		api, cleanup, err = setupGCP(ctx)
	case "aws":
		if app.Values.DbPassword == "" {
			app.Values.DbPassword = "xyzzy"
		}
		api, cleanup, err = setupAWS(ctx)
	case "local":
		if app.Values.DbHost == "" {
			app.Values.DbHost = "localhost"
		}
		if app.Values.DbPassword == "" {
			app.Values.DbPassword = "xyzzy"
		}
		api, cleanup, err = setupLocal(ctx)
	default:
		log.Fatalf("unknown -env=%s", app.Values.Env)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	// Set up URL routes.
	r := mux.NewRouter()

	r.HandleFunc("/", api.index)
	r.HandleFunc("/sign", api.sign)
	r.HandleFunc("/blob/{key:.+}", api.serveBlob)

	// Listen and serve HTTP.
	log.Printf("Running, connected to %q cloud", app.Values.Env)
	log.Fatal(api.Server.Srv.ListenAndServe(app.Values.Port, r))
}
