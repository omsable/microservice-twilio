package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	accountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	authToken  = os.Getenv("TWILIO_AUTH_TOKEN")
	from       = os.Getenv("TWILIO_FROM")
	to         = os.Getenv("TWILIO_TO")
)

var _ = Describe("Send SMS", func() {

	os.Setenv("ACCOUNT_SID", accountSID)
	os.Setenv("AUTH_TOKEN", authToken)
	sms := SMS{From: from, To: to, Message: "Testing twilio microservice"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(sms)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send sms message", func() {
		Context("send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Send SMS negative test", func() {

	sms := []byte(`{"status":false}`)

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(sms)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send sms message", func() {
		Context("send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
