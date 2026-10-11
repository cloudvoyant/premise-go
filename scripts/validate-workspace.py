"""Catch dependency selection failures across generated Go project modules."""

from pathlib import Path
import os
import subprocess
import tempfile

from template_support import copy_template

ROOT = Path(__file__).resolve().parents[1]


def prepare_workspace(workspace: Path) -> list[Path]:
    directories = []
    for source in sorted((ROOT / "templates").iterdir()):
        destination = workspace / "projects" / source.name.removeprefix("premise-")
        module = "github.com/acme/example/projects/" + destination.name
        copy_template(source, destination, module, "public")
        directories.append(destination)
    return directories


def validate_workspace(workspace: Path, directories: list[Path]) -> None:
    # Use the temporary go.work rather than any workspace inherited by the caller.
    environment = {key: value for key, value in os.environ.items() if key != "GOWORK"}
    module_paths = [str(directory.relative_to(workspace)) for directory in directories]
    subprocess.run(
        ["go", "work", "init", *module_paths],
        cwd=workspace,
        env=environment,
        check=True,
    )
    for directory in directories:
        print(f"Shared go.work: {directory.name}", flush=True)
        subprocess.run(["go", "test", "./..."], cwd=directory, env=environment, check=True)
        subprocess.run(["go", "vet", "./..."], cwd=directory, env=environment, check=True)
    print("All six renamed Go modules pass with shared workspace dependency selection")


def main() -> None:
    with tempfile.TemporaryDirectory(prefix="premise-go-workspace-") as temporary:
        workspace = Path(temporary)
        directories = prepare_workspace(workspace)
        validate_workspace(workspace, directories)


if __name__ == "__main__":
    main()
