package main

import (
	"fmt"
	"math"
	"net/http"
)

func dosqrt(a float64) float64 {
    return math.Sqrt(a)
}

func main() {

    sum := 0.0
    for i := 0.0; i < 1000000000.0; i++ {
    	sum += dosqrt(i)
    }

    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

    out := `<!DOCTYPE html>
        <html lang="en">
        <head>
        <meta charset="UTF-8">
        <title>Go HPA Server</title>
        </head>
        <body><b>Code.education Rocks!</b></body>
        </html>`;

        fmt.Fprintf(w, out);
    })

    http.ListenAndServe(":8000", nil)
}
