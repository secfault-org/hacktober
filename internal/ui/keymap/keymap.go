package keymap

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit           key.Binding
	UpDown         key.Binding
	GotoTop        key.Binding
	GotoBottom     key.Binding
	Select         key.Binding
	Submit         key.Binding
	EnterFlag      key.Binding
	SpawnContainer key.Binding
	StopContainer  key.Binding
	Back           key.Binding
	Help           key.Binding
}

func DefaultKeyMap() *KeyMap {
	keymap := new(KeyMap)

	keymap.Quit = key.NewBinding(
		key.WithKeys(
			"q",
			"ctrl+c",
		),
		key.WithHelp(
			"q",
			"quit",
		),
	)

	keymap.UpDown = key.NewBinding(
		key.WithKeys(
			"up",
			"down",
			"k",
			"j",
		),
		key.WithHelp(
			"↑↓ j/k",
			"navigate",
		),
	)

	keymap.GotoTop = key.NewBinding(
		key.WithKeys(
			"home",
			"g",
		),
		key.WithHelp(
			"g/home",
			"goto top",
		),
	)

	keymap.GotoBottom = key.NewBinding(
		key.WithKeys(
			"end",
			"G",
		),
		key.WithHelp(
			"G/end",
			"goto bottom",
		),
	)

	keymap.Select = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	)

	keymap.Submit = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit flag"),
	)

	keymap.EnterFlag = key.NewBinding(
		key.WithKeys("ctrl+f"),
		key.WithHelp("ctrl+f", "Enter Flag"),
	)

	keymap.SpawnContainer = key.NewBinding(
		key.WithKeys(
			"ctrl+r",
		),
		key.WithHelp(
			"ctrl+r",
			"spawn container",
		),
	)

	keymap.StopContainer = key.NewBinding(
		key.WithKeys(
			"ctrl+t",
		),
		key.WithHelp(
			"ctrl+t",
			"stop container",
		),
	)

	keymap.Back = key.NewBinding(
		key.WithKeys(
			"esc",
		),
		key.WithHelp(
			"esc",
			"back",
		),
	)

	keymap.Help = key.NewBinding(
		key.WithKeys(
			"?",
		),
		key.WithHelp(
			"?",
			"toggle help",
		),
	)

	return keymap
}
