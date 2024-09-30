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
