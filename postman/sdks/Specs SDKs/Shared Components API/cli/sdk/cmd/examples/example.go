package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"context"
	"example.com/led-display-api/sdk"
	"example.com/led-display-api/sdk/displaytext"
)

func main() {
	loadEnv()

	config := leddisplayapi.NewConfig()
	client := leddisplayapi.NewLedDisplayAPI(config)

	request := displaytext.CreateDisplayTextRequest{
		Text: leddisplayapi.Ptr("text"),
	}

	response, err := client.DisplayText.CreateDisplayText(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
