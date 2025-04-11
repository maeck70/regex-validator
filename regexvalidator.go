package regexvalidator

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

// ValidateRegex validates a value against a regex validator
// Returns true if the value matches the regex, false otherwise
// Example usage:
//
//	// Load the regex validators from a YAML file
//	Load("validators.yaml")
//
//	// Validate a value against a regex validator
//	res := ValidateRegex("ipv4", "10.20.30.40")
func ValidateRegex(validator string, value string) bool {
	return compiledRegexs[validator].MatchString(value)
}

// Load loads the regex validators from a YAML file
// The YAML file should contain a list of validators with their regex patterns
// Example YAML file:
//
//	validator:
//		validator1:
//			description: "Validator 1"
//			re: "^[a-zA-Z0-9]+$"
func Load(fileName string) {
	var validators validators_t

	err := loadValidators(fileName, &validators)
	if err != nil {
		log.Printf("Error loading validators: %v\n", err)
		return
	}

	compiledRegexs = make(map[string]*regexp.Regexp)

	for v := range validators.Validator {
		log.Printf("Loading RegEx validator: %s\n", v)
		compiledRegexs[v] = regexp.MustCompile(validators.Validator[v].Re)
	}
}

// IsLoaded checks if a regex validator is loaded
// Returns true if the validator is loaded, false otherwise
func IsLoaded(rekey string) bool {
	_, ok := compiledRegexs[rekey]
	return ok
}

var compiledRegexs map[string]*regexp.Regexp

type validator_t struct {
	Description string `yaml:"description"`
	Re          string `yaml:"re"`
}

type validators_t struct {
	Validator map[string]validator_t `yaml:"validators"`
}

func loadValidators(fileName string, validators *validators_t) error {

	// Load yaml file
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %s", err)
	}

	// Unmarshal yaml file
	err = yaml.Unmarshal(yamlFile, validators)
	if err != nil {
		return fmt.Errorf("error unmarshalling YAML file: %s", err)
	}

	return nil
}
