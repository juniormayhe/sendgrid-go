package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

func sendMail(from string, to string, subject string, message string) (response *rest.Response, error error) {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	payload := ` {
		"personalizations": [
			{
				"to": [
					{
						"email": "%s"
					}
				],
				"subject": "%s"
			}
		],
		"from": {
			"email": "%s"
		},
		"content": [
			{
				"type": "text/plain",
				"value": "%s"
			}
		]
	}`

	payload = fmt.Sprintf(payload, to, subject, from, message)

	request.Body = []byte(payload)

	return sendgrid.API(request)
}

func main() {

	// from := mail.NewEmail("JuniorMayhe", "me@juniormayhe.com")
	// subject := "Sending with SendGrid is Fun"

	// to := mail.NewEmail("Junior", "me@juniormayhe.com")
	// plainTextContent := "and easy to do anywhere, even with Go"
	// htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	// message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	// client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	// response, err := client.Send(message)
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(response.StatusCode)
	// 	fmt.Println(response.Body)
	// 	fmt.Println(response.Headers)
	// }
	response, err := sendMail("me@juniormayhe.com", "me@juniormayhe.com", "This is a test", "This is my message")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("success")
		fmt.Println(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(response.Headers)
	}
}
