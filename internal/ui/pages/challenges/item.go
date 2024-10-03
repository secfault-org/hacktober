package challenges

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/secfault-org/hacktober/internal/model/challenge"
	"github.com/secfault-org/hacktober/internal/ui/common"
	"io"
	"strings"
	"time"
)

type Item struct {
	Challenge *challenge.Challenge
}

func (i Item) ID() string          { return i.Challenge.Id }
func (i Item) Title() string       { return i.Challenge.Name }
func (i Item) Description() string { return i.Challenge.Description }
func (i Item) FilterValue() string { return i.Challenge.Name }

type ItemDelegate struct {
	common *common.Common
}

func NewItemDelegate(common *common.Common) *ItemDelegate {
	return &ItemDelegate{
		common: common,
	}
}

func (d ItemDelegate) Height() int {
	return d.common.Styles.ChallengeListItemContainer.GetVerticalFrameSize() + d.common.Styles.ChallengeListItemContainer.GetHeight()
}

func (d ItemDelegate) Spacing() int { return 1 }

func (d ItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item := listItem.(Item)
	stringBuilder := strings.Builder{}

	var isSelected = index == m.Index()

	styles := d.common.Styles.ChallengeListItem.Normal
	if isSelected {
		styles = d.common.Styles.ChallengeListItem.Active
	}

	title := item.Title()
	if item.Challenge.Locked() {
		title += " 🔒"
	}

	releaseStr := fmt.Sprintf(" Release: %s", item.Challenge.ReleaseDate.Format(time.DateOnly))
	if m.Width()-styles.Base.GetHorizontalFrameSize()-lipgloss.Width(releaseStr) <= 0 {
		releaseStr = ""
	}
	releaseStyle := styles.ReleaseDate.
		Align(lipgloss.Right).
		Width(m.Width() - styles.Base.GetHorizontalFrameSize() - lipgloss.Width(title))

	release := releaseStyle.Render(releaseStr)
	title = styles.Title.Render(title)
	desc := item.Description()
	desc = styles.Desc.Render(desc)

	stringBuilder.WriteString(lipgloss.JoinHorizontal(lipgloss.Bottom, title, release))
	stringBuilder.WriteRune('\n')
	stringBuilder.WriteString(desc)
	fmt.Fprint(w, styles.Base.Render(stringBuilder.String()))
}
