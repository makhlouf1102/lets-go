package dockerclient

import (
	"bufio"
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

type DockerContainer struct {
	ContainerID string
	IsRunning   bool
}

type ContainersMap struct {
	entities map[string]DockerContainer
}

func (cm *ContainersMap) AddContainer(programmingLanguage ProgrammingLanguage, container *DockerContainer) error {
	cm.entities[programmingLanguage.Name] = *container

	return nil
}

type ProgrammingLanguage struct {
	Name string
}

var publicContainers *ContainersMap = &ContainersMap{
	entities: make(map[string]DockerContainer),
}

func NewDockerContainer(id string, programmingLanguage ProgrammingLanguage) (*DockerContainer, error) {
	dc := &DockerContainer{
		ContainerID: id,
		IsRunning:   false,
	}

	publicContainers.AddContainer(programmingLanguage, dc)

	return dc, nil
}

func (dc *DockerContainer) Run(options container.StartOptions) error {
	ctx := context.Background()
	cli, err := GetDockerClient()

	if err != nil {
		fmt.Println("Problem while setting the client")
		return err
	}

	if err := cli.ContainerStart(ctx, dc.ContainerID, options); err != nil {
		fmt.Println("Problem while starting the client")
		return err
	}

	dc.IsRunning = false

	return nil
}

func (dc *DockerContainer) ExecuteCommand(execOptions container.ExecOptions, execAttachOptions container.ExecAttachOptions) ([]string, error) {
	ctx := context.Background()
	cli, err := GetDockerClient()

	if err != nil {
		fmt.Println("Problem while setting the client")
		return nil, err
	}

	// create an execution
	exec, err := cli.ContainerExecCreate(ctx, dc.ContainerID, execOptions)

	if err != nil {
		fmt.Println("Problem while creating an execution command")
		return nil, err
	}

	execID := exec.ID

	execStart, err := cli.ContainerExecAttach(ctx, execID, execAttachOptions)

	if err != nil {
		fmt.Println("Problem while executing an execution command")
		return nil, err
	}

	defer execStart.Close()

	scanner := bufio.NewScanner(execStart.Reader)

	var logs []string = make([]string, 0)

	for scanner.Scan() {
		logs = append(logs, scanner.Text())
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading exec output: %v", err)
		return nil, err
	}

	return logs, nil
}
