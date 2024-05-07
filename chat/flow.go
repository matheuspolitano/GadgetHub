package chat

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type FlowName string

var (
	ProductReview FlowName = "product_review"
)

type Template struct {
	Version string `mapstructure:"version"`
	Flows   []Flow `mapstructure:"flows"`
}

func (t *Template) GetFlow(flowName FlowName) (*Flow, error) {
	for _, i := range t.Flows {
		if i.Name == string(flowName) {
			return &i, nil
		}
	}
	return nil, fmt.Errorf("error: not found flowname %s", flowName)

}

type Flow struct {
	Name    string   `mapstructure:"name"`
	Actions []Action `mapstructure:"actions"`
}

func (t *Flow) GetPrimaryAction() (*Action, error) {
	for _, i := range t.Actions {
		if i.Primary {
			return &i, nil
		}
	}
	return nil, errors.New("error: not found primary action")

}

func loadChatTemplate(path string) (*Template, error) {
	viper.SetConfigFile(path)
	viper.SetConfigName("chatbot")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	var template *Template
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&template); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %v", err)
	}
	return template, nil
}
