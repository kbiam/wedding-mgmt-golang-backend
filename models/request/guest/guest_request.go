package guest_request

type AddGuestRequest struct {
	Name       string `json:"name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Relation   string `json:"relation"`
	Side       string `json:"side"`
	GuestCount int    `json:"guest_count"`
}

type UpdateInvitationStatusRequest struct {
	IsInvited bool `json:"is_invited" binding:"required"`
}