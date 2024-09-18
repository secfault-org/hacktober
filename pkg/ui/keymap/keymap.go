package keymap

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit   key.Binding
	UpDown key.Binding
	Select key.Binding
	Back   key.Binding
	Help   key.Binding
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

	keymap.Select = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
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
