package main

import (
	//"encoding/json"
	//"fmt"
	//"github.com/joho/godotenv"
	//"github.com/twilio/twilio-go"
	//twilioapi "github.com/twilio/twilio-go/rest/api/v2010"
	//"log"
	//"os"

	"fmt"
	"log"
	"os"

	"meowmos/internal/database"
	"meowmos/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func main() {

	url := os.Getenv("TURSO_DATABASE_URL")
	token := os.Getenv("TURSO_AUTH_TOKEN")

	db, err := database.ConnectTo(url+token, "libsql")
	if err != nil {
		log.Fatal("Could not connect to database." + err.Error())
	}


	fmt.Println("Connected to Turso database.")

	p := tea.NewProgram(tui.InitAppModel(db))
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
