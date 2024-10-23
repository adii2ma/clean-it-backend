// models/ticket.go
package models

type Ticket struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Time      string `json:"time"`
    Room      string `json:"room"`
    Serial    string `json:"serial"`
    Status    bool   `json:"status"` // false = Pending, true = Done
}
