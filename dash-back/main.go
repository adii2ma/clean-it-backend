// main.go
package main

import (
    "net/http"
    "myapp/db"
    "myapp/routes"
)

func main() {
    db.ConnectDB()
    routes.SetupRoutes()
    http.ListenAndServe(":8080", nil) // Serve on port 8080
}
