/**
 *
 * This is just some example code for Go.
 * It shows a self serving application that uses Thrust to display pages.
 * If you use this to build any projects @0x4445565a on twitter.
 * I love to see cool stuff being made!
 *
 */
package main

import (
  "fmt"
  "net"
  "strings"
  "strconv"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
  "github.com/miketheprogrammer/go-thrust/lib/bindings/window"
  "github.com/miketheprogrammer/go-thrust/lib/bindings/session"
  "github.com/miketheprogrammer/go-thrust/thrust"
  "github.com/miketheprogrammer/go-thrust/lib/commands"
)

func handlerDebug(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  for k, v := range r.Form {
    fmt.Println("key:", k)
    fmt.Println("val:", strings.Join(v, ""))
  }
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  t := loadTemplates("templates/index.gtpl")
  t.Execute(w, map[string]interface{}{
    "name": "Partner",
  })
  handlerDebug(w, r)
  if r.Method == "POST" {
    _ = addWindow(Current.uri + "new-window", commands.SizeHW{
      Width:  500,
      Height: 500,
    })
  }
}

func newWindowHandler(w http.ResponseWriter, r *http.Request) {
  t := loadTemplates("templates/new-window.gtpl")
  t.Execute(w, map[string]interface{}{})
  handlerDebug(w, r)
}

func angularExampleHandler(w http.ResponseWriter, r *http.Request) {
  t := loadTemplates("templates/angular.gtpl")
  t.Execute(w, map[string]interface{}{})
  handlerDebug(w, r)
}

func loadTemplates(name ...string) *template.Template {
  name = append(name, "templates/head.gtpl")
  name = append(name, "templates/head-angular.gtpl")
  name = append(name, "templates/nav.gtpl")
  name = append(name, "templates/footer.gtpl")
  t := template.Must(template.ParseFiles(
    name...
  ))

  return t
}

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
  if string(r.URL.Path[len(r.URL.Path) - 1]) == "/" {
    http.NotFound(w, r)
  } else {
    http.ServeFile(w, r, r.URL.Path[1:])
  }
}

/**
 * Thanks https://github.com/phayes/freeport
 * It's super simple
 */
func GetPort() string {
  addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
  if err != nil {
    panic(err)
  }

  l, err := net.ListenTCP("tcp", addr)
  if err != nil {
    panic(err)
  }
  defer l.Close()
  return strconv.Itoa(l.Addr().(*net.TCPAddr).Port)
}

var Current struct {
  session *session.Session
  port string
  uri string
  windows map[int]*window.Window
}

func addWindow(url string, size commands.SizeHW) int {
  index := len(Current.windows) - 1
  if index < 0 {
    index = 0
  } 
  Current.windows[index] = thrust.NewWindow(thrust.WindowOptions{
    RootUrl: url,
    Session: Current.session,
    Size: size,
  })
  Current.windows[index].Show()
  Current.windows[index].Focus()
  return index
}

func initThrust() {
  thrust.InitLogger()
  thrust.Start()
  
  Current.session = thrust.NewSession(false, false, ".cache")
  _ = addWindow(Current.uri, commands.SizeHW{
    Width:  750,
    Height: 750,
  })
}

func main() {
  fmt.Println("Checking for an open port...")
  Current.port = GetPort()
  Current.uri = "http://localhost:" + Current.port + "/"
  Current.windows = make(map[int]*window.Window)
  fmt.Println("Using port", Current.port)
  fmt.Println("Building Router...")
  r := mux.NewRouter()
  r.HandleFunc("/", rootHandler)
  r.HandleFunc("/angular", angularExampleHandler)
  r.HandleFunc("/new-window", newWindowHandler)
  http.HandleFunc("/css/", staticFileHandler)
  http.HandleFunc("/js/", staticFileHandler)
  http.Handle("/", r)

  fmt.Println("Init Thrust...")
  initThrust()

  fmt.Println("Ready and serving!")
  http.ListenAndServe(":" + Current.port, nil)
}
