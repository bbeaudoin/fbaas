package main

import (
        "strconv"
        "net/http"
        "log"
        "encoding/json"
)

func fizzbuzz(n int) []string {
   var output []string

   for i := 1; i<=n; i++ {
     if i%3 == 0 {
        if i%5 == 0 {
          output = append(output, "fizzbuzz")
        } else {
          output = append(output, "fizz")
        }
     } else if i%5 == 0 {
        output = append(output, "buzz")
     } else {
        output = append(output, strconv.Itoa(i))
     }
   }
   return output
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  o := fizzbuzz(20)

  js, err := json.Marshal(o)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func main() {
  http.HandleFunc("/", defaultHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
