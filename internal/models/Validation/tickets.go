package validation

type TicketResponse struct {
	ID          int
	Title       string
	Description string
	Status      bool
	CreatedAt   string
	UpdatedAt   string
}
type TicketRequest struct {
	Title       string
	Description string
}
