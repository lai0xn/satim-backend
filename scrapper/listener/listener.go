package listener

import (
	"context"
	"encoding/json"

	"github.com/charmbracelet/log"

	"github.com/lai0xn/satim-dolphin/data"
	"github.com/lai0xn/satim-dolphin/store"
	"github.com/lai0xn/satim-dolphin/tester"
)

func Listen(store store.Store) {
	subsriber := store.Subscribe(context.Background(), "test:checkout")
	defer subsriber.Close()
	ch := subsriber.Channel()
	log.Info("Started Listening")
	for msg := range ch {

		var d map[string]interface{}
		err := json.Unmarshal([]byte(msg.Payload), &d)
		if err != nil {
			log.Fatalf("Error parsing JSON payload: %v", err)
		}

		config := tester.Config{
			GreenIconLink:     "satim.cdn.dz/green_number.svg",
			TermsCheckboxName: "terms",
			PaymentMethod:     "CIB/EDAHABIA",
		}
		content := d["content"].(string)
		tester := tester.Tester{
			Content: content,
			Config:  config,
		}
		if err := tester.TestTermsAndConditions(); err != nil {
			log.Info(err)
		}
		log.Info("Terms and conditions passed")
		if err := tester.TestPaymentMethod(); err != nil {
			log.Info(err)
		}
		log.Info("Payment method found")
		log.Info(d["host"])
		data.LoadCheckout(d, store)
	}
}
