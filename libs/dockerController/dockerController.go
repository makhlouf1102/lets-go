package dockerController

import (
	"errors"
	dockerclient "lets-go/libs/dockerClient"
	localTypes "lets-go/types"
	"log"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

type ContainersMap struct {
	entities map[string]dockerclient.DockerContainer
}

func (cm *ContainersMap) AddContainer(programmingLanguage localTypes.ProgrammingLanguage, container *dockerclient.DockerContainer) error {
	cm.entities[programmingLanguage.Name] = *container

	return nil
}

func (cm *ContainersMap) Get(programmingLanguage localTypes.ProgrammingLanguage) (*dockerclient.DockerContainer, error) {
	tagetedContainer, ok := cm.entities[programmingLanguage.Name]
	if !ok {
		log.Printf("the container for : %s does not exists", programmingLanguage.Name)
		return nil, errors.New("the container does not exists")
	}

	return &tagetedContainer, nil

}

var publicContainers *ContainersMap
var once sync.Once

func GetContainersMap() (*ContainersMap, error) {
	once.Do(func() {
		publicContainers = &ContainersMap{
			entities: make(map[string]dockerclient.DockerContainer),
		}
	})

	return publicContainers, nil
}

func InitContainers() error {
	dockerFile, err := dockerclient.NewDockerFile("./docker/shared/DockerFile")

	if err != nil {
		return err
	}

	dockerImage, err := dockerFile.BuildImage(types.ImageBuildOptions{
		Tags:       []string{"js-letsgo-image"},
		Dockerfile: "DockerFile",
	})

	if err != nil {
		return err
	}

	hostConfig := &container.HostConfig{}

	// Empty `network.NetworkingConfig`
	networkingConfig := &network.NetworkingConfig{}

	containerName := "js-letsgo-container"

	config := &container.Config{}

	programmingLanguage := localTypes.ProgrammingLanguage{Name: "javascript"}

	dockerContainer, err := dockerImage.CreateContainer(config, hostConfig, networkingConfig, containerName, programmingLanguage)

	if err != nil {
		return err
	}

	pc, err := GetContainersMap()

	if err != nil {
		return err
	}

	if err := pc.AddContainer(programmingLanguage, dockerContainer); err != nil {
		return err
	}

	if err := dockerContainer.Run(container.StartOptions{}); err != nil {
		return err
	}

	return nil

}
