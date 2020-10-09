package gonpm

import (
	"bufio"
	"context"
	"fmt"
	"os/user"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/nlevee/uniq-package-manager/packager/tools"
)

func newDefaultOptions() tools.ContainerOptions {
	return tools.ContainerOptions{
		Image:        "docker.io/library/node",
		ImageVersion: "lts",
	}
}

func createDockerWrapper(packagePath string, opts tools.ContainerOptions) {
	ctx := context.Background()
	user, _ := user.Current()

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// pull image
	dockerImage := opts.Image + ":" + opts.ImageVersion
	tools.ImagePull(cli, dockerImage)

	// creation du conteneur pour php-package
	container, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        dockerImage,
		User:         user.Uid + ":" + user.Gid,
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
		Entrypoint:   strslice.StrSlice{"/usr/local/bin/npm"},
		Cmd:          opts.Cmd,
		WorkingDir:   "/app",
	}, &container.HostConfig{
		RestartPolicy: container.RestartPolicy{Name: "no"},
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: packagePath,
				Target: "/app",
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
	out, err := cli.ContainerLogs(ctx, container.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		panic(err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("wait for container", container.ID)
	if _, err := cli.ContainerWait(ctx, container.ID); err != nil {
		panic(err)
	}
}
