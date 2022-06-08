package command

import (
	"bingoBotGo/internal/trigger"
	types "bingoBotGo/internal/types"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	GeneratorHost = "localhost"
	GeneratorPort = 5000
)

type nameGenResponse struct {
	Error        bool   `json:"error"`
	Data         string `json:"data,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func MakeGenerateName(botName string) TriggeredCommand {
	return TriggeredCommand{
		Name:           "GenerateNameCommand",
		SelfTriggering: false,
		Trigger:        trigger.MakeNamePrefixedBasicStringMatch(botName, "give me a name"),
		Action:         generateNameAction,
	}
}

func generateNameAction(bot types.IBot, message *discordgo.Message) (Result, error) {
	log.Println("Generate Name command triggered")

	requestUrl := fmt.Sprintf("http://%s:%d/api/v1/ai/generator/strings/person_names", GeneratorHost, GeneratorPort)
	res, err := http.Get(requestUrl)
	if err != nil {
		log.Printf("Failed to generate name: %s\n", err)
		return FAILURE, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close HTTP response body")
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Failed to read HTTP response body")
		return FAILURE, err
	}

	var response nameGenResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Failed to unmarshal HTTP response into understandable format")
		return FAILURE, err
	}

	name := response.Data

	_, err = bot.SendMessageWithTyping(message.ChannelID, fmt.Sprintf("What about %s", name))
	if err != nil {
		log.Println("Failed to send message reply")
		return FAILURE, err
	}

	return SUCCESS, nil
}
