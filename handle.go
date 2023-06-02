package function

import (
	"context"
	"fmt"
	"net/http"
)

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, fmt.Sprintf("%.12f", calculatePi(10000)))
}

func calculatePi(n int) float64 {
    pi := 0.0
    sign := 1.0

    for i := 0; i < n; i++ {
        pi += sign / (2.0 * float64(i) + 1.0)
        sign *= -1.0
    }

    return 4.0 * pi
}



