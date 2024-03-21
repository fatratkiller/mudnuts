package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	discordToken      = "Discord Bot Token"
	discordChannelID  = "CHANNEL ID"
	geckoEndpoint     = "your api link"
	geckoRequestHeader = map[string]string{
		"Accept": "application/json;version=20230302",
	}
)

type Token struct {
	Attributes map[string]interface{} `json:"attributes"`
}

type TokenResponse struct {
	Data []Token `json:"data"`
}

func getTokens() []Token {
	client := &http.Client{}
	req, err := http.NewRequest("GET", geckoEndpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return []Token{}
	}

	for key, value := range geckoRequestHeader {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching tokens:", err)
		return []Token{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []Token{}
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return []Token{}
	}

	return tokenResp.Data
}

func postToDiscord(s *discordgo.Session, tokens []Token) {
	for _, token := range tokens[:20] { // Only post top 20 tokens
		fields := make([]*discordgo.MessageEmbedField, 0)
		for key, value := range token.Attributes {
			field := &discordgo.MessageEmbedField{
				Name:   key,
				Value:  fmt.Sprintf("%v", value),
				Inline: false,
			}
			fields = append(fields, field)
		}

		embed := &discordgo.MessageEmbed{
			Title:       "💰Today's Big Hit #crypto 🪙 in Mudnuts' Watchlist⚠️",
			Color:       0x00ff00,
			Fields:      fields,
			Description: "----------------------------------------------", // Separator
		}

		s.ChannelMessageSendEmbed(discordChannelID, embed)
	}
}

func main() {
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Logged in as:", s.State.User)
		go func() {
			for {
				tokens := getTokens()
				sort.Slice(tokens, func(i, j int) bool {
					return tokens[i].Attributes["gt_score"].(float64) > tokens[j].Attributes["gt_score"].(float64)
				})
				postToDiscord(s, tokens)
				time.Sleep(24 * time.Hour) // Sleep for 24 hours
			}
		}()
	})

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
	}
	defer dg.Close()

	// Keep the program running until interrupted.
	stop := make(chan os.Signal, 1)
	<-stop
}
