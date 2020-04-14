package main

import (
	"net/http"

	"github.com/rakeshkonni/go-play-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
