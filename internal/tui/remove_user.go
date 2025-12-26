package tui

import (
	"database/sql"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type RemoveUserModel struct {
	UserList list.Model
	Keys     ListKeyMap
	Database *sql.DB
	Status   string
}


func InitRemoveUserModel(db *sql.DB) RemoveUserModel {
	 
}


func (m RemoveUserModel) Init() tea.Cmd {
	return nil
}


func (m RemoveUserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}


func (m RemoveUserModel) View() string {
	return ""
}


func (m RemoveUserModel) handleKeyPress(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch {
		case key.Matches(msg, m.Keys.Up):
	}
}
