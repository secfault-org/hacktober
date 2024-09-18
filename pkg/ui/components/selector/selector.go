package selector

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/secfault-org/hacktober/pkg/ui/common"
)

type Selector struct {
	*list.Model
	common common.Common
	active int
}

type IdentifiableItem interface {
	list.DefaultItem
	ID() string
}

type ItemDelegate interface {
	list.ItemDelegate
}

type SelectMsg struct{ IdentifiableItem }

func New(common common.Common, items []IdentifiableItem, delegate ItemDelegate) *Selector {
	listItems := make([]list.Item, len(items))

	for i, items := range items {
		listItems[i] = items
	}
	listModel := list.New(listItems, delegate, common.Width, common.Height)
	listModel.Styles.NoItems = common.Styles.NoContent
	selector := &Selector{Model: &listModel, common: common}
	selector.SetSize(common.Width, common.Height)
	return selector
}

func (s *Selector) SelectedItem() IdentifiableItem {
	item := s.Model.SelectedItem()
	i, ok := item.(IdentifiableItem)
	if !ok {
		return nil
	}
	return i
}

func (s *Selector) SelectItemCmd() tea.Msg {
	return SelectMsg{s.SelectedItem()}
}

func (s *Selector) Init() tea.Cmd { return nil }

func (s *Selector) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.common.KeyMap.Select):
			cmds = append(cmds, s.SelectItemCmd)
		}
	}
	m, cmd := s.Model.Update(msg)

	s.Model = &m

	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	s.active = s.Index()
	return s, tea.Batch(cmds...)
}

func (s *Selector) SetSize(width, height int) {
	s.common.SetSize(width, height)
	s.Model.SetSize(width, height)
}
