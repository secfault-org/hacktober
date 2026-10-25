package selector

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/secfault-org/hacktober/internal/ui/keymap"
	"github.com/secfault-org/hacktober/internal/ui/styles"
)

type Selector struct {
	*list.Model
	screen styles.Themed
	keyMap *keymap.KeyMap
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

func New(scr styles.Themed, keyMap *keymap.KeyMap, items []IdentifiableItem, delegate ItemDelegate) *Selector {
	listItems := make([]list.Item, len(items))

	for i, items := range items {
		listItems[i] = items
	}
	listModel := list.New(listItems, delegate, 0, 0)
	listModel.Styles.NoItems = scr.Theme().NoContent

	return &Selector{
		Model:  &listModel,
		screen: scr,
		keyMap: keyMap,
	}
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

func (s *Selector) Update(msg tea.Msg) (*Selector, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keyMap.Select):
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
	s.Model.SetSize(width, height)
}

func (s *Selector) SetItems(items []IdentifiableItem) tea.Cmd {
	listItems := make([]list.Item, len(items))
	for i, item := range items {
		listItems[i] = item
	}
	return s.Model.SetItems(listItems)
}
