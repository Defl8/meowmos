package tui

import tea "github.com/charmbracelet/bubbletea"

type GoBackMsg struct{}

func GoBack() tea.Msg {
	return GoBackMsg{}
}

type SwitchViewMsg struct {
	View View
}

func SwitchView(v View) tea.Cmd {
	return func() tea.Msg {
		return SwitchViewMsg{v}
	}
}
