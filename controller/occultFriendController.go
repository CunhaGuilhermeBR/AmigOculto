package controller

import (
	"net/http"
	"src/amigOculto/models"
	"strings"
	"text/template"
)

const ROUTEOK = 301

type Pick struct {
	Pick string
	Id   string
}

var templates = template.Must(template.ParseGlob("public/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index", nil)
}

func CreateOccultFriend(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Create", nil)
}

func DrawFriend(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Draw", nil)
}

func Faq(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Index2", nil)
}

func GetPick(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idOc := r.FormValue("idOccultFriend")
		pick := models.GetRandomNotDraw(idOc)
		p := Pick{pick, idOc}
		templates.ExecuteTemplate(w, "ShowPick", p)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		models.CreateNew(r.FormValue("participants"))
	}
	http.Redirect(w, r, "/", ROUTEOK)
}

func Confirme(w http.ResponseWriter, r *http.Request) {
	idPick := r.URL.Query().Get("pick")
	idPickSlice := strings.Split(idPick, "?")
	models.RemoveNotDraw(idPickSlice[1], idPickSlice[0])
	http.Redirect(w, r, "/", ROUTEOK)
}
