package printer

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-runewidth"
)

var (
	_ tea.Model = (*model)(nil)
)

type model struct {
	index   int
	grasses []string
}

func newModel(grasses []string) *model {
	return &model{
		index:   0,
		grasses: grasses,
	}
}

func (m *model) Init() tea.Cmd {
	runewidth.DefaultCondition.EastAsianWidth = false

	return m.tick()
}

func (m *model) View() string {
	var v strings.Builder

	rows := 7

	columns := (m.index + rows - 1) / rows
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			pos := c*rows + r
			if pos < m.index {
				if c > 0 {
					v.WriteRune(' ')
				}
				v.WriteString(m.grasses[pos])
			}
		}
		v.WriteRune('\n')
	}

	return v.String()
}

type tickMsg struct{}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case tickMsg:
		m.index++
		if m.index == len(m.grasses) {
			return m, tea.Quit
		}
		return m, m.tick()
	}

	return m, nil
}

func (m *model) tick() tea.Cmd {
	return tea.Tick(10*time.Millisecond, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}
