package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/johanovskyt/kata-to-go/internal/codewars"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Create(kata *codewars.Kata, basePath string) error {
	projectPath := filepath.Join(basePath, kata.Slug)
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := g.initModule(kata.Slug, projectPath); err != nil {
		return err
	}

	if err := g.createReadme(kata, projectPath); err != nil {
		return err
	}

	if err := g.createMainFile(projectPath); err != nil {
		return err
	}

	return nil
}

func (g *Generator) initModule(moduleName, projectPath string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize go module: %w", err)
	}
	return nil
}

func (g *Generator) createReadme(kata *codewars.Kata, projectPath string) error {
	readme := fmt.Sprintf("# %s\n\n%s\n\n[View on Codewars](%s)\n", kata.Name, kata.Description, kata.URL)
	readmePath := filepath.Join(projectPath, "README.md")
	if err := os.WriteFile(readmePath, []byte(readme), 0644); err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	return nil
}

func (g *Generator) createMainFile(projectPath string) error {
	mainContent := `package main

func main() {
	// happy coding!
}
`
	mainPath := filepath.Join(projectPath, "main.go")
	if err := os.WriteFile(mainPath, []byte(mainContent), 0644); err != nil {
		return fmt.Errorf("failed to create main.go: %w", err)
	}
	return nil
}
