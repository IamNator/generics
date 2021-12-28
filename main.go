package main

import (
	"generics/controller"
	"generics/model"
	"generics/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"sync"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	str := storage.New()
	str.InitModels()
	cntrl := controller.New(str)
	r := gin.Default()
	r.POST("/", func(ctx *gin.Context) {
		var u model.User
		if er := ctx.ShouldBindJSON(&u); er != nil {
			ctx.String(http.StatusBadRequest, er.Error())
			return
		}

		u_, er := cntrl.RegisterUser(u)
		if er != nil {
			ctx.String(http.StatusUnprocessableEntity, er.Error())
			return
		}

		ctx.JSONP(http.StatusOK, u_)
	})

	port := ":8000"
	log.Printf("server running at port %s\n", port)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if er := http.ListenAndServe(port, r); er != nil {
			log.Println(er.Error())
		}
	}()

	wg.Wait()
}
