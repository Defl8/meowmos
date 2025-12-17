package tui

import (
	"fmt"
	"meowmos/internal/utils"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type MenuViewModel struct {
	Actions   []string
	CursorPos int
	Keys      KeyMap
}

func InitMenuViewModel() MenuViewModel {
	return MenuViewModel{
		Actions:   []string{"Add User", "Remove User", "Force Send"},
		CursorPos: 0,
		Keys:      DefaultKeyMap(),
	}
}

func (m MenuViewModel) Init() tea.Cmd {
	return nil
}

func (m MenuViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch {
		case key.Matches(msg, m.Keys.Up):
			m.CursorPos = utils.MenuWrap(m.CursorPos-1, len(m.Actions))

		case key.Matches(msg, m.Keys.Down):
			m.CursorPos = utils.MenuWrap(m.CursorPos+1, len(m.Actions))

		case key.Matches(msg, m.Keys.Enter):

		case key.Matches(msg, m.Keys.Back):
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MenuViewModel) View() string {
	output := "Welcome to the menu\n\n"

	for i, action := range m.Actions {
		cursor := " "
		if m.CursorPos == i {
			cursor = "->"
		}

		output += fmt.Sprintf("%s%s\n", cursor, action)
	}
	output += "\nPress q to quit\n\n"
	return output
}
