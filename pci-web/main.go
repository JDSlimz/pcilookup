package main

import(
    "html/template"
    "net/http"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("/home/josh/go/src/pci-web/templates/*"))

func main(){
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("/home/josh/go/src/pci-web/css/"))))
	http.HandleFunc("/", IndexHandler)
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", fs)
	http.ListenAndServe(":8001", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

    action := r.URL.Query().Get("action")
    if len(action) == 0 {
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
