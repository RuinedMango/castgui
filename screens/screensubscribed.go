package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

//////////////////////
/* TAB 1 Definition */
//////////////////////

type tabsub struct {
	title  string
	body   string
	active bool // New field
}

func NewSubScreen() TabModel {
	return &tabsub{
		title: "Subscribed",
		body:  "Content of Tab1",
	}
}

func (t *tabsub) Init() tea.Cmd {
	return nil
}

func (t *tabsub) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

func (t *tabsub) View() string {
	return fmt.Sprintf("%s\n", t.body)
}

func (t *tabsub) Title() string {
	return t.title
}

// Add this method to tab1, tab2, and tab3
func (t *tabsub) SetActive(isActive bool) {
	t.active = isActive
}
