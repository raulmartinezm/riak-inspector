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


func ConnectDB(r *http.Request, c martini.Context) string {

	host := r.FormValue("host")
	port := r.FormValue("port")
	client := riakpbc.NewClient([]string{
			host + ":" + port,
})

	c.Map(client)
	return "connected"

}


func GetObj(r *http.Request, riak *riakpbc.Client) string {
	bucket := r.FormValue("bucket")
	key := r.FormValue("key")

	one, err := riak.FetchObject(bucket, key)
	if err != nil {
		panic(err)
	}
	return string(one.GetContent()[0].GetValue())
}
