package dockerclient_test

import (
	dockerclient "lets-go/libs/dockerClient"
	"testing"
)

func TestBuildImage(t *testing.T) {

	dockerFile, err := dockerclient.NewDockerFile("../../docker/shared/common-base.Dockerfile")
	if err != nil {
		t.Fatalf("Failed to set the docker file struct : %v", err)
	}

	dockerImage, err := dockerFile.BuildImage()
	if err != nil {
		t.Fatalf("Failed while building the docker image : %v", err)
	}

	t.Fatal(dockerImage)
}
