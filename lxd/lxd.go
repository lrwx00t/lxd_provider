package lxd

import (
	"fmt"

	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

const (
	DefaultLxdUnixSocket = "/var/snap/lxd/common/lxd/unix.socket"
	DefaultTimeout       = -1
)

var (
	c   lxd.InstanceServer
	err error

	stopRequest = api.InstanceStatePut{
		Action:  "stop",
		Timeout: DefaultTimeout,
	}

	startRequest = api.InstanceStatePut{
		Action:  "start",
		Timeout: DefaultTimeout,
	}
)

func init() {
	c, err = lxd.ConnectLXDUnix(DefaultLxdUnixSocket, nil)
	if err != nil {
		panic(err)
	}
}

func NewSocket(unitSocket string) {
	c, err = lxd.ConnectLXDUnix(unitSocket, nil)
	if err != nil {
		panic(err)
	}
}

func GetInstanceServer() lxd.InstanceServer {
	return c
}

func StopContainerWithName(containername string) error {
	for _, cc := range GetLXDContainers() {
		if cc.Name == containername {
			StopInstance(cc, containername)

			return nil
		}
	}
	err := fmt.Errorf("couldn't stop any container with the name %s. the provided name doesn't match any existing container", containername)

	return err
}

func GetContainerWithName(containername string) bool {
	for _, cc := range GetLXDContainers() {
		if cc.Name == containername {
			return true
		}
	}

	return false
}

func StartContainerWithName(containername string) error {
	for _, cc := range GetLXDContainers() {
		if cc.Name == containername {
			StartInstance(cc, containername)

			return nil
		}
	}
	err := fmt.Errorf("couldn't start any container with the name %s. the provided name doesn't match any existing container", containername)

	return err
}

func ManageContainerWithName(containername, status string) error {
	if status == "start" {
		if err := StartContainerWithName(containername); err != nil {
			return err
		}

		return nil
	} else if status == "stop" {
		if StopContainerWithName(containername) != nil {
			return err
		}

		return nil
	}
	err := fmt.Errorf("container status %s is not supported", status)

	return err
}

func StopInstance(cont api.Container, containername string) {
	if cont.Status == "Stopped" {
		return
	}
	op, err := c.UpdateInstanceState(containername, stopRequest, "")
	if err != nil {
		panic(err)
	}

	err = op.Wait()
	if err != nil {
		panic(err)
	}
}

func StartInstance(cont api.Container, containername string) {
	if cont.Status == "Running" {
		return
	}
	op, err := c.UpdateInstanceState(containername, startRequest, "")
	if err != nil {
		panic(err)
	}

	err = op.Wait()
	if err != nil {
		panic(err)
	}
}

func GetLXDContainers() []api.Container {
	cs, err := c.GetContainers()
	if err != nil {
		panic(err)
	}

	return cs
}
