package printer

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-runewidth"
)

var (
	_ tea.Model = (*model)(nil)
)

type model struct {
	printer *Printer

	index   int
	grasses []string
}

func newModel(p *Printer, grasses []string) *model {
	return &model{
		printer: p,

		index:   0,
		grasses: grasses,
	}
}

func (m *model) Init() tea.Cmd {
	runewidth.DefaultCondition.EastAsianWidth = false

	return m.tick()
}

func (m *model) View() string {
	return m.printer.grasses(m.grasses, m.index)
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
