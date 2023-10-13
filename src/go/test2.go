package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "i":
			// "To Infinity and Beyond!" - Trigger some action
			fmt.Println("Launching...")
		case "s":
			// "Star Command, Come in Star Command" - Maybe search?
			fmt.Println("Searching...")
		}
	}
	return m, nil
}

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
	chatStyle := lipgloss.NewStyle().
		Padding(1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("205"))
	chatArea := chatStyle.Render("Sox: Hey there! What can I do for you?")

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

func main() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app:\n%s", err)
	}
}
