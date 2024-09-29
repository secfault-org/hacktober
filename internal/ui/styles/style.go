package styles

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	NoContent lipgloss.Style

	App                        lipgloss.Style
	ChallengeListItemContainer lipgloss.Style

	ChallengeListItem struct {
		Normal struct {
			Base        lipgloss.Style
			Title       lipgloss.Style
			Desc        lipgloss.Style
			ReleaseDate lipgloss.Style
		}
		Active struct {
			Base        lipgloss.Style
			Title       lipgloss.Style
			Desc        lipgloss.Style
			ReleaseDate lipgloss.Style
		}
	}

	ChallengeDetail struct {
		Base  lipgloss.Style
		Title lipgloss.Style
		Body  lipgloss.Style
	}

	Statusbar struct {
		Base    lipgloss.Style
		Help    lipgloss.Style
		Info    lipgloss.Style
		Spinner lipgloss.Style
	}

	Footer      lipgloss.Style
	HelpKey     lipgloss.Style
	HelpValue   lipgloss.Style
	HelpDivider lipgloss.Style
}

func DefaultStyles(renderer *lipgloss.Renderer) *Styles {

	style := new(Styles)

	style.App = renderer.NewStyle().
		Margin(0, 0)

	style.NoContent = renderer.NewStyle().
		MarginTop(1).
		MarginLeft(2).
		Foreground(lipgloss.Color("242"))

	style.ChallengeListItemContainer = renderer.NewStyle().
		PaddingLeft(1).
		Border(lipgloss.Border{
			Left: " ",
		}, false, false, false, true).
		Height(2)

	style.ChallengeListItem.Normal.Base = renderer.NewStyle().
		PaddingLeft(1).
		Border(lipgloss.Border{Left: " "}, false, false, false, true).
		Height(2)

	style.ChallengeListItem.Normal.Title = renderer.NewStyle().Bold(true)

	style.ChallengeListItem.Normal.Desc = renderer.NewStyle().
		Foreground(lipgloss.Color("243"))

	style.ChallengeListItem.Normal.ReleaseDate = renderer.NewStyle().
		Foreground(lipgloss.Color("243"))

	style.ChallengeListItem.Active.Base = style.ChallengeListItem.Normal.Base.
		BorderStyle(lipgloss.Border{Left: "┃"}).
		BorderForeground(lipgloss.Color("176"))

	style.ChallengeListItem.Active.Title = style.ChallengeListItem.Normal.Title.
		Foreground(lipgloss.Color("212"))

	style.ChallengeListItem.Active.Desc = style.ChallengeListItem.Normal.Desc.
		Foreground(lipgloss.Color("246"))

	style.ChallengeListItem.Active.ReleaseDate = style.ChallengeListItem.Normal.ReleaseDate.
		Foreground(lipgloss.Color("212"))

	style.ChallengeDetail.Base = renderer.NewStyle()

	style.ChallengeDetail.Title = renderer.NewStyle().
		Padding(0, 2)

	style.ChallengeDetail.Body = renderer.NewStyle().
		Margin(1, 0)

	style.Statusbar.Base = renderer.NewStyle().
		Height(1)

	style.Statusbar.Spinner = renderer.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("33"))

	style.Statusbar.Info = renderer.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("235")).
		Foreground(lipgloss.Color("243"))

	style.Statusbar.Help = renderer.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color("243")).
		Background(lipgloss.Color("237"))

	style.Footer = renderer.NewStyle().
		MarginTop(1).
		Padding(0, 1).
		Height(1)

	style.HelpKey = renderer.NewStyle().
		Foreground(lipgloss.Color("241"))

	style.HelpValue = renderer.NewStyle().
		Foreground(lipgloss.Color("239"))

	style.HelpDivider = renderer.NewStyle().
		Foreground(lipgloss.Color("237")).
		SetString(" • ")

	return style
}
