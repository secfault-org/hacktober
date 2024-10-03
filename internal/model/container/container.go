package container

type Id = string
type Port = uint16
type State string

const (
	Starting State = "starting"
	Running        = "running"
	Stopping       = "stopping"
	Stopped        = "stopped"
)

type Container struct {
	ID       Id
	State    State
	HostPort Port
}

func (c State) ToEmoji() string {
	switch c {
	case Starting:
		return "🚀"
	case Running:
		return "🏃"
	case Stopping:
		return "🛑"
	case Stopped:
		return "💤"
	default:
		return "❓"
	}
}
