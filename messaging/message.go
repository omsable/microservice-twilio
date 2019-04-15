package messaging

import (
	"encoding/json"
	result "github.com/heaptracetechnology/microservice-twilio/result"
	"github.com/sfreiberg/gotwilio"
	"net/http"
	"os"
)

type SMS struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}

//Send Email
func Send(responseWriter http.ResponseWriter, request *http.Request) {

	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	decoder := json.NewDecoder(request.Body)
	var param SMS
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	res, exp, respErr := twilio.SendSMS(param.From, param.To, param.Message, "", "")
	if respErr != nil {
		result.WriteErrorResponse(responseWriter, respErr)
		return
	}

	if res != nil {
		bytes, _ := json.Marshal(res)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
	} else if exp != nil {
		bytes, _ := json.Marshal(res)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
	}
}
