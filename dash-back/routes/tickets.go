// routes/tickets.go
package routes

import (
    "net/http"
    "myapp/controllers"
)

func SetupRoutes() {
    http.HandleFunc("/api/tickets", controllers.GetTickets)
    http.HandleFunc("/api/tickets/update", controllers.UpdateTicketStatus)
}
