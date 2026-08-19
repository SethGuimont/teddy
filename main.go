
package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"
)

type GuestbookEntry struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type Photo struct {
	Filename string `json:"filename"`
	Caption  string `json:"caption"`
}

type PageData struct {
	Title, MetaDescription string
	Year                   int
	Entries                []GuestbookEntry
	Photos                 []Photo
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

func loadJSON(path string, v any) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(v); err != nil {
		log.Fatal(err)
	}
}

func main() {
	t := template.Must(template.ParseGlob("templates/_*.html"))
	t = template.Must(t.ParseGlob("templates/*.html"))
	y := time.Now().Year()

	var entries []GuestbookEntry
	loadJSON("data/guestbook.json", &entries)

	var photos []Photo
	loadJSON("data/photos.json", &photos)

	write(t, "index.html", "index.html", PageData{Title: "Home", MetaDescription: "A memorial for Teddy.", Year: y})
	write(t, "gallery.html", "gallery/index.html", PageData{Title: "Gallery", MetaDescription: "Photos of Teddy.", Year: y, Photos: photos})
	write(t, "guestbook.html", "guestbook/index.html", PageData{Title: "Guestbook", MetaDescription: "Leave a message for Teddy.", Year: y, Entries: entries})
	write(t, "thanks.html", "thanks/index.html", PageData{Title: "Thank You", MetaDescription: "Thank you for your submission.", Year: y})
}
