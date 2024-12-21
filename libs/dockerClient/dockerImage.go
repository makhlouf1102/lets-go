package dockerclient

import "github.com/docker/docker/api/types"

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

