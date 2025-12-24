package tui

import (
	"database/sql"
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	currentView View
	menuModel   MenuViewModel
	addModel    AddUserModel
}

func InitAppModel(db *sql.DB) AppModel {
	return AppModel{
		currentView: menuView,
		menuModel:   InitMenuViewModel(),
		addModel:    InitAddUserModel(db),
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
		case addUserView:
			m, cmd := a.addModel.Update(msg)
			a.addModel = m.(AddUserModel)
			return a, cmd
		}

	case SwitchViewMsg: // Switch to a different view
		a.switchView(msg.View)
	}

	return a, nil
}

func (a AppModel) View() string {
	switch a.currentView {
	case menuView:
		return a.menuModel.View()
	case addUserView:
		return a.addModel.View()
	}
	return ""
}

func (a AppModel) switchView(v View) {
	switch v {
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

// TODO: Implement once all views are created
// func (a AppModel) handleKeyPress(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch a.currentView {
// 	case menuView:
// 		m, cmd := a.menuModel.Update(msg)
// 		a.menuModel = m.(MenuViewModel)
// 		return a, cmd
// 	case addUserView:
// 		m, cmd := a.addModel.Update(msg)
// 		a.addModel = m.(AddUserModel)
// 		return a, cmd
// 	}
// }
