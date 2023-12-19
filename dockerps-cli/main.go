package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client" // v1.13.1
	"github.com/spf13/cobra"
)

func runPs(dockerCli *client.Client) {
	containers, err := dockerCli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("containers ID: ")
	for _, container := range containers {
		fmt.Printf("%s \n", container.ID[:10])

	}
}

func main() {
	dockerCli, _ := client.NewEnvClient()
	var cmdPs = &cobra.Command{
		Use:   "ps",
		Short: "list containers",
		Run: func(cmd *cobra.Command, args []string) {
			runPs(dockerCli)
		},
	}
	var rootCmd = &cobra.Command{Use: "play_docker"}
	rootCmd.AddCommand(cmdPs)
	rootCmd.Execute()
}
