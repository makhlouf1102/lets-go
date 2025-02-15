package dockerclient

import (
	"context"
	"fmt"
	localTypes "lets-go/types"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
)

type DockerImage struct {
	ImageID        string
	DockerFileRef  *DockerFile
	DockerImageRef *types.ImageInspect
}

func NewDockerImage(dockerFileRef *DockerFile, dockerImageRef *types.ImageInspect) (*DockerImage, error) {
	return &DockerImage{
		ImageID:        dockerImageRef.ID,
		DockerFileRef:  dockerFileRef,
		DockerImageRef: dockerImageRef,
	}, nil
}

func (di *DockerImage) CreateContainer(config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string, programmingLanguage localTypes.ProgrammingLanguage) (*DockerContainer, error) {
	ctx := context.Background()
	cli, err := GetDockerClient()

	if err != nil {
		fmt.Println("Problem while setting the client")
		return nil, err
	}

	config.Image = di.ImageID

	filter := filters.NewArgs(
		filters.Arg("name", containerName),
	)

	// check if the container exists
	containerList, err := cli.ContainerList(ctx, container.ListOptions{
		Filters: filter,
		All:     true,
	})

	if err != nil {
		return nil, err
	}

	var containerID string

	if len(containerList) > 0 {
		containerID = containerList[0].ID
	} else {
		creationResponse, err := cli.ContainerCreate(ctx, config, hostConfig, networkingConfig, nil, containerName)

		if err != nil {
			fmt.Println("problem while creating the container")
			return nil, err
		}

		containerID = creationResponse.ID

	}

	dockerContainer, err := NewDockerContainer(containerID)

	if err != nil {
		fmt.Println("problem while creating instanciating the container")
		return nil, err
	}

	return dockerContainer, nil
}
