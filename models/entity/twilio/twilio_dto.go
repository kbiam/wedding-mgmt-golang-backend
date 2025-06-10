package twilio_dto

type Template struct {
	ID string `json:"id"`
	Language string `json:"language"`
	Variables map[string]string `json:"variables"`
}