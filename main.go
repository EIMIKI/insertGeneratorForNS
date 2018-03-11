package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	template *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t.once.Do(func() {
		t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	if r.Form.Get("nation_name") != "" {
		reqData := formData{}
		reqData.setFromData(r.Form)
		createSql(&reqData)
	}
	t.template.Execute(w, nil)
}
func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
