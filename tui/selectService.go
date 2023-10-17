package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"omdbposter/omdbapi"
	"strings"
)

func RunSelect(options []omdbapi.MovieSearched, header string) SelectModel {
	return SelectModel{
		Header:  header,
		choices: options,
	}
}

func (m SelectModel) Init() tea.Cmd {
	return nil
}

func (m SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.SelectCode = -1
			return m, tea.Quit

		case "b":
			m.SelectCode = 0
			return m, tea.Quit

		case "enter":
			m.SelectCode = 1
			m.Choice = m.choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m SelectModel) View() string {
	s := strings.Builder{}
	s.WriteString(m.Header + "\n\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i].Title)
		s.WriteString("\n")
	}
	s.WriteString("(b чтобы изменить / ctrl+c чтобы выйти)\n")

	return s.String()
}
