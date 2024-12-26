package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

//////////////////////
/* TAB 1 Definition */
//////////////////////

type tabsearch struct {
	title  string
	body   string
	active bool // New field
}

func NewSearchScreen() TabModel {
	return &tabsearch{
		title: "Search",
		body:  "",
	}
}

func (t *tabsearch) Init() tea.Cmd {
	return nil
}

func (t *tabsearch) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if t.active {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "i":
				t.body += " < "
			}
		}
	}

	return t, nil
}

func (t *tabsearch) View() string {
	return fmt.Sprintf("%s\n", t.body)
}

func (t *tabsearch) Title() string {
	return t.title
}

// Add this method to tab1, tab2, and tab3
func (t *tabsearch) SetActive(isActive bool) {
	t.active = isActive
}
