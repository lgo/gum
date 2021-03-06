package datastores

import (
	"errors"

	"github.com/xLegoz/gum/registry"
)

const defaultConfig = `
version: 9.5
`

func postgresStart(optionzs map[string]interface{}) error {
	// Compile step:
	// var version = "3.5"
	// folderPath, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(folderPath)
	// file := "app.py"
	// configCmds := []string{}
	// cmds := append([]string{
	// 	fmt.Sprintf("pip install -r requirements.txt && postgres3.5 %s", file),
	// }, configCmds...)
	// envs := []string{}
	//
	// var options = docker.CreateContainerOptions{
	// 	Name: "GUM_Python",
	// 	Config: &docker.Config{
	// 		WorkingDir: "/app",
	// 		Cmd:        cmds,
	// 		Env:        envs,
	// 		Entrypoint: []string{
	// 			"sh",
	// 			"-c",
	// 		},
	// 		Image: fmt.Sprintf("iron/postgres:%s-dev", version),
	// 		ExposedPorts: map[docker.Port]struct{}{
	// 			"5000/tcp": {},
	// 		},
	// 	},
	// 	HostConfig: &docker.HostConfig{
	// 		PortBindings: map[docker.Port][]docker.PortBinding{
	// 			"5000/tcp": []docker.PortBinding{docker.PortBinding{HostIP: "localhost", HostPort: "8000"}},
	// 		},
	// 		Binds: []string{
	// 			fmt.Sprintf("%s:/app", folderPath),
	// 		},
	// 	},
	// }
	// err = containers.Pull("iron/postgres", fmt.Sprintf("%s-dev", version))
	// if err != nil {
	// 	panic(err)
	// }
	// container, err := containers.Client.CreateContainer(options)
	// if err != nil {
	// 	panic(err)
	// }
	// err = containers.Client.StartContainer(container.ID, nil)
	//
	// if err != nil {
	// 	panic(err)
	// }
	// // _, err = containers.Client.WaitContainer(container.ID)
	//
	// if err != nil {
	// 	panic(err)
	// }
	// // err = containers.Client.RemoveContainer(docker.RemoveContainerOptions{
	// // 	ID:            container.ID,
	// // 	RemoveVolumes: true,
	// // 	Force:         true,
	// // })
	//
	// if err != nil {
	// 	panic(err)
	// }
	return nil
}

func postgresStop(options map[string]interface{}) error {
	return nil
}

func postgresPrepare(options map[string]interface{}) error {
	return nil
}

func postgresShouldRestart() {
	// TODO(xLegoz): logic for storing checksum of last binary against new binary to see if the application needs to be restarted
	// will be used for a filewatcher in development
}

func postgresVersions(options map[string]interface{}) error {
	if version, ok := options["version"]; ok && version != "3.5" {
		return errors.New("Non-supported Go version.")
	}

	return nil
}

func init() {
	registry.RegisterService(
		registry.Handler{
			Name:     "postgresql",
			Priority: 1,
			Versions: registry.WrapHandler(postgresVersions),
			Start:    registry.WrapHandler(postgresStart),
			Stop:     registry.WrapHandler(postgresStop),
			Prepare:  registry.WrapHandler(postgresPrepare),
			// Reload:   reload,
		},
	)
}
