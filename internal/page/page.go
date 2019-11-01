package page

import (
	"html/template"
	"net/http"
	"io/ioutil"

	"github.com/julienschmidt/httprouter"
)

var templates *template.Template

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func ParseFiles() {
	files := []string{
		"tmpl/edit.html",
		"tmpl/view.html",
		"tmpl/viewTitle.html",
		"tmpl/home.html",
		"tmpl/login.html",
	}
	templates = template.Must(template.ParseFiles(files...))
}

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, "home", nil)
	}
}

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, "login", nil)
	}
}

func PostView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		pwd := r.FormValue("pwd")
		if user != "bob" && pwd != "dylan" {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name: "token",
			Value: "secret",
			HttpOnly: true,
			SameSite: http.SameSiteDefaultMode,
			// the Secure field requires TLS
		})
		render(w, "view", nil)
	}
}

func GetView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, "view", nil)
	}
}

func ViewTitle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := param("title", r)
		p, err := loadPage(title)
		if err != nil {
			http.Redirect(w, r, "/edit/"+title, http.StatusFound)
			return
		}
		render(w, "viewTitle", p)
	}
}

func EditTitle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := param("title", r)
		p, err := loadPage(title)
		if err != nil {
			p = &Page{Title: title} // create a page on the fly
		}
		render(w, "edit", p)
	}
}

func SaveTitle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := param("title", r)
		body := r.FormValue("body")
		p := &Page{Title: title, Body: []byte(body)}
		err := p.save()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	}
}

// render renders a template on w
func render(w http.ResponseWriter, name string, p *Page) {
	err := templates.ExecuteTemplate(w, name+".html", p)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// param returns a named parameter from the request context
func param(s string, r *http.Request) string {
	return httprouter.ParamsFromContext(r.Context()).ByName(s)
}
