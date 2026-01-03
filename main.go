package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:   "new",
		Usage:  "Create a new project",
		Action: newProject,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func newProject(ctx context.Context, command *cli.Command) error {
	fmt.Println("TODO")

	return nil
}
