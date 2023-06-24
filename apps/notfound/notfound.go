package notfound

import (
	"html/template"
	"log"
	"net/http"

	"github.com/webteleport/webteleport/ufo"
	"k0s.io/pkg/middleware"
)

func Arg0(args []string, fallback string) string {
	if len(args) > 0 {
		return args[0]
	}
	return fallback
}

const TEMPLATE = `<!DOCTYPE html>
<html lang="en" charset="utf-8">
<head>
	<title>{{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
</head>
<body>
	<h2>{{.Title}}</h2>
	<p>{{.Message}}</p>
	<code>{{.Code}}</code>
	<p>
	  <a href="{.Link}">webteleport/ufo</a>
	</p>
</body>
</html>
`

type Data struct {
	Title   string
	Message string
	Code    string
	Link    string
}

func NotFoundHandler() http.Handler {
	tmpl, err := template.New("404").Parse(TEMPLATE)
	if err != nil {
		log.Fatalln(err)
	}
	data := Data{
		Title:   "🙈 host not found",
		Message: `You can teleport your local app to this site. Try:`,
		Code:    `$ ufo teleport https://ufo.k0s.io http://127.0.0.1:3000`,
		Link:    "https://github.com/webteleport/ufo",
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
}

func Run(args []string) error {
	return ufo.Serve(Arg0(args, "https://ufo.k0s.io"), middleware.LoggingMiddleware(NotFoundHandler()))
}
