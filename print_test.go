package log_test

import (
	"strings"
	"testing"
	"time"

	"github.com/myste1tainn/hexlog"
	"github.com/spf13/viper"
)

func TestPrint(t *testing.T) {
	viper.AddConfigPath("./")
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}

	log.InitGlobalLogger()
	// log.Warnf("test")

	// log.Errorf("error %v", map[string]any{
	// 	"request": "path/to/somewhere",
	// 	"method":  "POST",
	// 	"body": map[string]any{
	// 		"firstName": "John",
	// 		"lastName":  "Doe",
	// 		"action":    "Greeting",
	// 		"arrs": []map[string]any{
	// 			{
	// 				"firstName": "John1",
	// 				"lastName":  "Doe1",
	// 				"action":    "Greeting1",
	// 			},
	// 			{
	// 				"firstName": "John2",
	// 				"lastName":  "Doe2",
	// 				"action":    "Greeting2",
	// 			},
	// 			{
	// 				"firstName": "John3",
	// 				"lastName":  "Doe3",
	// 				"action":    "Greeting3",
	// 			},
	// 		},
	// 	},
	// 	"t": time.Now(),
	// })

	// log.Debugf("debug %v", Req{
	// 	Request: "path/to/somewhere",
	// 	Method:  "POST",
	// 	Body: T{
	// 		FirstName: "John",
	// 		LastName:  "Doe",
	// 		Action:    "Greeting",
	// 		Arrs: []T{
	// 			{
	// 				FirstName: "John1",
	// 				LastName:  "Doe1",
	// 				Action:    "Greeting1",
	// 			},
	// 			{
	// 				FirstName: "John2",
	// 				LastName:  "Doe2",
	// 				Action:    "Greeting2",
	// 			},
	// 			{
	// 				FirstName: "John3",
	// 				LastName:  "Doe3",
	// 				Action:    "Greeting3",
	// 			},
	// 		},
	// 	},
	// 	T: time.Now(),
	// })

	log.Infof("info %v", map[string]any{
		"request": "path/to/somewhere",
		"method":  "POST",
		"body": T{
			FirstName: "John",
			LastName:  "Doe",
			Action:    "Greeting",
			Arrs: []T{
				{
					FirstName: "John1",
					LastName:  "Doe1",
					Action:    "Greeting1",
				},
				{
					FirstName: "John2",
					LastName:  "Doe2",
					Action:    "Greeting2",
				},
				{
					FirstName: "John3",
					LastName:  "Doe3",
					Action:    "Greeting3",
				},
			},
		},
		"t": time.Now(),
	})
}

type Req struct {
	Request string
	Method  string
	Body    T
	T       time.Time
}

type T struct {
	FirstName string
	LastName  string
	Action    string
	Arrs      []T
}
