package tester

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/log"

	"github.com/lai0xn/satim-proxy/redis"
)

func TestFraud(key string, data map[string]interface{}) {
	amount, err := redis.Client.Get(context.Background(), key).Result()
	if err != nil {
		log.Warn(err)
		return
	}
	am, _ := strconv.ParseFloat(amount, 64)
	s_am := data["Amount"].(float64)
	fmt.Println(am)
	fmt.Println(s_am)
	if am*100 == s_am {
		log.Info("Fraud Test Passed")
	} else {
		log.Warn("Fraud Test Failed")
	}
	b, _ := json.Marshal(data)
	redis.Client.Set(context.Background(), "naviguih.com:checkout", string(b), time.Hour*1)
}
