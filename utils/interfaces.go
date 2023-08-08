package utils

import (
	"gopkg.in/yaml.v3"
	"sort"
)

func MapValueToYAMLString(mapInterface interface{}) (string, error) {
	data, err := yaml.Marshal(mapInterface)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func SortEntries(input []byte) map[string]interface{} {
	// Unmarshal the YAML content into a map
	data := make(map[string]interface{})
	err := yaml.Unmarshal(input, &data)
	if err != nil {
		panic(err)
	}

	// Sort the keys
	sortedKeys := sortMapKeys(data)

	// Create a new map with sorted keys
	sortedData := make(map[string]interface{})
	for _, key := range sortedKeys {
		sortedData[key] = data[key]
	}

	// Print the sorted map
	return sortedData
}

func sortMapKeys(entries map[string]interface{}) []string {
	keys := make([]string, 0, len(entries))

	// Get keys from Map
	for entry := range entries {
		keys = append(keys, entry)
	}

	sort.Strings(keys)

	return keys
}
