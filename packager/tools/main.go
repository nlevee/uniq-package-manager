package tools

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
)

// ContainerOptions
type ContainerOptions struct {
	Image        string
	ImageVersion string
	Cmd          strslice.StrSlice
}

// GetFilePath
func GetFilePath(dir string, filename string) string {
	filepath := path.Join(dir, filename)
	fmt.Println("check if file exist and is readable : ", filepath)
	if info, err := os.Stat(filepath); err == nil && info.Mode().IsRegular() {
		return filepath
	}
	return ""
}

// CleanContainer
func CleanContainer(cli *client.Client, containerID string) {
	ctx := context.Background()
	fmt.Println("remove container", containerID)
	cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
		Force: true,
	})
}

// ShowContainerLog output all container log
func ShowContainerLog(cli *client.Client, containerID string) {
	ctx := context.Background()
	out, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{
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
}

// SetupEndHandler
func SetupEndHandler(cli *client.Client, containerID string) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		CleanContainer(cli, containerID)
		os.Exit(1)
	}()
}

// ImagePull
func ImagePull(cli *client.Client, nodeImage string) {
	ctx := context.Background()

	out, err := cli.ImagePull(ctx, nodeImage, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	if _, err := ioutil.ReadAll(out); err != nil {
		panic(err)
	}
}
