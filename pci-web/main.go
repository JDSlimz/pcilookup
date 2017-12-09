package main

import(
    "html/template"
    "net/http"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))//home/josh/go/src/pci-web/
func main(){
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))//home/josh/go/src/pci-web/
    http.HandleFunc("/contribute.json", contribute)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8001", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Security-Policy:", "default-src https:")

    action := r.URL.Query().Get("action")
    if len(action) == 0{
        err := templates.ExecuteTemplate(w, "indexPage", nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    } else if action == "message" {
        //Send email then
        name := r.URL.Query().Get("name")
        email := r.URL.Query().Get("email")
        message := r.URL.Query().Get("message")

        mail(name, email, message)

        err := templates.ExecuteTemplate(w, "indexPage", nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    } else {
        err := templates.ExecuteTemplate(w, "resultsPage", nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func contribute(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "contribute.json")
}