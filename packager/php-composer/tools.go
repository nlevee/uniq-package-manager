package gocomposer

import (
	"context"
	"fmt"
	"os/user"
	"path"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/nlevee/uniq-package-manager/packager/tools"
)

func newDefaultOptions() tools.ContainerOptions {
	return tools.ContainerOptions{
		Image:        "docker.io/library/composer",
		ImageVersion: "latest",
	}
}

func createDockerWrapper(composerPath string, opts tools.ContainerOptions) {
	ctx := context.Background()
	user, _ := user.Current()

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	dockerImage := opts.Image + ":" + opts.ImageVersion
	tools.ImagePull(cli, dockerImage)

	// creation du conteneur pour php-composer
	container, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        dockerImage,
		User:         user.Uid + ":" + user.Gid,
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
		Entrypoint:   strslice.StrSlice{"/usr/bin/composer", "-n"},
		Cmd:          opts.Cmd,
	}, &container.HostConfig{
		RestartPolicy: container.RestartPolicy{Name: "no"},
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: composerPath,
				Target: "/app",
			},
			{
				Type:   mount.TypeBind,
				Source: path.Join(home, ".cache/composer"),
				Target: "/opt/composer/cache",
			},
		},
	}, nil, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("create container", container.ID)

	// check du signal pour suppression du container en SIGTERM
	tools.SetupEndHandler(cli, container.ID)

	defer tools.CleanContainer(cli, container.ID)

	// Start created container
	fmt.Println("start container", container.ID)
	if err := cli.ContainerStart(ctx, container.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	// Show container logs
	tools.ShowContainerLog(cli, container.ID)

	// wait for container end
	fmt.Println("wait for container", container.ID)
	if _, err := cli.ContainerWait(ctx, container.ID); err != nil {
		panic(err)
	}
}
