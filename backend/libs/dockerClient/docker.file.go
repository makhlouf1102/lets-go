package dockerclient

import (
	"archive/tar"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
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
		Name: "DockerFile",
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

func (df *DockerFile) BuildImage(imageBuildOptions types.ImageBuildOptions) (*DockerImage, error) {
	ctx := context.Background()
	cli, err := GetDockerClient()

	if err != nil {
		fmt.Println("Problem while setting the client")
		return nil, err
	}

	builder, err := cli.ImageBuild(ctx, df.TarContent, imageBuildOptions)
	if err != nil {
		fmt.Println("Problem while creating the docker builder")
		return nil, err
	}

	defer builder.Body.Close()

	// Parse the build logs to extract the image ID
	var builtImageID string
	scanner := bufio.NewScanner(builder.Body)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line) // Log the build output for debugging
		var logEntry struct {
			Aux struct {
				ID string `json:"ID"`
			} `json:"aux"`
		}
		if err := json.Unmarshal([]byte(line), &logEntry); err == nil && logEntry.Aux.ID != "" {
			builtImageID = logEntry.Aux.ID
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading build output: %v\n", err)
		return nil, err
	}

	if builtImageID == "" {
		fmt.Println("Could not find built image ID in the build logs")
		return nil, fmt.Errorf("failed to extract image ID from build output")
	}

	imageInspect, _, err := cli.ImageInspectWithRaw(ctx, builtImageID)
	if err != nil {
		fmt.Printf("Error inspecting built image: %v\n", err)
		return nil, err
	}

	dockerImage, err := NewDockerImage(df, &imageInspect)

	if err != nil {
		fmt.Printf("Error while creating an instance of a docker image: %v\n", err)
		return nil, err
	}
	return dockerImage, nil
}
