package twilio_request

type TwilioApiBody struct {
	Phone string `json:"phone" binding:"required"`
	TemplateName string `json:"templateName" binding:"required"`
	Name string `json:"name" binding:"required"`
}