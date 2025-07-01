package dto

type OrderSummaryResponse struct {
	TotalPedidosMes  int            `json:"totalPedidosMes"`
	VentasTotalesMes float64        `json:"ventasTotalesMes"`
	TicketPromedio   float64        `json:"ticketPromedio"`
	PorEstado        map[string]int `json:"porEstado"`
}
