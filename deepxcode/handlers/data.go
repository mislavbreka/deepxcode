package handlers

import "os"
import "time"
// import "fmt"
import "log"
import "strings"
import "path/filepath"
import "net/http"
import "html/template"
import "github.com/gorilla/mux"

var templates = ParseTemplates()

// ParseTemplates ahja
func ParseTemplates() *template.Template {
    templ := template.New("")

    err := filepath.Walk("./data/html/", func(path string, info os.FileInfo, err error) error {
        if strings.Contains(path, ".gohtml") {
            _, err = templ.ParseFiles(path)
            if err != nil {
                log.Println(err)
            }
        }
        return err
    })

    if err != nil {
        panic(err)
    }

    return templ
}

// Handle jfask
func Handle(router *mux.Router) {
    router.HandleFunc("/", landingHandler)
    router.HandleFunc("/login/", loginHandler)
    router.HandleFunc("/register/", registerHandler)
    router.HandleFunc("/profile/", profileHandler)
    router.HandleFunc("/data/", dataHandler)
	router.PathPrefix("/data/").Handler(http.StripPrefix("/data/", http.FileServer(http.Dir("./data"))))
}

func landingHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("username")
    if err != nil {
        cookie = &http.Cookie{Name: "username", Value: "test", Path: "/", Expires: time.Unix(0, 0)}
        http.SetCookie(w, cookie)
    }
    err = templates.ExecuteTemplate(w, "landing-page.gohtml", cookie.Value)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "login-page.gohtml", nil)
    if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
func registerHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "register-page.gohtml", nil)
    if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "profile-page.gohtml", nil)
    if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
    err := templates.ExecuteTemplate(w, "data-page.gohtml", nil)
    if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

