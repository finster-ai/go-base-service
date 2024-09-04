package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var AppConfig NestedConfig

type NestedConfig map[string]interface{}

// LoadConfig loads the configuration into the AppConfig variable
func LoadConfig(configFilePath string) {
	viper.SetConfigFile(configFilePath)

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal the configuration into a raw map structure
	var rawConfig map[string]interface{}
	if err := viper.Unmarshal(&rawConfig); err != nil {
		log.Fatalf("Unable to decode config into map: %v", err)
	}

	// Assign the unmarshalled map to the AppConfig variable
	AppConfig = NestedConfig(rawConfig)
}

// Get retrieves the value for a given keyPath, traversing nested maps
func (c NestedConfig) Get(keyPath string) interface{} {
	keys := strings.Split(keyPath, ".")
	var result interface{} = c

	for _, key := range keys {
		switch v := result.(type) {
		case map[string]interface{}:
			result = v[key]
		case NestedConfig:
			result = v[key]
		default:
			log.Printf("Expected map[string]interface{} but got %T for key: %s", result, key)
			return nil
		}
	}

	return result
}

// GetString retrieves a string value for the given keyPath
func (c NestedConfig) GetString(keyPath string) string {
	value := c.Get(keyPath)
	if value == nil {
		log.Printf("String value not found for key: %s", keyPath)
		return ""
	}
	strValue, ok := value.(string)
	if !ok {
		log.Printf("Expected string but got %T for key: %s", value, keyPath)
		return ""
	}
	return strValue
}

// Other methods (GetInt, GetStringSlice, GetStringMap) remain the same...
