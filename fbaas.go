package main

import (
        "strconv"
        "regexp"
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

var rNum = regexp.MustCompile(`^/\d+$`)
var rSlash = regexp.MustCompile(`^/$`)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
  var o []string
  if rNum.MatchString(r.URL.Path) {
    p := r.URL.Path
    s := p[1:]
    n, er := strconv.Atoi(s)
    if er != nil {
      http.Error(w, er.Error(), http.StatusInternalServerError)
    }
    o = fizzbuzz(n)
  } else if rSlash.MatchString(r.URL.Path) {
    o = fizzbuzz(20)
  } else {
    o = append(o, "Help: valid paths are / or /(number). /(number) outputs (number) lines of fizzbuzz. / outputs 20 lines. any other request gets this message")
  }

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
