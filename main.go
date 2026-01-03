package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v3"
)

type kata struct {
	Name string `json:"name"`
}

const url = "https://www.codewars.com/api/v1"

func main() {
	cmd := &cli.Command{
		Name:   "new",
		Usage:  "Create a new project",
		Action: newProject,
		Arguments: []cli.Argument{
			&cli.StringArg{Name: "id"},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func newProject(ctx context.Context, command *cli.Command) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	reqUrl := fmt.Sprintf("%s/code-challenges/%s", url, command.StringArg("id"))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get response: %w", err)
	}

	if resp.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(resp.Body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	kata := kata{}
	err = json.Unmarshal(body, &kata)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	fmt.Println(kata.Name)

	return nil
}
