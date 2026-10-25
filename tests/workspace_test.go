package tests

import (
	"bytes"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// Tests ----------------------------------------------------------------------

func TestScaffold(t *testing.T) {
	t.Run("when all Go templates join one workspace", func(t *testing.T) {
		// Arrange
		root, registry, binary := workspaceFixture(t)
		projects := []struct {
			template string
			path     string
		}{
			{
				template: "premise-cobra-cli",
				path:     "apps/cli",
			},
			{
				template: "premise-connect-grpc-service",
				path:     "apps/service",
			},
			{
				template: "premise-htmx-app",
				path:     "apps/web",
			},
			{
				template: "premise-go-lib",
				path:     "libs/math",
			},
			{
				template: "premise-pulumi-component-lib",
				path:     "libs/components",
			},
			{
				template: "premise-pulumi-component-package",
				path:     "libs/provider",
			},
		}
		var clientRules []byte

		// Act
		for index, project := range projects {
			t.Logf("Generate %s", project.path)
			answers := "name=" + filepath.Base(project.path) + ",module=github.com/zexample/workspace"
			if strings.HasPrefix(project.path, "libs/") {
				answers += ",package.public=true"
			}
			runCommand(t, root, binary, "generate", registry+":"+project.template, "--answer", answers)
			if index == 0 {
				baseline, err := os.ReadFile(filepath.Join(root, ".golangci.yml"))
				if err != nil {
					t.Fatal(err)
				}
				clientRules = append([]byte("# Client quality rules\n"), bytes.ReplaceAll(baseline, []byte("max-complexity: 30"), []byte("max-complexity: 25"))...)
				if err := os.WriteFile(filepath.Join(root, ".golangci.yml"), clientRules, 0o644); err != nil {
					t.Fatal(err)
				}
			}
		}
		runCommand(t, root, "go", "test", "./...")
		runCommand(t, root, "go", "vet", "./...")

		// Assert
		t.Run("it shares one module and preserves client quality rules", func(t *testing.T) {
			var modules []string
			err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if entry.Name() == "go.mod" {
					modules = append(modules, path)
				}
				if entry.Name() == "go.work" || strings.HasPrefix(entry.Name(), ".premise-stage-") {
					t.Errorf("unexpected generated path: %s", path)
				}
				return nil
			})
			if err != nil || len(modules) != 1 || modules[0] != filepath.Join(root, "go.mod") {
				t.Fatalf("module paths = %v, error = %v", modules, err)
			}
			actual, err := os.ReadFile(filepath.Join(root, ".golangci.yml"))
			if err != nil || !bytes.Equal(actual, clientRules) {
				t.Fatalf("client rules changed: %s, %v", actual, err)
			}
			if _, err := os.Stat(filepath.Join(root, ".editorconfig")); err != nil {
				t.Fatal(err)
			}
			for _, project := range projects {
				if _, err := os.Stat(filepath.Join(root, project.path, ".golangci.yml")); !os.IsNotExist(err) {
					t.Fatalf("project owns separate formatting rules: %s, %v", project.path, err)
				}
			}
		})
	})

	t.Run("when generated code violates workspace quality rules", func(t *testing.T) {
		// Arrange
		root, registry, binary := workspaceFixture(t)
		runCommand(t, root, binary, "generate", registry+":premise-go-lib", "--answer", "name=example,module=github.com/zexample/workspace,package.public=false")
		project := filepath.Join(root, "libs", "example")
		parsed, err := parser.ParseFile(token.NewFileSet(), filepath.Join(project, "greeting.go"), nil, parser.PackageClauseOnly)
		if err != nil {
			t.Fatal(err)
		}
		file := filepath.Join(project, "quality.go")
		if err := os.WriteFile(file, []byte("package "+parsed.Name.Name+"\nfunc Unformatted( )int{return 1}\n"), 0o644); err != nil {
			t.Fatal(err)
		}

		// Act
		output, formattingError := executeCommand(project, "mise", "run", "format:check")

		// Assert
		t.Run("it rejects unformatted code and repairs it through the format task", func(t *testing.T) {
			if formattingError == nil {
				t.Fatalf("unformatted code passed: %s", output)
			}
			runCommand(t, project, "mise", "run", "format")
			runCommand(t, project, "mise", "run", "format:check")
		})

		// Arrange
		code := "package " + parsed.Name.Name + "\nfunc Complex(value int) int {\n" + strings.Repeat("if value > 0 { value-- }\n", 31) + "return value\n}\n"
		if err := os.WriteFile(file, []byte(code), 0o644); err != nil {
			t.Fatal(err)
		}
		runCommand(t, project, "mise", "run", "format")

		// Act
		output, complexityError := executeCommand(project, "mise", "run", "lint")

		// Assert
		t.Run("it rejects functions over the shared complexity limit", func(t *testing.T) {
			if complexityError == nil || !strings.Contains(string(output), "cyclop") {
				t.Fatalf("expected complexity failure, got %v:\n%s", complexityError, output)
			}
		})
	})
}

func TestScaffoldWorkflows(t *testing.T) {
	t.Run("when a Go library creates client workflows", func(t *testing.T) {
		// Arrange
		root, registry, binary := workspaceFixture(t)

		// Act
		runCommand(t, root, binary, "generate", registry+":premise-go-lib", "--answer", "name=example,module=github.com/zexample/workspace,package.public=false")

		// Assert
		t.Run("it uses the same released action as the registry", func(t *testing.T) {
			actionPattern := regexp.MustCompile(`cloudvoyant/premise@v[^\s]+`)
			for _, workflow := range []string{"on-commit.yml", "on-merge.yml"} {
				registryWorkflow, err := os.ReadFile(filepath.Join(registry, ".github", "workflows", workflow))
				if err != nil {
					t.Fatal(err)
				}
				expected := actionPattern.Find(registryWorkflow)
				clientWorkflow, err := os.ReadFile(filepath.Join(root, ".github", "workflows", workflow))
				if err != nil {
					t.Fatal(err)
				}
				actions := actionPattern.FindAll(clientWorkflow, -1)
				if len(expected) == 0 || len(actions) == 0 {
					t.Fatalf("missing released action in %s", workflow)
				}
				for _, action := range actions {
					if !bytes.Equal(action, expected) {
						t.Fatalf("%s uses %s, registry uses %s", workflow, action, expected)
					}
				}
			}
		})
	})
}

// Private helpers ------------------------------------------------------------

func workspaceFixture(t *testing.T) (string, string, string) {
	t.Helper()
	registry, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}
	binary := os.Getenv("PREMISE_BIN")
	if binary == "" {
		binary, err = exec.LookPath("premise")
		if err != nil {
			t.Fatal(err)
		}
	}
	root := t.TempDir()
	runCommand(t, root, binary, "init")
	return root, registry, binary
}

func executeCommand(directory, binary string, arguments ...string) ([]byte, error) {
	command := exec.Command(binary, arguments...)
	command.Dir = directory
	command.Env = append(os.Environ(), "GOWORK=off", "PREMISE_TEMPLATE_TEST=1", "MISE_YES=1")
	return command.CombinedOutput()
}

func runCommand(t *testing.T, directory, binary string, arguments ...string) {
	t.Helper()
	output, err := executeCommand(directory, binary, arguments...)
	if err != nil {
		t.Fatalf("%s %s: %v\n%s", binary, strings.Join(arguments, " "), err, output)
	}
}
