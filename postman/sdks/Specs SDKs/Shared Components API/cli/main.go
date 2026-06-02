package main

import (
	_ "example.com/led-display-api/cmd/displaysummary"
	_ "example.com/led-display-api/cmd/displaytext"
	_ "example.com/led-display-api/cmd/predefinedicons"
	_ "example.com/led-display-api/config"
	"example.com/led-display-api/root"
	_ "example.com/led-display-api/setupauth"
)

func main() {
	root.Execute()
}
