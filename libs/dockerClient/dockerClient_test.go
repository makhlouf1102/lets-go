package dockerclient_test

import (
	dockerclient "lets-go/libs/dockerClient"
	"testing"
)

func TestBuildImage(t *testing.T) {

	dockerFile, err := dockerclient.NewDockerFile("../../docker/shared/Dockerfile")
	if err != nil {
		t.Fatalf("Failed to set the docker file struct : %v", err)
	}

	dockerImage, err := dockerFile.BuildImage("test")
	if err != nil {
		t.Fatalf("Failed while building the docker image : %v", err)
	}

	if len(dockerImage.DockerImageRef.Created) == 0 {
		t.Fatalf("The object is empty")
	}

	t.Logf("created image with ID : %s and the docker version is : %s", dockerImage.ImageID, dockerImage.DockerImageRef.DockerVersion)

}
