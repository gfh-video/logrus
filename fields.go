package logrus

import (
	"strings"

	cases "golang.org/x/text/cases"
	language "golang.org/x/text/language"

	structs_to_map "github.com/fatih/structs"
)

// Common global fields that will be set in every log
var GlobalFields Fields = make(Fields)

// Converts snake_case or PascalCase to camelCase
func toCamelCase(input string) string {
	// Split the input string into words using underscores as separators
	words := strings.Split(input, "_")

	if len(words) > 0 {

		// Capitalize the first letter of each word (except the first one)
		for i := 1; i < len(words); i++ {
			caser := cases.Title(language.Und)
			words[i] = caser.String(words[i])
		}

		// Join the words to create the camelCase string
		joinedWords := strings.Join(words, "")

		if len(joinedWords) > 0 {
			return strings.ToLower(joinedWords[0:1]) + joinedWords[1:]
		}
	}

	return ""
}

// Call to create empty Fields object
func NewFields_Empty() Fields {
	return Fields{}
}

// Call to create Fields from any struct object, using camelCase mapping
func NewFields(anyObject interface{}) Fields {

	logFields := Fields{}

	// Check if this is a map of string to string
	map_string_int, ok := anyObject.(map[string]int)
	if ok {
		for key, val := range map_string_int {
			// if key != "" {
			logFields[toCamelCase(key)] = val
			// }
		}
		return logFields
	}

	// Check if this is a map of string to string
	map_string_string, ok := anyObject.(map[string]string)
	if ok {
		for key, val := range map_string_string {
			// if key != "" {
			logFields[toCamelCase(key)] = val
			// }
		}
		return logFields
	}

	// Check if this is a map of string to string
	map_string_bool, ok := anyObject.(map[string]bool)
	if ok {
		for key, val := range map_string_bool {
			logFields[toCamelCase(key)] = val
		}
		return logFields
	}

	// This is a struct
	for key, val := range structs_to_map.Map(anyObject) {
		logFields[toCamelCase(key)] = val
	}
	return logFields
}

func (x *Fields) AddFields(extraFields map[string]interface{}) {
	for key, val := range extraFields {
		(*x)[toCamelCase(key)] = val
	}
}
