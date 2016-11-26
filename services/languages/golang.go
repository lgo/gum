package languages

import (
	"errors"
	"fmt"
	"os"

	"github.com/fsouza/go-dockerclient"
	"github.com/xLegoz/gum/containers"
	"github.com/xLegoz/gum/registry"
	// "gopkg.in/yaml.v2"
)

const DefaultConfig = `
version: 1.7
`

func golangStart(options map[string]interface{}) error {
	return nil
}

func golangStop(options map[string]interface{}) error {
	return nil
}

func golangPrepare(optionzs map[string]interface{}) error {
	// Compile step:
	var version = "1.7"
	folderPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(folderPath)
	var options = docker.CreateContainerOptions{
		Name: "GolangBuilder",
		Config: &docker.Config{
			WorkingDir: "/app",
			Cmd:        []string{},
			Env: []string{
				"CGO_ENABLE=0",
			},
			Entrypoint: []string{
				"go",
				"build",
				"-a",
				"--installsuffix",
				"cgo",
				"--ldflags='-s'",
				"-o",
				"hello",
			},
			Image: fmt.Sprintf("golang:%s", version),
		},
		HostConfig: &docker.HostConfig{
			Binds: []string{
				fmt.Sprintf("%s:/go", os.Getenv("GOPATH")),
				fmt.Sprintf("%s:/app", folderPath),
			},
		},
	}
	err = containers.Pull("golang", version)
	if err != nil {
		panic(err)
	}
	container, err := containers.Client.CreateContainer(options)
	if err != nil {
		panic(err)
	}
	err = containers.Client.StartContainer(container.ID, nil)

	if err != nil {
		panic(err)
	}
	_, err = containers.Client.WaitContainer(container.ID)

	if err != nil {
		panic(err)
	}
	// err = containers.Client.RemoveContainer(docker.RemoveContainerOptions{
	// 	ID:            container.ID,
	// 	RemoveVolumes: true,
	// 	Force:         true,
	// })

	if err != nil {
		panic(err)
	}
	return err
}

func shouldRestart() {
	// TODO(xLegoz): logic for storing checksum of last binary against new binary to see if the application needs to be restarted
	// will be used for a filewatcher in development
}

func golangVersions(options map[string]interface{}) error {
	if version, ok := options["version"]; ok && version != "1.7" {
		return errors.New("Non-supported Go version.")
	}

	return nil
}

func init() {
	registry.RegisterLanguage(
		registry.Handler{
			Name:     "go",
			Priority: 1,
			Versions: registry.WrapHandler(golangVersions),
			Start:    registry.WrapHandler(golangStart),
			Stop:     registry.WrapHandler(golangStop),
			Prepare:  registry.WrapHandler(golangPrepare),
		},
	)
}
