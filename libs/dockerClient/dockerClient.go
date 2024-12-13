package dockerclient

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerFile struct {
	path       string
	TarContent io.Reader
}

func NewDockerFile(path string) (*DockerFile, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	defer tw.Close()

	header := &tar.Header{
		Name: "common-base.DockerFile",
		Mode: 0600,
		Size: int64(len(fileContent)),
	}

	if err := tw.WriteHeader(header); err != nil {
		return nil, err
	}

	if _, err := tw.Write(fileContent); err != nil {
		return nil, err
	}

	return &DockerFile{
		path:       path,
		TarContent: bytes.NewReader(buf.Bytes()),
	}, nil
}

func (df *DockerFile) BuildImage() (*DockerImage, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		fmt.Println("Problem while setting the client")
		return nil, err
	}

	defer cli.Close()

	builder, err := cli.ImageBuild(ctx, df.TarContent, types.ImageBuildOptions{
		Dockerfile:     "common-base.DockerFile", // Specify the name here
	})
	if err != nil {
		fmt.Println("Problem while creating the docker builder")
		return nil, err
	}

	defer builder.Body.Close()

	scanner := bufio.NewScanner(builder.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading build output: %v\n", err)
		return nil, err
	}

	// Return a placeholder DockerImage for now
	return &DockerImage{}, nil

	// var builtImage types.ImageInspect
	// if err := json.NewDecoder(builder.Body).Decode(&builtImage); err != nil {
	// 	fmt.Println("Problem while decoding the body")
	// 	return nil, err
	// }

	// return NewDockerImage(df, &builtImage)
}

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
