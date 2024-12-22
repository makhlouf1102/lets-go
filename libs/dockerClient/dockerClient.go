package dockerclient

import (
	"sync"

	"github.com/docker/docker/client"
)

var dockerClient *client.Client
var once sync.Once

func GetDockerClient() (*client.Client, error) {
	var err error
	once.Do(func() {
		dockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	})
	return dockerClient, err
}
