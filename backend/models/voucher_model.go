package models

type CheckRequest struct {
	FlightNumber string `json:"flightNumber"`
	Date         string `json:"date"`
}

type CheckResponse struct {
	Exists bool `json:"exists"`
}

type GenerateRequest struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	FlightNumber string `json:"flightNumber"`
	Date         string `json:"date"`
	Aircraft     string `json:"aircraft"`
}

type GenerateResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats,omitempty"`
	Error   string   `json:"error,omitempty"`
}

type Voucher struct {
	ID           int
	CrewName     string
	CrewID       string
	FlightNumber string
	FlightDate   string
	AircraftType string
	Seat1        string
	Seat2        string
	Seat3        string
	CreatedAt    string
}