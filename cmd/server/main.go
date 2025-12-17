package main

import (
	//"encoding/json"
	//"fmt"
	//"github.com/joho/godotenv"
	//"github.com/twilio/twilio-go"
	//twilioapi "github.com/twilio/twilio-go/rest/api/v2010"
	//"log"
	//"os"

	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"meowmos/internal/tui"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func main() {
	p := tea.NewProgram(tui.InitAppModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	//accSID := os.Getenv("TWILIO_ACC_SID")
	//authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	//senderNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	//
	//client := twilio.NewRestClientWithParams(twilio.ClientParams{
	//	Username: accSID,
	//	Password: authToken,
	//})
	//
	//params := &twilioapi.CreateMessageParams{}
	//params.SetTo("+17802434269")
	//params.SetFrom(senderNumber)
	//params.SetBody("Hello, World!")
	//
	//resp, err := client.Api.CreateMessage(params)
	//if err != nil {
	//	fmt.Println("Error sending SMS message: " + err.Error())
	//} else {
	//	response, _ := json.Marshal(*resp)
	//	fmt.Println("Response: " + string(response))
	//}
}
