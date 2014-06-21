package main


import (
    "github.com/go-martini/martini"
    "net/http"
)

func main() {
  m := martini.Classic()
  
  m.Post("/getobj", GetObj)
  m.Run()
}


func GetObj(r *http.Request) string {
    bucket := r.FormValue("bucket")
    key := r.FormValue("key")
    return "Bucket:" + bucket + ", key:" + key
}
