package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
)

// Calculate pi using Gregory-Leibniz series:   (4/1) - (4/3) + (4/5) - (4/7) + (4/9) - (4/11) + (4/13) - (4/15) ...
func calculatePi(iterations int) float64 {
    var result float64 = 0.0
    var sign float64 = 1.0
    var denominator float64 = 1.0
    for i := 0; i < iterations; i++ {
        result = result + (sign * 4/denominator)
        denominator = denominator + 2
        sign = -sign
    }
    return result
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Print("Pi calculator received a request.")
    iterations, err := strconv.Atoi(r.URL.Query()["iterations"][0])
    if err != nil {
        fmt.Fprintf(w, "iterations parameter not valid\n")
        return
    }
    fmt.Fprintf(w, "%.10f\n", calculatePi(iterations))
}

func main() {
    log.Print("Pi calculator started.")

    http.HandleFunc("/", handler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
