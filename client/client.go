package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput   textinput.Model
	lastMessage string
}

// Initial first view
func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Your Message:"
	ti.Focus()

	return model{
		textInput:   ti,
		lastMessage: "",
	}
}

// Set up of TUI
func (m model) Init() tea.Cmd {
	return nil
}

// Update - Called on every event
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.lastMessage = m.textInput.Value()
			m.textInput.Reset()
			return m, cmd
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf("Your previous message:\n%s\n\n\nEnter your Message: \n\n%s\n\n", m.lastMessage, m.textInput.View()) + "\n"
}

// Bootstrap Application
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
