package twilio

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/gin-gonic/gin"
	twilio "github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

// type TwilioConfig struct {
// 	AccountSID string `json:"account_sid" binding:"required"`
// 	AuthToken  string `json:"auth_token" binding:"required"`
// 	FromNumber string `json:"from_number" binding:"required"`
// }
var twilioSID = os.Getenv("TWILIO_ACCOUNT_SID")
var twilioAuthToken = os.Getenv("TWILIO_AUTH_TOKEN")

type TwilioApiBody struct {
	Phone string `json:"phone" binding:"required"`
	TemplateName string `json:"templateName" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Template struct {
	ID string `json:"id"`
	Language string `json:"language"`
	Variables map[string]string `json:"variables"`
}

func SendMessage(c *gin.Context){

	fmt.Println(twilioSID, twilioAuthToken)
	var request TwilioApiBody
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	templates := map[string]Template{
    "invitation": {
      ID: "HXb6230bfac43e5a80430e282c621f68ca",
      Language: "en-US",
      Variables: map[string]string{
        "1": request.Name, // guest name
        "2": "Pavithra O", // bride's name
        "3": "Lokeshwar Kumar T", // groom's name
        "4": "May 22nd and 23rd, 2025", // date
        "5": "Chennai", // venue
      },
    },
}
	

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioSID,
		Password: twilioAuthToken,
	})

	templateRequest := templates[request.TemplateName]

	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:+19162324303")
	params.SetTo(fmt.Sprintf("whatsapp:%s",request.Phone))
	params.SetContentSid(templateRequest.ID)
	variablesJSON, err := json.Marshal(templateRequest.Variables)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to encode template variables"})
		return
	}
	params.SetContentVariables(string(variablesJSON))


	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if resp.Body != nil {
			c.JSON(200, gin.H{"message": "Message sent successfully", "response": resp})
		} else {
			c.JSON(500, gin.H{"error": "Failed to send message"})
		}
	}

}