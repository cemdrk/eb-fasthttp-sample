package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func handler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "RequestURI is %q", ctx.RequestURI())
}


func main() {
	fasthttp.ListenAndServe(":5000", handler)
}
