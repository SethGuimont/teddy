package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"
)

type PageData struct {
	Title, MetaDescription string
	Year                   int
}

func write(t *template.Template, tplName, out string, pd PageData) {
	dst := filepath.Join("dist", out)
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := t.ExecuteTemplate(f, tplName, pd); err != nil {
		log.Fatal(err)
	}
}
func main() {

	t := template.Must(template.ParseGlob("templates/_*.html"))
	t = template.Must(t.ParseGlob("templates/*.html")) // or *.gohtml
	y := time.Now().Year()
	// Home + sections (match your actual filenames)
	write(t, "index.html", "index.html", PageData{"Home","Explore our prefinishing services.", y})
	// Contact & Thank You pages (static)
	write(t, "guestbook.html", "guestbook/index.html", PageData{"guestbook.","Explore our prefinishing services.", y})
	// a simple thanks page you create at templates/thanks.html
	write(t, "gallery.html", "gallery/index.html", PageData{"Gallery.", "Explore our prefinishing services.",y})
}
