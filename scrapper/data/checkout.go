package data

import (
	"context"
	"fmt"
	"regexp"

	"github.com/lai0xn/satim-dolphin/store"
)

func LoadCheckout(data map[string]interface{}, store store.Store) {
	pattern := `([0-9]+(?:\.[0-9]{2})?)\s?(?:DZD|DZ\s?DA|\$)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(data["content"].(string), 1)

	if len(matches) > 0 {
		amount := matches[0][1]
		store.Set(context.Background(), data["host"].(string)+":amount", amount)
	} else {
		fmt.Println("No match found")
	}
}
