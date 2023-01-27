package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	showPage0 state = iota
	showPage1
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)
)

type model struct {
	state state
	page0 page0Model
	page1 page1Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Always pass WindowSizeMsg to all pages.
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		var model tea.Model
		model, _ = m.page0.Update(msg)
		m.page0 = model.(page0Model)
		model, _ = m.page1.Update(msg)
		m.page1 = model.(page1Model)
		return m, nil
	}

	var cmd tea.Cmd
	var model tea.Model
	switch m.state {
	case showPage0:
		model, cmd = m.page0.Update(msg)
		m.page0 = model.(page0Model)
		if m.page0.action != "" {
			m.state = showPage1
		}
		return m, cmd

	case showPage1:
		model, cmd = m.page1.Update(msg)
		m.page1 = model.(page1Model)
		// Potential logic for changing the state goes here. You change the
		// state based on how the update affected the section model.
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	switch m.state {
	case showPage0:
		return m.page0.View()
	case showPage1:
		return m.page1.View()
	default:
		return m.page0.View()
	}
}

func Run() []string {
	// TODO: Print normal help message.
	p := tea.NewProgram(model{page1: newPage1Model()})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	// TODO: Remove exit when we finish.
	os.Exit(0)
	return nil
}
