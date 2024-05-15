package chat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONToMap(t *testing.T) {
	// Define test cases
	tests := []struct {
		jsonStr     string
		expectedMap map[string]string
		expectError bool
	}{
		{
			jsonStr:     `{"name": "John", "age": "30", "city": "New York"}`,
			expectedMap: map[string]string{"name": "John", "age": "30", "city": "New York"},
			expectError: false,
		},
		{
			jsonStr:     `{"key1": "value1", "key2": "value2"}`,
			expectedMap: map[string]string{"key1": "value1", "key2": "value2"},
			expectError: false,
		},
		{
			jsonStr:     `{"singleKey": "singleValue"}`,
			expectedMap: map[string]string{"singleKey": "singleValue"},
			expectError: false,
		},
		{
			jsonStr:     `{"nested": {"innerKey": "innerValue"}}`,
			expectedMap: nil,
			expectError: true,
		},
		{
			jsonStr:     `{invalid json}`,
			expectedMap: nil,
			expectError: true,
		},
	}

	for _, test := range tests {
		result, err := JSONToMap(test.jsonStr)

		if test.expectError {
			assert.Error(t, err, "Expected an error but got nil")
		} else {
			assert.NoError(t, err, "Did not expect an error but got one")
			assert.Equal(t, test.expectedMap, result, "Expected map does not match result map")
		}
	}
}
func TestMapToJSON(t *testing.T) {
	// Define test cases
	tests := []struct {
		inputMap     map[string]string
		expectedJSON string
		expectError  bool
	}{
		{
			inputMap:     map[string]string{"name": "John", "age": "30", "city": "New York"},
			expectedJSON: `{"age":"30","city":"New York","name":"John"}`,
			expectError:  false,
		},
		{
			inputMap:     map[string]string{"key1": "value1", "key2": "value2"},
			expectedJSON: `{"key1":"value1","key2":"value2"}`,
			expectError:  false,
		},
		{
			inputMap:     map[string]string{"singleKey": "singleValue"},
			expectedJSON: `{"singleKey":"singleValue"}`,
			expectError:  false,
		},
	}

	for _, test := range tests {
		result, err := MapToJSON(test.inputMap)

		if test.expectError {
			assert.Error(t, err, "Expected an error but got nil")
		} else {
			assert.NoError(t, err, "Did not expect an error but got one")
			assert.JSONEq(t, test.expectedJSON, result, "Expected JSON string does not match result JSON string")
		}
	}
}
