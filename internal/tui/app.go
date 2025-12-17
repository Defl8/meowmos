package tui

import tea "github.com/charmbracelet/bubbletea"

type AppModel struct {
	currentView View
	menuModel   MenuViewModel
}

func InitAppModel() AppModel {
	return AppModel{
		currentView: menuView,
		menuModel:   InitMenuViewModel(),
	}
}

func (a AppModel) Init() tea.Cmd {
	return nil
}

func (a AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg: // key press
		switch a.currentView {
		case menuView:
			m, cmd := a.menuModel.Update(msg)
			a.menuModel = m.(MenuViewModel)
			return a, cmd
		}

	case SwitchViewMsg: // Switch to a different view
		switch msg.View {
		case menuView:
			a.currentView = menuView

		case addUserView:
			a.currentView = addUserView

		case editUserView:
			a.currentView = editUserView

		case forceSendView:
			a.currentView = forceSendView
		}
	}

	return a, nil
}

func (a AppModel) View() string {
	switch a.currentView {
	case menuView:
		return a.menuModel.View()
	}
	return ""
}
