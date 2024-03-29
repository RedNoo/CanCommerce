package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	//"time"

	dbUtil "CanCommerce/db"

	apis "CanCommerce/api"
	models "CanCommerce/models"

	services "CanCommerce/services"
	//jwt "CanCommerce/middleware"
)

var MainUrl string = "http://localhost:9099"

func init() {
	dbUtil.Connect()

	/*
		payload := jwt.Payload{Sub: "222", Name: "redno", Exp: time.Now().Add(1000000)}
		encodedResult := jwt.Encode(payload, "456")
		fmt.Println("encoded" + encodedResult)

		result, _ := jwt.Decode(encodedResult, "456")
		p := result.(jwt.Payload)
		fmt.Println("decoded" + p.Name)
	*/
}

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type App struct {
	ProductHandler *ProductHandler
	AdminHandler   *AdminHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header.Get("Authorization"))
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	if head == "product" {
		h.ProductHandler.ServeHTTP(res, req)
		return
	} else if head == "admin" {
		h.AdminHandler.ServeHTTP(res, req)
		return
	}

	return

	//http.Error(res, "Not Found", http.StatusNotFound)
}

type AdminHandler struct {
}

type PageData struct {
	PageTitle string
	Url       string
	Data      interface{}
}

func (h *AdminHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	fmt.Println(head)

	if head == "" {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))

		data := PageData{
			PageTitle: "My TODO list",
			Url:       MainUrl,
		}
		tmpl.Execute(res, data)
		return
	} else if head == "product" {

		head, req.URL.Path = ShiftPath(req.URL.Path)
		fmt.Println("product head : " + head)

		if head == "" || head == "home" {
			tmpl := template.Must(template.ParseFiles("./templates/common/header.html", "./templates/product/index.html", "./templates/common/footer.html"))

			count := services.GetProductCount()
			data := PageData{
				PageTitle: "My Products Home",
				Data:      count,
			}
			tmpl.ExecuteTemplate(res, "index", data)
		} else if head == "list" {
			tmpl := template.Must(template.ParseFiles("./templates/common/header.html", "./templates/product/list.html", "./templates/common/footer.html"))
			products := services.GetProducts(1)
			data := PageData{
				PageTitle: "My Products Home",
				Data:      products,
			}
			tmpl.ExecuteTemplate(res, "list", data)
		} else {
			tmpl := template.Must(template.ParseFiles("./templates/404.html"))
			tmpl.Execute(res, nil)
		}
	}

	return
}

type ProductHandler struct {
	CategoryHandler *CategoryHandler
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
}

func (h *ProductHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	enableCors(&res)

	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	fmt.Println("head: " + head)

	switch head {
	case "add":
		switch req.Method {
		case "POST":
			decoder := json.NewDecoder(req.Body)
			var p models.Product

			err := decoder.Decode(&p)
			if err != nil {
				panic(err)
			}

			apis.AddProduct(res, p)

			//title := p.Title
			//fmt.Println("title: " + title)
			//res.Write([]byte(title))

		default:
			http.Error(res, "Only POST is allowed", http.StatusMethodNotAllowed)
		}
	case "list":
		keys, ok := req.URL.Query()["page"]

		if !ok || len(keys[0]) < 1 {
			log.Println("Url Param 'key' is missing")
			return
		}

		i, err := strconv.Atoi(keys[0])
		if err != nil {
			panic(err)
		}

		apis.ListProducts(res, i)

		//res.Write([]byte(title))

	case "category":
		// We can't just make ProfileHandler an http.Handler; it needs the
		// user id. Let's instead…
		h.CategoryHandler.Handler(req.URL.Path).ServeHTTP(res, req)
	case "account":
		// Left as an exercise to the reader.
	default:
		http.Error(res, "Not Found", http.StatusNotFound)
	}
	return

}

type CategoryHandler struct {
}

func (h *CategoryHandler) Handler(id string) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var head string
		head, req.URL.Path = ShiftPath(req.URL.Path)
		fmt.Println("category head: " + head)
	})
}

func main() {

	a := &App{
		ProductHandler: new(ProductHandler),
	}

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", a)

	err := http.ListenAndServe(":9099", nil)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
