package main

import (
  "net/http"
  "strconv"
)

var hum float64
var temp float64
var speed float64


func get(w http.ResponseWriter, r *http.Request) {
  Fetchweb()	
  Computation()
  message := strconv.FormatFloat(speed, 'f', 1, 64)

   w.Write([]byte(message))
}

func main() {
  http.HandleFunc("/", get)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}