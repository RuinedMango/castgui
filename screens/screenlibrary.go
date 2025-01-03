package screens

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

//////////////////////
/* TAB 1 Definition */
//////////////////////

type tablib struct {
	title  string
	body   string
	active bool // New field
}

func NewLibraryScreen() TabModel {
	return &tablib{
		title: "Library",
		body:  "Content of Tab1",
	}
}

func (t *tablib) Init() tea.Cmd {
	return nil
}

func (t *tablib) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

func (t *tablib) View() string {
	return fmt.Sprintf("%s\n", t.body)
}

func (t *tablib) Title() string {
	return t.title
}

// Add this method to tab1, tab2, and tab3
func (t *tablib) SetActive(isActive bool) {
	t.active = isActive
}
