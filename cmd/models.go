package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var models = &cobra.Command{
	Use:   "models",
	Short: "models - codgen structs from payloads",
	Long:  `Generate Go structs from payloads downloaded from the Shopify API reference documentation. Run this command after running the payloads command.`,
	Run: func(cmd *cobra.Command, args []string) {
		modelsFn()
	},
}

const templateText = `package payload
{{if .NeedsTimeImport}}
import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
	"time"
)
{{else}}
import (
	"encoding/json"
	"fmt"
	"github.com/sadsciencee/shopify-webhooks-go/pkg/webhook"
)
{{end}}
func New{{.StructName}}(webhook webhook.ValidWebhook) (*{{.StructName}}, error) {
	payload := &{{.StructName}}Payload{}
	info, err := webhook.GetInfo()
	if err != nil {
		return nil, err
	}
	reader, err := webhook.GetReader()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(payload); err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}
	return &{{.StructName}}{
		info: info,
		data: *payload,
	}, nil
}

type {{.StructName}} struct {
	info webhook.Info
	data {{.StructName}}Payload
}

func (webhook *{{.StructName}}) GetInfo() (webhook.Info, error) {
	return webhook.info, nil
}
func (webhook *{{.StructName}}) GetData() ({{.StructName}}Payload, error) {
	return webhook.data, nil
}

{{.PayloadStruct}}
`

type TemplateData struct {
	StructName      string
	PayloadStruct   string
	NeedsTimeImport bool
}

func toPascalCase(s string) string {
	s = strings.ReplaceAll(s, ".", "_")
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "-", "_")

	words := strings.Split(s, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[0:1]) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, "")
}

func convertTimeFields(generated string) string {
	re := regexp.MustCompile(`(\s+[A-Za-z]+(?:_at|_date|At|Date)\s+)string(\s+.json:"[^"]+")`)
	return re.ReplaceAllString(generated, "${1}time.Time${2}")
}

func hasTimeField(data interface{}) bool {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, _ := range v {
			if strings.HasSuffix(strings.ToLower(key), "_at") ||
				strings.HasSuffix(strings.ToLower(key), "_date") ||
				strings.HasSuffix(key, "At") ||
				strings.HasSuffix(key, "Date") {
				return true
			}
		}
	}
	return false
}

func modelsFn() {
	cwd, err := os.Getwd()
	fmt.Printf("Working directory: %s\n", cwd)

	tmpl, err := template.New("webhook").Parse(templateText)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	files, err := filepath.Glob("fixtures/*.json")
	if err != nil {
		fmt.Printf("Error finding JSON files: %v\n", err)
		return
	}

	fmt.Printf("Found %d files\n", len(files))
	for _, file := range files {
		fmt.Printf("Processing file: %s\n", file)

		jsonData, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		var data interface{}
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			fmt.Printf("Error parsing JSON from %s: %v\n", file, err)
			continue
		}

		base := strings.TrimSuffix(filepath.Base(file), ".json")
		structName := toPascalCase(base)

		generatedStructs, err := gojson.Generate(bytes.NewReader(jsonData), gojson.ParseJson, structName+"Payload", "webhooks", []string{"json"}, false, true)
		if err != nil {
			fmt.Printf("Error generating struct for %s: %v\n", file, err)
			continue
		}

		// Extract just the struct definition
		structParts := strings.Split(string(generatedStructs), "\n")
		var relevantParts []string
		for _, line := range structParts {
			if !strings.Contains(line, "package") {
				relevantParts = append(relevantParts, line)
			}
		}
		payloadStruct := strings.Join(relevantParts, "\n")

		// Convert time fields
		payloadStruct = convertTimeFields(payloadStruct)

		needsTimeImport := hasTimeField(data)

		outFileName := fmt.Sprintf("topic_%s.go", base)
		outFile, err := os.Create(outFileName)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", outFileName, err)
			continue
		}

		err = tmpl.Execute(outFile, TemplateData{
			StructName:      structName,
			PayloadStruct:   payloadStruct,
			NeedsTimeImport: needsTimeImport,
		})
		if err != nil {
			fmt.Printf("Error executing template for %s: %v\n", outFileName, err)
			outFile.Close()
			continue
		}

		outFile.Close()
		fmt.Printf("Generated %s\n", outFileName)
	}
}
