package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Config é uma estrutura para representar o arquivo YAML.
type Config map[string]interface{}

// parseYaml lê o arquivo YAML e retorna um map genérico.
func parseYaml(filePath string) (Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	config := make(Config)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// generateProperties converte a estrutura do YAML para o formato de propriedades do Java.
func generateProperties(config Config, parentKey string, properties *strings.Builder) {
	for key, value := range config {
		fullKey := key
		if parentKey != "" {
			fullKey = parentKey + "." + key
		}
		switch v := value.(type) {
		case map[interface{}]interface{}:
			subConfig := make(Config)
			for k, val := range v {
				subConfig[k.(string)] = val
			}
			generateProperties(subConfig, fullKey, properties)
		case map[string]interface{}:
			generateProperties(v, fullKey, properties)
		default:
			properties.WriteString(fullKey + "=" + toString(v) + "\n")
		}
	}
}

// toString converte um valor genérico para uma string.
func toString(value interface{}) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", value), "\n", ""), "\r", ""), "\t", ""), "\b", ""), "\f", ""), "\u0000", ""))
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <path-to-yaml-file>\n", filepath.Base(os.Args[0]))
	}
	yamlPath := os.Args[1]

	config, err := parseYaml(yamlPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v\n", err)
	}

	var properties strings.Builder
	generateProperties(config, "", &properties)

	propertiesPath := strings.TrimSuffix(yamlPath, filepath.Ext(yamlPath)) + ".properties"
	err = ioutil.WriteFile(propertiesPath, []byte(properties.String()), 0644)
	if err != nil {
		log.Fatalf("Error writing properties file: %v\n", err)
	}

	log.Printf("Properties file generated: %s\n", propertiesPath)
}
