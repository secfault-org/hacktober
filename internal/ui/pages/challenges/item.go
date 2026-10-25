package challenges

import (
	"fmt"
	"io"
	"strings"
	"time"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"github.com/secfault-org/hacktober/internal/ui/styles"
)

type Item struct {
	Challenge *challenge.Challenge
}

func (i Item) ID() string          { return i.Challenge.Id }
func (i Item) Title() string       { return i.Challenge.Name }
func (i Item) Description() string { return i.Challenge.Description }
func (i Item) FilterValue() string { return i.Challenge.Name }

type ItemDelegate struct {
	screen styles.Themed
}

func NewItemDelegate(scr styles.Themed) *ItemDelegate {
	return &ItemDelegate{
		screen: scr,
	}
}

func (d ItemDelegate) Height() int {
	return d.screen.Theme().ChallengeListItemContainer.GetVerticalFrameSize() +
		d.screen.Theme().ChallengeListItemContainer.GetHeight()
}

func (d ItemDelegate) Spacing() int { return 1 }

func (d ItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item := listItem.(Item)
	stringBuilder := strings.Builder{}

	var isSelected = index == m.Index()

	itemStyle := d.screen.Theme().ChallengeListItem.Normal
	if isSelected {
		itemStyle = d.screen.Theme().ChallengeListItem.Active
	}

	title := item.Title()
	if item.Challenge.Locked() {
		title += " 🔒"
	}

	releaseStr := fmt.Sprintf(" Release: %s", item.Challenge.ReleaseDate.Format(time.DateOnly))
	if m.Width()-itemStyle.Base.GetHorizontalFrameSize()-lipgloss.Width(releaseStr) <= 0 {
		releaseStr = ""
	}
	releaseStyle := itemStyle.ReleaseDate.
		Align(lipgloss.Right).
		Width(m.Width() - itemStyle.Base.GetHorizontalFrameSize() - lipgloss.Width(title))

	release := releaseStyle.Render(releaseStr)
	title = itemStyle.Title.Render(title)
	desc := item.Description()
	desc = itemStyle.Desc.Render(desc)

	stringBuilder.WriteString(lipgloss.JoinHorizontal(lipgloss.Bottom, title, release))
	stringBuilder.WriteRune('\n')
	stringBuilder.WriteString(desc)
	fmt.Fprint(w, itemStyle.Base.Render(stringBuilder.String()))
}
