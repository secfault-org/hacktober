package container

type Id = string
type Port = uint16

type Container struct {
	ID       Id
	HostPort Port
}
