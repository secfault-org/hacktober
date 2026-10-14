package styles

import "charm.land/lipgloss/v2"

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
		Timer   lipgloss.Style
	}

	Footer      lipgloss.Style
	HelpKey     lipgloss.Style
	HelpValue   lipgloss.Style
	HelpDivider lipgloss.Style
}

func DefaultStyles() *Styles {

	style := new(Styles)

	style.App = lipgloss.NewStyle().
		Margin(0, 0)

	style.NoContent = lipgloss.NewStyle().
		MarginTop(1).
		MarginLeft(2).
		Foreground(lipgloss.Color("242"))

	style.ChallengeListItemContainer = lipgloss.NewStyle().
		PaddingLeft(1).
		Border(lipgloss.Border{
			Left: " ",
		}, false, false, false, true).
		Height(2)

	style.ChallengeListItem.Normal.Base = lipgloss.NewStyle().
		PaddingLeft(1).
		Border(lipgloss.Border{Left: " "}, false, false, false, true).
		Height(2)

	style.ChallengeListItem.Normal.Title = lipgloss.NewStyle().Bold(true)

	style.ChallengeListItem.Normal.Desc = lipgloss.NewStyle().
		Foreground(lipgloss.Color("243"))

	style.ChallengeListItem.Normal.ReleaseDate = lipgloss.NewStyle().
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

	style.ChallengeDetail.Base = lipgloss.NewStyle()

	style.ChallengeDetail.Title = lipgloss.NewStyle().
		Padding(0, 2)

	style.ChallengeDetail.Body = lipgloss.NewStyle().
		Margin(1, 0)

	style.Statusbar.Base = lipgloss.NewStyle().
		Height(1)

	style.Statusbar.Spinner = lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("33"))

	style.Statusbar.Info = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("235")).
		Foreground(lipgloss.Color("243"))

	style.Statusbar.Help = lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color("243")).
		Background(lipgloss.Color("237"))

	style.Statusbar.Timer = lipgloss.NewStyle().
		Padding(0, 1).
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("33"))

	style.Footer = lipgloss.NewStyle().
		MarginTop(1).
		Height(1)

	style.HelpKey = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))

	style.HelpValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color("239"))

	style.HelpDivider = lipgloss.NewStyle().
		Foreground(lipgloss.Color("237")).
		SetString(" • ")

	return style
}
