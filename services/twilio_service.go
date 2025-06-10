package services

import (
	"encoding/json"
	"errors"
	"fmt"
	twilio_dto "gorm/models/entity/twilio"
	twilio_request "gorm/models/request/twilio"
	"os"

	twilio "github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

var twilioSID = os.Getenv("TWILIO_ACCOUNT_SID")
var twilioAuthToken = os.Getenv("TWILIO_AUTH_TOKEN")

func SendMessage(body twilio_request.TwilioApiBody)(*api.ApiV2010Message, error){

	templates := map[string]twilio_dto.Template{
    "invitation": {
      ID: "HXb6230bfac43e5a80430e282c621f68ca",
      Language: "en-US",
      Variables: map[string]string{
        "1": body.Name, // guest name
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
	templateRequest := templates[body.TemplateName]


	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:+19162324303")
	params.SetTo(fmt.Sprintf("whatsapp:%s",body.Phone))
	params.SetContentSid(templateRequest.ID)
	variablesJSON, err := json.Marshal(templateRequest.Variables)
	if err != nil {
		return nil, errors.New("failed to encode variables")
	}
	params.SetContentVariables(string(variablesJSON))

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return nil, err
	}
	return resp, nil
}