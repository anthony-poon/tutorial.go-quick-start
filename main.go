// Declare 'package main' to make this program an executable.
package main

import (
	"anthony-poon.com/quickstart/translate"
	"fmt"
	"log"
	"os"
	"time"
)
// Declare a function
func main()  {
	// Initialize and assign a value to a variable
	name := "John Doe"
	t := time.Now()
	var greeting string
	if t.Hour() < 12 {
		greeting = "Good morning"
	} else {
		greeting = "Good Afternoon"
	}
	// Echo / Print result
	greeting = greeting + fmt.Sprintf(" %s, it is %s currently", name, t.Format(time.Kitchen))
	fmt.Println(greeting)
	langs := []string {
		"zh-TW",
		"ja",
		"ko",
		"es",
	}
	for _, lang := range langs {
		translated, err := translate.TranslateText(lang, greeting)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(translated)
	}
}