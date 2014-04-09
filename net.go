package main

import (
  "fmt"
  "net/http"
  "os"
)

var(

)

func serverInit() {

  http.Handle("/", http.FileServer(http.Dir(".")))

  //Starting server
  fmt.Println("Server starting...")
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Printf("http failed: %s\n", err.Error())
    os.Exit(1)
  }
}


/*
func wsHandler(ws *websocket.Conn) {
  defer ws.Close()

  ws.Request()
  fmt.Println("Client is Connected")

}

*/


/*
func main() {
  serverInit()
}
*/

