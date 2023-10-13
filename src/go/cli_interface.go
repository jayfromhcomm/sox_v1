// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	"github.com/charmbracelet/bubbles/progress"
// 	"github.com/charmbracelet/bubbles/textinput"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// 	"github.com/go-resty/resty/v2"
// )

// type GPT3Response struct {
// 	Choices []struct {
// 		Text string `json:"text"`
// 	} `json:"choices"`
// }

// type model struct {
// 	progressBar  progress.Model
// 	percent      float64
// 	textInput    textinput.Model
// 	gpt3Response string
// }

// func getGPT3Response(prompt string) (string, error) {
// 	client := resty.New()
// 	apiKey := os.Getenv("API_KEY") // Make sure to set this environment variable

// 	resp, err := client.R().
// 		SetHeader("Authorization", "Bearer "+apiKey).
// 		SetBody(map[string]interface{}{"prompt": prompt, "max_tokens": 50}).
// 		Post("https://api.openai.com/v1/engines/davinci-codex/completions")

// 	if err != nil {
// 		return "", err
// 	}

// 	var gpt3Resp GPT3Response
// 	err = json.Unmarshal(resp.Body(), &gpt3Resp)
// 	if err != nil {
// 		return "", err
// 	}

// 	return gpt3Resp.Choices[0].Text, nil
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd

// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "q":
// 			return m, tea.Quit
// 		case "Enter":
// 			resp, err := getGPT3Response(m.textInput.Value())
// 			if err != nil {
// 				// Handle error
// 			}
// 			m.gpt3Response = resp
// 		}
// 	}

// 	// Update the text input
// 	m.textInput, cmd = m.textInput.Update(msg)

// 	return m, cmd
// }

// func (m *model) View() string {
// 	// Style the text input
// 	inputView := m.textInput.View()

// 	// Style the box
// 	boxStyle := lipgloss.NewStyle().
// 		Padding(1).
// 		Border(lipgloss.RoundedBorder()).
// 		BorderForeground(lipgloss.Color("205")).
// 		Background(lipgloss.Color("15"))

// 	catStyle := lipgloss.NewStyle().
// 		Foreground(lipgloss.Color("226"))

// 	renderedProgress := m.progressBar.View()

// 	var content string
// 	if m.gpt3Response == "" {
// 		// Show loading bar
// 		content = renderedProgress
// 	} else {
// 		// Show GPT-3 response
// 		content = m.gpt3Response
// 	}

// 	ui := []string{
// 		boxStyle.Render(
// 			fmt.Sprintf("⚡ %s 🐱\n%s\n%s",
// 				catStyle.Render("Sox: Hey there! What can I do for you?"),
// 				content,
// 				inputView,
// 			),
// 		),
// 	}

// 	return "\n" + lipgloss.JoinVertical(lipgloss.Center, ui...)
// }

// func main() {
// 	p := progress.NewModel(progress.WithDefaultGradient())
// 	input := textinput.NewModel()
// 	input.Placeholder = "Type here..."
// 	input.Focus()
// 	input.CharLimit = 156
// 	input.Width = 20

// 	m := model{
// 		progressBar: p,
// 		textInput:   input,
// 	}

// 	program := tea.NewProgram(&m)
// 	if err := program.Start(); err != nil {
// 		fmt.Fprintf(os.Stderr, "Error starting app:\n%s", err)
// 	}
// }
