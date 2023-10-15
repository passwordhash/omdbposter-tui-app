package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

//func main() {
//	p := tea.NewProgram(initialModel())
//	if _, err := p.RunSelect(); err != nil {
//		log.Fatal(err)
//	}
//}

type (
	errMsg error
)

func RunInput(placeholder string, header string) InputModel {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return InputModel{
		Header:    header,
		TextInput: ti,
		err:       nil,
	}
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if len(m.TextInput.Value()) > 3 {
				return m, tea.Quit
			}
			m.Header = "Вы ввели слишком короткое название"
		case tea.KeyCtrlC, tea.KeyEsc:
			m.IsExit = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m InputModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%v\n\n%s",
		m.Header,
		m.TextInput.View(),
		"(esc чтобы выйти)",
	) + "\n"
}
