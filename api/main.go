package handler

import (
	"context"
	_ "encoding/json"
	"fmt"
	. "handler/views"
	"net/http"
	"os"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()
	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))

	// Define route handlers
	server.GET("/main-page", func(c *Context) {
		SendMainPage(c.Writer, c.Req)
	})
	server.GET("/profile", func(c *Context) {
		SendProfilePage(c.Writer, c.Req)
	})
	server.GET("/tag/", func(c *Context) {
		SendTagPage(c.Writer, c.Req)
	})
	server.GET("/info", func(c *Context) {
		SendInfoPage(c.Writer, c.Req)
	})
	server.GET("/img", func(c *Context) {
		SendCatImg(c.Writer, c.Req)
	})

	server.Handle(w, r)
}

func SendMainPage(w http.ResponseWriter, r *http.Request) {
	ShowHome().Render(context.TODO(), w)
}
func SendProfilePage(w http.ResponseWriter, r *http.Request) {
	ShowProfile("tomek").Render(context.TODO(), w)

}
func SendTagPage(w http.ResponseWriter, r *http.Request) {
	ids := []int{1, 2, 3}
	comments := []string{"Super Dzień", "'Obcy na dzielni' takie 2/10", "Pizzeria Fabio 4.5/10"}
	authors := []string{"tomek", "lolek", "pies4"}
	ShowTag(ids, authors, comments).Render(context.TODO(), w)
}
func SendInfoPage(w http.ResponseWriter, r *http.Request) {
	ShowInfo().Render(context.TODO(), w)
}
func SendDebug(w http.ResponseWriter, r *http.Request) {
	ShowDebug().Render(context.TODO(), w)
}
func SendLoginError(w http.ResponseWriter, r *http.Request) {
	LoginError().Render(context.TODO(), w)
}
func SendCatImg(w http.ResponseWriter, r *http.Request) {
	img, err := os.ReadFile("img.png")
	if err != nil {
		http.Error(w, "Błąd odczytu pliku obrazka", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(img)
}
func SendFullPost(w http.ResponseWriter, r *http.Request, content string, author string, comcontent []string, comauthor []string) {
	ShowFullPost("Super Dzień!!!", "Tomek", []string{"U mnie też super!", "U mnie nie"}, []string{"Pies4", "Adam1919"}).Render(context.TODO(), w)
}
