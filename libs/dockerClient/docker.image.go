package dockerclient

import (
	"context"
	"fmt"
	localTypes "lets-go/types"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
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

	creationResponse, err := cli.ContainerCreate(ctx, config, hostConfig, networkingConfig, nil, containerName)

	if err != nil {
		fmt.Println("problem while creating the container")
		return nil, err
	}

	dockerContainer, err := NewDockerContainer(creationResponse.ID, programmingLanguage)

	if err != nil {
		fmt.Println("problem while creating instanciating the container")
		return nil, err
	}

	return dockerContainer, nil
}
