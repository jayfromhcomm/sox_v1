package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	progress progress.Model
	percent  float64
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	case string:
		if msg == "tick" {
			m.percent += 2 // Now you can modify it
			if m.percent >= 100 {
				return m, tea.Quit
			}
		}
	}
	return m, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg { return "tick" })
}

func (m *model) View() string {
	boxStyle := lipgloss.NewStyle().
		Padding(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("205")).
		Background(lipgloss.Color("15"))

	catStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("226"))

	renderedProgress := m.progress.View()

	ui := []string{
		boxStyle.Render(
			fmt.Sprintf("⚡ %s 🐱\n%s",
				catStyle.Render("Assistant: Fill in your text prompt here"),
				renderedProgress,
			),
		),
	}

	return "\n" + lipgloss.JoinVertical(lipgloss.Center, ui...)
}

func main() {
	p := progress.NewModel(progress.WithDefaultGradient())
	m := model{progress: p}
	program := tea.NewProgram(&m)
	if err := program.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app:\n%s", err)
	}
}
