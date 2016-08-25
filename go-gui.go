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
  "github.com/miketheprogrammer/go-thrust/thrust"
)

func handlerDebug(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
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
  //r.ParseForm()
  handlerDebug(w, r)
}

func loadTemplates(name ...string) *template.Template {
  name = append(name, "templates/head.gtpl")
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

func initThrust(port string) {
  //thrust.InitLogger()
  thrust.Start()
  
  session := thrust.NewSession(false, false, "cache")
  thrustWindow := thrust.NewWindow(thrust.WindowOptions{
    RootUrl: "http://localhost:" + port + "/",
    Session: session,
  })
  thrustWindow.Show()
  //thrustWindow.Maximize()
  thrustWindow.Focus()
}

func main() {
  fmt.Println("Checking for an open port...")
  port := GetPort()

  fmt.Println("Using port", port)
  fmt.Println("Building Router...")
  r := mux.NewRouter()
  r.HandleFunc("/", rootHandler)
  http.HandleFunc("/css/", staticFileHandler)
  http.Handle("/", r)

  fmt.Println("Init Thrust...")
  initThrust(port)

  fmt.Println("Ready and serving!")
  http.ListenAndServe(":" + port, nil)
}
