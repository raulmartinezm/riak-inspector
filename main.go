package main

import (
  "github.com/codegangsta/martini"
  "net/http"
)

func main() {
  m := martini.Classic()

  m.Post("/connectdb", connectdb)

  m.Run() 
}


func connectdb (r *http.Request) string {
  return r.FormValue("host")+":"+r.FormValue("port")
}
