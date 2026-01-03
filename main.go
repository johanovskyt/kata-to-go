package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/johanovskyt/kata-to-go/internal/codewars"
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
	client := codewars.NewClient()

	kata, err := client.GetKata(ctx, command.StringArg("id"))
	if err != nil {
		return err
	}

	fmt.Println(kata)

	return nil
}
