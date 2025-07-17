package main

import (
    "lib/internal/routers"
)

func main() {
    r := routers.SetupRouter()
    r.Run() // default to :8080
}

