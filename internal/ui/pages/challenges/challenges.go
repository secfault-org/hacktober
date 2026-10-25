package challenges

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"github.com/secfault-org/hacktober/internal/ui/common"
	"github.com/secfault-org/hacktober/internal/ui/components/selector"
	"github.com/secfault-org/hacktober/internal/ui/keymap"
	"github.com/secfault-org/hacktober/internal/ui/styles"
)

type ChallengesPage struct {
	screen     styles.Themed
	keyMap     *keymap.KeyMap
	challenges []challenge.Challenge
	selector   *selector.Selector
}

func New(scr styles.Themed, keyMap *keymap.KeyMap, challs []challenge.Challenge) *ChallengesPage {

	list := selector.New(
		scr,
		keyMap,
		[]selector.IdentifiableItem{},
		NewItemDelegate(scr),
	)

	list.Title = "Challenges - Hacktober 2024"
	list.SetShowTitle(true)
	list.SetShowHelp(false)
	list.SetShowStatusBar(false)
	list.DisableQuitKeybindings()

	return &ChallengesPage{
		challenges: challs,
		keyMap:     keyMap,
		selector:   list,
	}
}

func (c *ChallengesPage) Init() tea.Cmd {
	items := make([]selector.IdentifiableItem, len(c.challenges))
	for i := range c.challenges {
		items[i] = Item{Challenge: &c.challenges[i]}
	}

	return tea.Batch(
		c.selector.Init(),
		c.selector.SetItems(items),
	)
}

func (c *ChallengesPage) Update(msg tea.Msg) (common.Page, tea.Cmd) {
	s, cmd := c.selector.Update(msg)
	c.selector = s
	return c, cmd
}

func (c *ChallengesPage) View() string {
	return c.selector.View()
}

func (c *ChallengesPage) getMargins() (wm, hm int) {
	wm = 2
	hm = 2
	return
}

func (c *ChallengesPage) SetSize(width, height int) {
	wm, hm := c.getMargins()
	c.selector.SetSize(width-wm, height-hm)
}

func (c *ChallengesPage) ShortHelp() []key.Binding {
	keyBindings := make([]key.Binding, 0)
	keyBindings = append(keyBindings,
		c.keyMap.UpDown,
		c.keyMap.Select,
	)

	return keyBindings
}

func (c *ChallengesPage) FullHelp() [][]key.Binding {
	bindings := [][]key.Binding{}
	bindings = append(bindings, []key.Binding{
		c.selector.KeyMap.CursorUp,
		c.selector.KeyMap.CursorDown,
		c.keyMap.Select,
	})

	return bindings
}
