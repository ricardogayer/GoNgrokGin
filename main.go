// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"golang.ngrok.com/ngrok"
// 	"golang.ngrok.com/ngrok/config"
// )

// func main() {
// 	if err := run(context.Background()); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func run(ctx context.Context) error {
// 	tun, err := ngrok.Listen(ctx,
// 		config.HTTPEndpoint(),
// 		ngrok.WithAuthtokenFromEnv(),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("tunnel created:", tun.URL())

// 	return http.Serve(tun, http.HandlerFunc(handler))
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "<h1>Hello from ngrok-go.</h1>")
// }

// go get golang.ngrok.com/ngrok

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	log.Println("tunnel created:", tun.URL())

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/", handler)

	return http.Serve(tun, router)
}

func handler(c *gin.Context) {

	p1 := Produto{
		Id:        100,
		Descricao: "Bomba hidráulica",
		Valor:     4500.56,
	}

	p2 := Produto{
		Id:        200,
		Descricao: "Motor Weg 2HP",
		Valor:     7550.00,
	}

	p3 := Produto{
		Id:        300,
		Descricao: "Redutor 1:3",
		Valor:     14500.56,
	}

	var produtos []Produto

	produtos = append(produtos, p1, p2, p3)

	c.JSON(http.StatusOK, produtos)
}

type Produto struct {
	Id        int64
	Descricao string
	Valor     float64
}
