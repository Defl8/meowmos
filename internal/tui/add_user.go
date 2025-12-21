package tui

import (
	"database/sql"
	"meowmos/internal/database"
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
	Database    *sql.DB
}

func (m AddUserModel) ClearInputs() {
	for i := range m.InputFields {
		m.InputFields[i].SetValue("")
	}
}

func InitAddUserModel(db *sql.DB) AddUserModel {
	m := AddUserModel{
		FocusPos:    0,
		InputFields: make([]textinput.Model, 3),
		Keys:        MultiFieldKeys(),
		Database:    db,
	}

	for i := range m.InputFields {
		t := textinput.New()

		switch i {
		case 0:
			t.Placeholder = "First Name"
			t.Focus()
			t.CharLimit = 32
			t.Width = t.CharLimit + 5

		case 1:
			t.Placeholder = "Last Name"
			t.CharLimit = 32
			t.Width = t.CharLimit + 5

		case 2:
			t.Placeholder = "Phone Number"
			t.CharLimit = 32
			t.Width = t.CharLimit + 5
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

		case key.Matches(msg, m.Keys.Save):
			user := database.NewUser(m.InputFields[0].Value(), m.InputFields[1].Value(), m.InputFields[2].Value())
			err := database.AddUser(m.Database, *user)
			if err != nil {
				return m, nil
			}

			m.ClearInputs()

			return m, SwitchView(menuView)


		case key.Matches(msg, m.Keys.Back):
			m.ClearInputs()
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

	// headers for the fields
	headers := []string{"Enter your first name", "Enter your last name", "Enter your phone number"}
	for i := range m.InputFields {
		builder.WriteString(headers[i])
		builder.WriteRune('\n')
		builder.WriteString(m.InputFields[i].View())
		if i < len(m.InputFields)-1 {
			builder.WriteRune('\n')
		}
	}

	builder.WriteString("\n\nPress ctrl+s to save the user\nPress q to go back\n")
	return builder.String()
}
