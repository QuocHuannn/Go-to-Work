package main

import (
	"github.com/QuocHuannn/Go-to-Work/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8002")
}
