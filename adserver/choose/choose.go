package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)


type Page struct {
	Title string
	Body []byte
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("viewHandler")

	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	
	if err != nil {
		p = &Page{Title: title}
	}

	/*fmt.Fprintf(w, "<h1>Editing %s</h1>" + 
		"<form action=\"/save/%s\" method=\"POST\">" +
		"<textarea name=\"body\">%s</textarea><br>" +
		"<input type=\"submit\" value=\"Save\">" +
		"</form>",
		p.Title, p.Title, p.Body)*/

	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func main() {
	//p1 := &Page {Title: "tp", Body: []byte("ttpp")}
	//p1.save()

	//p2, _ := loadPage("tp")
	//fmt.Println(string(p2.Body))

	//http.HandleFunc("/", handler)
	http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save", saveHandler)
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)

	fmt.Println("end")
}