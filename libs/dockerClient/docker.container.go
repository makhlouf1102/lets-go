package dockerclient

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

type DockerContainer struct {
	ContainerID string
	IsRunning   bool
}

func NewDockerContainer(id string) (*DockerContainer, error) {
	return &DockerContainer{
		ContainerID: id,
		IsRunning:   false,
	}, nil
}

func (dc *DockerContainer) Run(options container.StartOptions) error {
	ctx := context.Background()
	cli, err := GetDockerClient()

	if err != nil {
		fmt.Println("Problem while setting the client")
		return err
	}

	if err := cli.ContainerStart(ctx, dc.ContainerID, options); err != nil {
		fmt.Println("Problem while starting the client")
		return err
	}

	dc.IsRunning = false

	return nil
}
