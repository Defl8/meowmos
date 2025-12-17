package tui

import tea "github.com/charmbracelet/bubbletea"

type currentView int

const (
	menuView currentView = iota
	addUserView
	editUserView
	forceSendView
)

type AppModel struct {
	currentView currentView
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
	switch a.currentView {
	case menuView:
		m, cmd := a.menuModel.Update(msg)
		a.menuModel = m.(MenuViewModel)
		return a, cmd
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
