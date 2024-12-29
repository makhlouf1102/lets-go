package dockerclient

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerContainer struct {
	ContainerID string
	IsRunning   bool
}

func NewDockerContainer(id string) (*DockerContainer, error) {
	dc := &DockerContainer{
		ContainerID: id,
		IsRunning:   false,
	}

	return dc, nil
}

func (dc *DockerContainer) Run(options container.StartOptions) error {

	if dc.ContainerID == "" {
		return errors.New("container ID is empty, cannot start the container")
	}

	ctx := context.Background()
	cli, err := GetDockerClient()

	if err != nil {
		fmt.Println("Problem while setting the client")
		return err
	}

	if err := cli.ContainerStart(ctx, dc.ContainerID, options); err != nil {
		log.Printf("Failed to start container '%s': %v", dc.ContainerID, err)
		return err
	}

	containerJSON, err := cli.ContainerInspect(ctx, dc.ContainerID)

	if err != nil {
		if client.IsErrNotFound(err) {
			fmt.Printf("Container '%s' does not exist.\n", dc.ContainerID)
			return err
		} else {
			return err
		}
	} else {
		if containerJSON.State.Running {
			fmt.Printf("The container '%s' is running.\n", dc.ContainerID)
			dc.IsRunning = true
		} else {
			fmt.Printf("The container '%s' is not running.\n", dc.ContainerID)
			return errors.New("The container is not running")
		}
	}

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
		log.Printf("Problem while creating an execution command : %v", err)
		return nil, err
	}

	execID := exec.ID

	if len(execID) == 0 {
		log.Printf("The execID is empty")
		return nil, errors.New("the exec id is empty")
	}

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
		log.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading exec output: %v", err)
		return nil, err
	}

	if len(logs) == 0 {
		log.Printf("no logs were returned")
		return nil, errors.New("no logs were returned")
	}

	return logs, nil
}
