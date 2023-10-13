package main

import (
	"fmt"
	"os"
	
	
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	progressBar progress.Model
	percent     float64
	textInput   textinput.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}

	// Update the text input
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m *model) View() string {
	// Style the text input
	inputView := m.textInput.View()

	// Style the box
	boxStyle := lipgloss.NewStyle().
		Padding(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("205")).
		Background(lipgloss.Color("15"))

	catStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("226"))

	renderedProgress := m.progressBar.View()

	ui := []string{
		boxStyle.Render(
			fmt.Sprintf("⚡ %s 🐱\n%s\n%s",
				catStyle.Render("Sox: Hey there! What can I do for you?"),
				renderedProgress,
				inputView,
			),
		),
	}

	return "\n" + lipgloss.JoinVertical(lipgloss.Center, ui...)
}

func main() {
	p := progress.NewModel(progress.WithDefaultGradient())
	input := textinput.NewModel()
	input.Placeholder = "Type here..."
	input.Focus()
	input.CharLimit = 156
	input.Width = 20

	m := model{
		progressBar: p,
		textInput:   input,
	}

	program := tea.NewProgram(&m)
	if err := program.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app:\n%s", err)
	}
}
