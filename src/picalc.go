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

// Purchases a 'product' equal to database id.
func purchaseProduct(product int) string{
    id := product
    sqlStatement := 'SELECT col FROM products WHERE id=$1'
    row := db.QueryRow(sqlStatement, id)
    err := row.Scan(&col)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("Zero rows found")
        } else {
            panic(err)
        }
    }
    return row
}

func handlerPicalc(w http.ResponseWriter, r *http.Request) {
    log.Print("Pi calculator received a request.")
    iterations, err := strconv.Atoi(r.URL.Query()["iterations"][0])
    if err != nil {
        fmt.Fprintf(w, "iterations parameter not valid\n")
        return
    }
    fmt.Fprintf(w, "%.10f\n", calculatePi(iterations))
}

func handlerPurchase(w http.ResponseWriter, r *http.Request) {
    log.Print("Purchase function received a request.")
    product, err := strconv.Atoi(r.URL.Query()["product"][0])
    if err != nil {
        fmt.Fprintf(w, "product parameter not valid\n")
        return
    }
    fmt.Fprint(w, "%s", purchaseProduct(product))
}

func main() {
    log.Print("Pi calculator started.")

    http.HandleFunc("/picalc", handlerPicalc)
    http.HandleFunc("/purchase", handlerPurchase)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
