package globals
import(
  "html/template"
  "regexp"
  "gopkg.in/mgo.v2"
)

var Templates = template.Must(template.ParseFiles(
  "templates/edit.html", "templates/view.html", "templates/index.html"))

var ValidPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

var Collection *mgo.Collection
