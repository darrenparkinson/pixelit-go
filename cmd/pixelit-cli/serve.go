package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type AutomationType uint8

const (
	Clock AutomationType = iota + 1
	Instagram
	TikTok
)

var (
	AutomationTypeValue = map[string]uint8{
		"clock":     1,
		"instagram": 2,
		"tiktok":    3,
	}
	AutomationTypeName = map[uint8]string{
		1: "clock",
		2: "instagram",
		3: "tiktok",
	}
)

func (t AutomationType) String() string {
	return AutomationTypeName[uint8(t)]
}

func (t AutomationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *AutomationType) UnmarshalJSON(data []byte) (err error) {
	var types string
	if err := json.Unmarshal(data, &types); err != nil {
		return err
	}
	if *t, err = ParseAutomationType(types); err != nil {
		return err
	}
	return nil
}

// AutomationTypeHookFunc allows us to unmarshal our custom type with mapstructure
// Great article here explains this: https://sagikazarmark.hu/blog/decoding-custom-formats-with-viper/
func AutomationTypeHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type, // data type
		t reflect.Type, // target data type
		data interface{}, // raw data
	) (interface{}, error) {
		// Check if the data type matches the expected one
		// in this case we're expecting a string really so
		// if it's not that, then just return
		if f.Kind() != reflect.String {
			return data, nil
		}
		// Check if the target type matches the expected one
		// which in our case is an AutomationType
		if t != reflect.TypeOf(AutomationType(0)) {
			return data, nil
		}
		// Format/decode/parse the data and return the new value
		at, err := ParseAutomationType(data.(string))
		if err != nil {
			return data, err
		}
		return at, nil
	}
}

type Automation struct {
	GlobalDuration    int                `mapstructure:"globalDuration"`
	AutomationScreens []AutomationScreen `mapstructure:"screens"`
}

type AutomationScreen struct {
	Name     string         `mapstructure:"name"`
	Type     AutomationType `mapstructure:"type" json:"type"`
	Duration int            `mapstructure:"duration"`
	Username string         `mapstructure:"username"`
}

func ParseAutomationType(s string) (AutomationType, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := AutomationTypeValue[s]
	if !ok {
		return AutomationType(0), fmt.Errorf("%q is not a valid automation type", s)
	}
	return AutomationType(value), nil
}

func (s *AutomationScreen) Run() {
	switch s.Type {
	case Clock:
		pxc.SendClock()
	case Instagram:
		pxc.SendInstagramFollowers(s.Username)
	case TikTok:
		pxc.SendTiktokFollowers(s.Username)
	}
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Read configuration and perform timed actions",
	Run:   serveActions,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serveActions(cmd *cobra.Command, args []string) {
	var automation Automation
	err := viper.UnmarshalKey("automation", &automation, viper.DecodeHook(AutomationTypeHookFunc()))
	if err != nil {
		log.Fatal("error loading automation configuration", err)
	}
	if automation.GlobalDuration == 0 {
		log.Fatal("global duration must be set to run this command")
	}

	// ticker := time.NewTicker(time.Duration(automation.GlobalDuration) * time.Second)
	done := make(chan bool, 1) // send to this when we get interrupt and call ticker.Stop()
	screens := make(chan AutomationScreen)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		log.Println("caught signal", s.String(), "stopping...")
		// ticker.Stop()
		done <- true
		log.Println("ticker stopped")
	}()

	screenCount := len(automation.AutomationScreens)
	log.Println("added", screenCount, "screens to cycle, with global duration of", automation.GlobalDuration, "seconds")
	go func(automationScreens []AutomationScreen) {
		currentScreen := 0
		for {
			screen := automationScreens[currentScreen]
			currentScreen++
			if currentScreen == len(automationScreens) {
				currentScreen = 0
			}
			screens <- screen
		}
	}(automation.AutomationScreens)

	go func(globalDuration int) {
		for {
			screen := <-screens
			log.Println("running", screen.Name, screen.Type, screen.Username, screen.Duration)
			screen.Run()
			if screen.Duration > 0 {
				log.Println("sleeping for", screen.Duration, "seconds")
				time.Sleep(time.Duration(screen.Duration) * time.Second)
			} else {
				log.Println("sleeping for global screen duration of", globalDuration, "seconds")
				time.Sleep(time.Duration(globalDuration) * time.Second)
			}

		}
	}(automation.GlobalDuration)

	<-done // wait until we're done

}
