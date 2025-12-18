package tui

import (
	"meowmos/internal/utils"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type AddUserModel struct {
	FocusPos    int
	InputFields []textinput.Model
	Keys        MultiFieldKeyMap
}

func InitAddUserModel() AddUserModel {
	m := AddUserModel{
		FocusPos:    0,
		InputFields: make([]textinput.Model, 3),
		Keys:        MultiFieldKeys(),
	}

	for i := range m.InputFields {
		t := textinput.New()

		switch i {
		case 0:
			t.Placeholder = "First Name"
			t.Focus()
		case 1:
			t.Placeholder = "Last Name"

		case 2:
			t.Placeholder = "Phone Number"
		}
		m.InputFields[i] = t
	}
	return m
}

func (m AddUserModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m AddUserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Tab):
			currentInputPos := m.FocusPos
			m.FocusPos = utils.MenuWrap(m.FocusPos+1, len(m.InputFields))
			m.InputFields[currentInputPos].Blur()
			m.InputFields[m.FocusPos].Focus()
		
		case key.Matches(msg, m.Keys.Back):
			switchViewMsg := SwitchView(menuView)
			return m, switchViewMsg
		}
	}
	var cmd tea.Cmd
	m.InputFields[m.FocusPos], cmd = m.InputFields[m.FocusPos].Update(msg)
	return m, cmd
}

func (m AddUserModel) View() string {
	var builder strings.Builder
	for i := range m.InputFields {
		builder.WriteString(m.InputFields[i].View())
		if i < len(m.InputFields)-1 {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}
