package cloud

import (
	"bytes"
	"github.com/gofunct/stencil/app"
	"github.com/gofunct/stencil/app/server"
	"github.com/gorilla/mux"
	"gocloud.dev/runtimevar"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"context"
)

type Api struct {
	App *app.App
	Server *server.Server
	mu sync.RWMutex
}

func NewApi(a *app.App, s *server.Server) *Api {
	return &Api{
		mu: sync.RWMutex{},
		App:    a,
		Server: s,
	}
}


// index serves the server's landing page. It lists the 100 most recent
// greetings, shows a cloud environment banner, and displays the message of the
// day.
func (app *Api) index(w http.ResponseWriter, r *http.Request) {
	var data struct {
		MOTD      string
		Env       string
		BannerSrc string
		Greetings []greeting
	}
	app.mu.RLock()
	data.MOTD = app.App.Values.RunVar
	app.mu.RUnlock()
	switch app.App.Values.Env {
	case "gcp":
		data.Env = "GCP"
		data.BannerSrc = "/blob/gcp.png"
	case "aws":
		data.Env = "AWS"
		data.BannerSrc = "/blob/aws.png"
	case "local":
		data.Env = "Local"
		data.BannerSrc = "/blob/gophers.jpg"
	}

	const query = "SELECT content FROM (SELECT content, post_date FROM greetings ORDER BY post_date DESC LIMIT 100) AS recent_greetings ORDER BY post_date ASC;"
	q, err := app.Server.Db.QueryContext(r.Context(), query)
	if err != nil {
		log.Println("main page SQL error:", err)
		http.Error(w, "could not load greetings", http.StatusInternalServerError)
		return
	}
	defer q.Close()
	for q.Next() {
		var g greeting
		if err := q.Scan(&g.Content); err != nil {
			log.Println("main page SQL error:", err)
			http.Error(w, "could not load greetings", http.StatusInternalServerError)
			return
		}
		data.Greetings = append(data.Greetings, g)
	}
	if err := q.Err(); err != nil {
		log.Println("main page SQL error:", err)
		http.Error(w, "could not load greetings", http.StatusInternalServerError)
		return
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, data); err != nil {
		log.Println("template error:", err)
		http.Error(w, "could not render page", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	if _, err := w.Write(buf.Bytes()); err != nil {
		log.Println("writing response:", err)
	}
}

type greeting struct {
	Content string
}

var tmpl = template.Must(template.New("index.html").Parse(`<!DOCTYPE html>
<title>Guestbook - {{.Env}}</title>
<style type="text/css">
html, body {
	font-family: Helvetica, sans-serif;
}
blockquote {
	font-family: cursive, Helvetica, sans-serif;
}
.banner {
	height: 125px;
	width: 250px;
}
.greeting {
	font-size: 85%;
}
.motd {
	font-weight: bold;
}
</style>
<h1>Guestbook</h1>
<div><img class="banner" src="{{.BannerSrc}}"></div>
{{with .MOTD}}<p class="motd">Admin says: {{.}}</p>{{end}}
{{range .Greetings}}
<div class="greeting">
	Someone wrote:
	<blockquote>{{.Content}}</blockquote>
</div>
{{end}}
<form action="/sign" method="POST">
	<div><textarea name="content" rows="3"></textarea></div>
	<div><input type="submit" value="Sign"></div>
</form>
`))

// sign is a form action handler for adding a greeting.
func (app *Api) sign(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	content := r.FormValue("content")
	if content == "" {
		http.Error(w, "content must not be empty", http.StatusBadRequest)
		return
	}
	const sqlStmt = "INSERT INTO greetings (content) VALUES (?);"
	_, err := app.Server.Db.ExecContext(r.Context(), sqlStmt, content)
	if err != nil {
		log.Println("sign SQL error:", err)
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// serveBlob handles a request for a static asset by retrieving it from a bucket.
func (app *Api) serveBlob(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	blobRead, err := app.Server.Bucket.NewReader(r.Context(), key, nil)
	if err != nil {
		// TODO(light): Distinguish 404.
		// https://github.com/google/go-cloud/issues/2
		log.Println("serve blob:", err)
		http.Error(w, "blob read error", http.StatusInternalServerError)
		return
	}
	// TODO(light): Get content type from blob storage.
	// https://github.com/google/go-cloud/issues/9
	switch {
	case strings.HasSuffix(key, ".png"):
		w.Header().Set("Content-Type", "image/png")
	case strings.HasSuffix(key, ".jpg"):
		w.Header().Set("Content-Type", "image/jpeg")
	default:
		w.Header().Set("Content-Type", "Api/octet-stream")
	}
	w.Header().Set("Content-Length", strconv.FormatInt(blobRead.Size(), 10))
	if _, err = io.Copy(w, blobRead); err != nil {
		log.Println("Copying blob:", err)
	}
}

// watchMOTDVar listens for changes in v and updates the app's message of the
// day. It is run in a separate goroutine.
func (a *Api) watchMOTDVar(v *runtimevar.Variable) {
	ctx := context.Background()
	for {
		snap, err := v.Watch(ctx)
		if err != nil {
			log.Printf("watch MOTD variable: %v", err)
			continue
		}
		log.Println("updated MOTD to", snap.Value)
		a.mu.Lock()
		a.App.Values.RunVar = snap.Value.(string)
		a.mu.Unlock()
	}
}