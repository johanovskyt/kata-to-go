package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/johanovskyt/kata-to-go/internal/codewars"
	"github.com/johanovskyt/kata-to-go/internal/project"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:   "new",
		Usage:  "Create a new project",
		Action: newProject,
		Arguments: []cli.Argument{
			&cli.StringArg{Name: "id"},
			&cli.StringArg{Name: "path"},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func newProject(ctx context.Context, command *cli.Command) error {
	id := command.StringArg("id")
	if id == "" {
		return fmt.Errorf("id argument is required")
	}

	client := codewars.NewClient()

	kata, err := client.GetKata(ctx, id)
	if err != nil {
		return err
	}

	generator := project.NewGenerator()
	basePath := command.StringArg("path")
	if err := generator.Create(kata, basePath); err != nil {
		return err
	}

	fmt.Printf("Project created successfully at %s/%s\n", basePath, kata.Slug)

	return nil
}
