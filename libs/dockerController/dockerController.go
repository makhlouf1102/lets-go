package dockerController

import (
	"errors"
	dockerclient "lets-go/libs/dockerClient"
	localTypes "lets-go/types"
	"log"
	"sync"
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
