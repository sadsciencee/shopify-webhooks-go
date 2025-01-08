package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var payloads = &cobra.Command{
	Use:   "payloads",
	Short: "payloads - download payloads",
	Long:  `Download payloads from the Shopify API reference documentation. Run this command first before running the models command.`,
	Run: func(cmd *cobra.Command, args []string) {
		payloadsFn()
	},
}

func payloadsFn() {
	fixturesDir := "./fixtures"
	if _, err := os.Stat(fixturesDir); err == nil {
		fmt.Printf("Error: %s directory already exists\n", fixturesDir)
		os.Exit(1)
	}
	if err := os.Mkdir(fixturesDir, 0755); err != nil {
		fmt.Printf("Error creating directory %s: %v\n", fixturesDir, err)
		os.Exit(1)
	}
	url := "https://shopify.dev/docs/api/webhooks/2025-01?reference=graphql#getting-started-creating-subscriptions-using-graphql-admin-api"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Error parsing HTML: %v\n", err)
		return
	}

	var processNode func(*html.Node)
	processNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "accordion-item" {
					processAccordionItem(n, fixturesDir)
					return
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processNode(c)
		}
	}

	processNode(doc)
}

func processAccordionItem(n *html.Node, fixturesDir string) {
	var topic, jsonContent string

	var findContent func(*html.Node)
	findContent = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "h3" {
				topic = getTextContent(n)
			} else if n.Data == "code" {
				jsonContent = getTextContent(n)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findContent(c)
		}
	}
	findContent(n)

	if topic != "" && jsonContent != "" {
		filename := strings.Replace(topic, "/", "_", -1) + ".json"
		if filename == "Modify payloads.json" || filename == "Filter events.json" {
			return
		}

		var prettyJSON map[string]interface{}
		if err := json.Unmarshal([]byte(jsonContent), &prettyJSON); err != nil {
			fmt.Printf("Error parsing JSON for %s: %v\n", topic, err)
			return
		}

		prettyContent, err := json.MarshalIndent(prettyJSON, "", "    ")
		if err != nil {
			fmt.Printf("Error formatting JSON for %s: %v\n", topic, err)
			return
		}

		fullPath := filepath.Join(fixturesDir, filename)

		if err := os.WriteFile(fullPath, prettyContent, 0644); err != nil {
			fmt.Printf("Error writing file %s: %v\n", filename, err)
			return
		}

		fmt.Printf("Successfully saved %s\n", filename)
	}
}

func getTextContent(n *html.Node) string {
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			text += c.Data
		}
		text += getTextContent(c)
	}
	return strings.TrimSpace(text)
}
