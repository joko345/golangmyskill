package main

import (
	"net/http"

	"github.com/joko345/golangmyskill/mongo-go/control"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := control.NewUserControl(getSession()) //newUser akan menerima value dari getSession
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r) //r untuk handle crud
}

func getSession() *mgo.Session { //*pointer merupakan return parameter
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
