package story

import (
	"html/template"
	"net/http"
	"strings"

	jp "github.com/Its-Maniaco/AdventureGame/JsonParse"
)

func init() {
	// os.Chdir("Template")
	// tmplfile := "story.html"
	tmpl = template.Must(template.New("").Parse(defaultTemplate))

}

var tmpl *template.Template

func NewStoryHandler(s jp.Story) http.Handler {
	return storyHandler{s: s}
}

type storyHandler struct {
	s jp.Story
}

// storyHandler has a method named ServeHTTP, so it satisfies the http.Handler interface, instance of storyHandler passed to http.ListenAndServe, calls the ServeHTTP method internally.
func (sh storyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimSpace(req.URL.Path)

	// if there is nothing after base url or only a slash -> intro
	if path == "" || path == "/" {
		path = "/intro"
	}

	// get Chapter name
	path = path[1:]
	if chapter, ok := sh.s[path]; ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			panic(err)
		}
	}

}

var defaultTemplate = `
<!DOCTYPE html>
<html>

<head>
    <title> Choose Your Own Adventure </title>
</head>

<body>
    {{.Title}}
    {{range .Story}}
    <p>{{.}}</p>
    {{end}}

    <ul>
        {{range .Options}}
        <li> <a href="./{{.Arc}}">{{.Text}}</a> </li>
		{{end}}
    </ul>

</body>

</html>`
