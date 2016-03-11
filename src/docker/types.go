package docker

import (
	"github.com/Sirupsen/logrus"
	dockerClient "github.com/fsouza/go-dockerclient"
)

var (
	log = logrus.WithFields(logrus.Fields{"package": "docker"})
)

// ContainerStore exposes methods to handle container lifecycle events.
// Instances of this interface should allow threadsafe reads and writes.
type ContainerStore interface {
	AddContainerByID(id string) error
	IAMRoles() []string
	IAMRoleForIP(ip string) (string, error)
	IAMRoleForID(ip string) (string, error)
	RemoveContainer(name string)
	SyncRunningContainers() error
}

// EventHandler instances implement DockerEventsChannel() which performs actions
// based on Docker events. Listen() is a blocking function which performs an
// action based on the events written to the channel.
type EventHandler interface {
	Listen(<-chan *dockerClient.APIEvents) error
}

// RawClient specifies the subset of commands that EventHandlers use from the
// go-dockerclient.
type RawClient interface {
	InspectContainer(id string) (*dockerClient.Container, error)
	ListContainers(opts dockerClient.ListContainersOptions) ([]dockerClient.APIContainers, error)
}