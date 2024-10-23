// controllers/ticketsController.go
package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "myapp/models"
    "myapp/db"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
    rows, err := db.DB.Query("SELECT id, name, time, room, serial, status FROM tickets")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    tickets := []models.Ticket{}
    for rows.Next() {
        var ticket models.Ticket
        if err := rows.Scan(&ticket.ID, &ticket.Name, &ticket.Time, &ticket.Room, &ticket.Serial, &ticket.Status); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tickets = append(tickets, ticket)
    }

    json.NewEncoder(w).Encode(tickets)
}

func UpdateTicketStatus(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    newStatus := r.URL.Query().Get("status")

    _, err := db.DB.Exec("UPDATE tickets SET status = $1 WHERE id = $2", newStatus, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
