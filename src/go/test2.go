// The above code is a Go program that creates a command-line interface (CLI) for a personal assistant
// called "Sox". It uses the GPT-3 API to generate responses to user prompts and displays them in the
// CLI interface. The program also includes various UI elements such as a header, sidebar, tabs, chat
// area, content viewport, progress bar, and help keys.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load() // Load .env file
}

type model struct {
	gpt3Response string // field storing gpt response
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter: // Handle 'Enter' key
			go m.callGPT3API("Your prompt here") // Replace with actual user input
		case tea.KeyRunes:
			switch msg.Runes[0] {
			case 'q':
				return m, tea.Quit
			case 's':
				fmt.Println("Star Command, Come in Star Command") // Replace with actual functionality
			case 'i':
				fmt.Println("To Infinity and Beyond!") // Replace with actual functionality
			}
		}
	}
	return m, nil
}

// New function to make the GPT-3 API call
func (m *model) callGPT3API(prompt string) {
	apiKey := os.Getenv("API_KEY") // Read from .env file
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"

	payload := map[string]string{
		"prompt":     prompt,
		"max_tokens": "50",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Handle error
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	m.gpt3Response = result["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)
}

// View style
func (m model) View() string {
	// Header
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).
		Padding(1)
	header := headerStyle.Render("🐱 Sox: StarCommand Personal Assistant")

	// Sidebar with Search and History
	sidebarStyle := lipgloss.NewStyle().
		Padding(1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
	sidebar := sidebarStyle.Render("Search:\n\nHistory:")

	// Description and Graphic
	descStyle := lipgloss.NewStyle().
		Padding(1)
	desc := descStyle.Render("Your friendly assistant.")
	graphic := descStyle.Render("Animated Graphic Here")

	// Help Keys
	helpKeysStyle := lipgloss.NewStyle().
		Padding(1)
	helpKeys := helpKeysStyle.Render("Help: q to quit | i for Infinity | s for Star Command")

	// Tabs
	activeTabStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	inactiveTabStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	tabs := []string{"Chat", "Settings", "About"}
	for i, tab := range tabs {
		if i == 0 {
			tabs[i] = activeTabStyle.Render(tab)
		} else {
			tabs[i] = inactiveTabStyle.Render(tab)
		}
	}
	tabRow := lipgloss.JoinHorizontal(lipgloss.Center, tabs...)

	// Chat Area
	// chatStyle := lipgloss.NewStyle().
	// 	Padding(1).
	// 	Border(lipgloss.NormalBorder()).
	// 	BorderForeground(lipgloss.Color("205"))
	chatArea := lipgloss.NewStyle().
		Padding(1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("205")).
		Render("Sox: " + m.gpt3Response) // Modified line

	// Content Viewport
	viewportStyle := lipgloss.NewStyle().
		Padding(1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
	viewport := viewportStyle.Render("Content goes here.")

	// Progress Bar
	progressStyle := lipgloss.NewStyle().
		Padding(1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
	progressBar := progressStyle.Render("Progress Bar Here")

	// Putting it all together
	ui := []string{
		header,
		lipgloss.JoinHorizontal(lipgloss.Left, desc, graphic),
		tabRow,
		lipgloss.JoinHorizontal(lipgloss.Left, sidebar, chatArea),
		viewport,
		helpKeys,
		progressBar,
	}

	return "\n" + lipgloss.JoinVertical(lipgloss.Left, ui...)
}

func Padding(i int) {
	panic("unimplemented")
}

func main() {
	p := tea.NewProgram(&model{})
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app:\n%s", err)
	}
}
