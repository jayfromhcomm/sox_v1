package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

type gpt3ResponseMsg string

type tickMsg time.Time

type model struct {
	gpt3Response string // field storing GPT-3 response
	userInput    string // field storing user input
	phase        int    // for animation
	activeTab    int    // for tabs
}

// Function to create a command that sends a gpt3ResponseMsg back to the main loop
func sendGpt3ResponseMsg(msg string) tea.Cmd {
	return func() tea.Msg {
		return gpt3ResponseMsg(msg)
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
	}
}

func (m model) Init() tea.Cmd {
	return tickCmd
}

func tickCmd() tea.Msg {
	time.Sleep(500 * time.Millisecond)
	return tickMsg(time.Now())
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.Type == tea.KeyEsc {
			return m, tea.Quit // Quit the application
		}
		switch msg.Type {
		case tea.KeyEnter:
			if strings.HasPrefix(m.userInput, "meow ") {
				specialCommand := strings.TrimPrefix(m.userInput, "meow ")
				switch specialCommand {
				case "s":
					fmt.Println("Star Command, Come in Star Command")
				case "buzz":
					fmt.Println("To Infinity and Beyond!")
				}
			} else {
				return m, m.callGPT3API(m.userInput)
			}
			m.userInput = "" // Clear the input after sending
		case tea.KeyRunes:
			m.userInput += string(msg.Runes[0])
		case tea.KeySpace:
			m.userInput += " "
		case tea.KeyBackspace:
			if len(m.userInput) > 0 {
				m.userInput = m.userInput[:len(m.userInput)-1]
			}
		}
	case tickMsg:
		m.phase++
		if m.phase > 3 {
			m.phase = 0
		}
	case gpt3ResponseMsg:
		m.gpt3Response = string(msg)
		fmt.Println("Updated gpt3Response:", m.gpt3Response)
		// Insert the new conversation into MongoDB
		InsertConversation(Conversation{
			UserInput: m.userInput,
			BotOutput: m.gpt3Response,
			Timestamp: time.Now(),
		})
		return m, nil
	}
	return m, nil
}

func (m *model) callGPT3API(prompt string) tea.Cmd {
	return func() tea.Msg {
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			return gpt3ResponseMsg("API key is empty")
		}

		url := "https://api.openai.com/v1/chat/completions" // Corrected URL
		messages := []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		}
		payload := map[string]interface{}{
			"model":    "gpt-3.5-turbo", // specify the model here
			"messages": messages,
		}
		jsonPayload, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return gpt3ResponseMsg("API call failed: " + err.Error())
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("API Response:", string(body)) // Debug print

		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			return gpt3ResponseMsg("Error unmarshalling response: " + err.Error())
		}

		if choices, ok := result["choices"].([]interface{}); ok {
			if len(choices) > 0 {
				if choice, ok := choices[0].(map[string]interface{}); ok {
					if message, ok := choice["message"].(map[string]interface{}); ok {
						if content, ok := message["content"].(string); ok {
							fmt.Println("Received text from API:", content) // Debug print
							return gpt3ResponseMsg(content)
						}
					}
				}
			}
		} else {
			return gpt3ResponseMsg("Unexpected API response format")
		}
		return gpt3ResponseMsg("No response")
	}
}

func cmd(gpt3ResponseMsg gpt3ResponseMsg) {
	panic("unimplemented")
}

func (m model) View() string {
	// Header
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).
		Padding(1)
	header := headerStyle.Render("🐱 Sox: StarCommand Personal Assistant")

	// Sidebar
	sidebarItems := []string{"Search", "History", "Settings"}
	sidebarStyle := lipgloss.NewStyle().
		Padding(1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
	sidebar := sidebarStyle.Render(lipgloss.JoinVertical(lipgloss.Left, sidebarItems...))

	// Tabs
	tabStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("205"))

	tabs := []string{"Chat", "Settings", "About"}
	for i, tab := range tabs {
		if i == m.activeTab {
			tabs[i] = tabStyle.Foreground(lipgloss.Color("205")).Render(tab)
		} else {
			tabs[i] = tabStyle.Foreground(lipgloss.Color("240")).Render(tab)
		}
	}
	tabRow := lipgloss.JoinHorizontal(lipgloss.Center, tabs...)

	// Chat History and Input
	chatHistoryStyle := lipgloss.NewStyle().
		Width(50).
		Height(15).
		Padding(1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("205"))

	chatInputStyle := lipgloss.NewStyle().
		Width(50).
		Padding(1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("205"))

	// Render GPT-3 response as Markdown
	r, _ := glamour.NewTermRenderer(
		glamour.WithStyles(glamour.ASCIIStyleConfig),
		glamour.WithWordWrap(100),
	)
	gpt3ResponseStyled, _ := r.Render(m.gpt3Response)
	fmt.Println("gpt3ResponseStyled:", gpt3ResponseStyled) // Debug print
	userInputStyled, _ := r.Render(m.userInput + "|")      // Added cursor

	chatHistory := chatHistoryStyle.Render("Sox: " + gpt3ResponseStyled)
	chatInput := chatInputStyle.Render("You: " + userInputStyled)

	// Combine chat history and input
	chatArea := lipgloss.JoinVertical(lipgloss.Left, chatHistory, chatInput)

	// Combine all UI components
	ui := []string{
		header,
		tabRow, // Moved tabs to be above the chat area
		lipgloss.JoinHorizontal(lipgloss.Left, sidebar, chatArea),
	}

	return "\n" + lipgloss.JoinVertical(lipgloss.Left, ui...)
}

func main() {
	p := tea.NewProgram(&model{})
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app:\n%s", err)
	}

	conversations := GetAllConversations()
	fmt.Println("Conversations from DB:")
	for i, conv := range conversations {
		fmt.Printf("Conversation %d: UserInput: %s, BotOutput: %s, Timestamp: %s\n", i+1, conv.UserInput, conv.BotOutput, conv.Timestamp)
	}
}
