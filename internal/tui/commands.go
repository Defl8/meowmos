package tui

import tea "github.com/charmbracelet/bubbletea"

type GoBackMsg struct{}

func GoBack() tea.Msg {
	return GoBackMsg{}
}
