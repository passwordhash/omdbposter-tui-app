package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"omdbposter/omdbapi"
	"strings"
)

const useHighPerformanceRenderer = false

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

func (m PagerModel) Init() tea.Cmd {
	return nil
}

func (m PagerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(formatContent(m.Movie, 8))
			m.ready = true

			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m PagerModel) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m PagerModel) headerView() string {
	title := titleStyle.Render(m.Movie.Title)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m PagerModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func formatContent(m omdbapi.Movie, wordsPerRow int) string {
	var plot string
	var row string
	var words = strings.Split(m.Plot, " ")
	for i, w := range words {
		row += " " + w
		if (i+1)%wordsPerRow == 0 || i == len(words)-1 {
			plot += fmt.Sprintf("--- %v\n", row)
			row = ""
		}
	}
	return fmt.Sprintf(`
### %v  // %v // %v 
--- Genre: %v

= Cast: %v
= Writer: %v

%v

*** Country: %v
*** Poster: %v

### Statistic:
# Awards & Imdb rating: %v | %v
# Collected at the box office: %v
`, m.Title, m.Year, m.ImdbRating, m.Genre, m.Actors,
		m.Writer, plot, m.Country, m.Poster, m.Actors,
		m.ImdbRating, m.BoxOffice)
}
