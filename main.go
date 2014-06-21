package main

import (
	"github.com/codegangsta/martini"
	"net/http"
	"github.com/mrb/riakpbc"
)

func main() {
	m := martini.Classic()

	m.Post("/connectdb", ConnectDB)
	m.Post("/getobj", GetObj)
	m.Run()
}


func ConnectDB(r *http.Request) string {
	return r.FormValue("host") + ":" + r.FormValue("port")

}


func GetObj(r *http.Request) string {
	bucket := r.FormValue("bucket")
	key := r.FormValue("key")
	return "Bucket:" + bucket + ", key:" + key
}
