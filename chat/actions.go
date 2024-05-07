package chat

import (
	"fmt"

	"github.com/spf13/viper"
)

type Flow struct {
	Name    string   `mapstructure:"name"`
	Actions []Action `mapstructure:"actions"`
}

type Action struct {
	Name             string            `mapstructure:"name"`
	Primary          bool              `mapstructure:"primary,omitempty"`
	NextAction       string            `mapstructure:"next_action,omitempty"`
	Description      string            `mapstructure:"description"`
	SavePayload      bool              `mapstructure:"save_payload,omitempty"`
	PayloadKey       string            `mapstructure:"payload_key,omitempty"`
	ParseFunc        string            `mapstructure:"parse_func,omitempty"`
	Regex            []string          `mapstructure:"regex,omitempty"`
	Response         string            `mapstructure:"response,omitempty"`
	IfElseNextAction *IfElseNextAction `mapstructure:"if_else_next_action,omitempty"`
	CallFunction     string            `mapstructure:"callFunction,omitempty"`
}

type IfElseNextAction struct {
	If struct {
		Regex      []string `mapstructure:"regex"`
		NextAction string   `mapstructure:"next_action"`
		Response   string   `mapstructure:"response"`
	} `mapstructure:"if"`
	Else struct {
		Regex      []string `mapstructure:"regex"`
		NextAction string   `mapstructure:"next_action"`
		Response   string   `mapstructure:"response"`
	} `mapstructure:"else"`
}

func loadFlowConfig(path string) (*[]Flow, error) {
	viper.SetConfigFile(path)
	viper.SetConfigName("chatbot")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	var flows *[]Flow
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	if err := viper.UnmarshalKey("flows", &flows); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}
	return flows, nil
}
