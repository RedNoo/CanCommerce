package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	dbUtil "CanCommerce/db"

	apis "CanCommerce/api"
	models "CanCommerce/models"

	jwt "CanCommerce/middleware"
)

func init() {
	dbUtil.Connect()

	payload := jwt.Payload{Sub: "222", Name: "redno", Exp: time.Now().Add(1000000)}
	encodedResult := jwt.Encode(payload, "456")
	fmt.Println("encoded" + encodedResult)

	result, _ := jwt.Decode(encodedResult, "456")
	p := result.(jwt.Payload)
	fmt.Println("decoded" + p.Name)
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
}

type ProductHandler struct {
	CategoryHandler *CategoryHandler
}

func (h *ProductHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	if head == "product" {
		h.ProductHandler.ServeHTTP(res, req)
		return
	}
	http.Error(res, "Not Found", http.StatusNotFound)
}

func main() {

	a := &App{
		ProductHandler: new(ProductHandler),
	}
	err := http.ListenAndServe(":9099", a)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
