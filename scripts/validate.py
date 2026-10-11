"""Validate isolated generated modules and their complete Mise contracts."""

from pathlib import Path
import os
import subprocess
import tempfile
import tomllib

from template_support import PUBLICATION_PLACEHOLDER, copy_template

ROOT = Path(__file__).resolve().parents[1]
COMMON_TASKS = [
    "install",
    "build",
    "test",
    "lint",
    "lint:fix",
    "format",
    "format:check",
    "env-pull",
    "publish:rc",
    "publish",
    "clean",
]
APPLICATION_TASKS = ["run", "dev", "deploy", "e2e"]


def publication_modes(source: Path) -> list[str]:
    """Exercise both policies when a library exposes the publication question."""
    policy = (source / ".premise-go.toml").read_text()
    if PUBLICATION_PLACEHOLDER in policy:
        return ["public", "private-local"]
    return ["private-local"]


def validate_project(source: Path, project: Path, publication: str, environment: dict) -> None:
    configuration = tomllib.loads((project / "mise.toml").read_text())
    policy = tomllib.loads((project / ".premise-go.toml").read_text())
    assert policy["publication"] == publication, (source.name, policy)
    task_names = COMMON_TASKS.copy()
    if "run" in configuration["tasks"]:
        task_names.extend(APPLICATION_TASKS)

    subprocess.run(
        ["mise", "trust", str(project / "mise.toml")],
        check=True,
        env=environment,
    )
    for task in task_names:
        assert task in configuration["tasks"], (source.name, task)
        print(f"{source.name} ({publication}): {task}", flush=True)
        subprocess.run(["mise", "run", task], cwd=project, env=environment, check=True)
    if "test:sdk" in configuration["tasks"]:
        subprocess.run(["mise", "run", "test:sdk"], cwd=project, env=environment, check=True)


def main() -> None:
    # Publication tasks validate locally and never upload packages in this mode.
    environment = dict(os.environ, GOWORK="off", PREMISE_TEMPLATE_TEST="1", MISE_YES="1")
    for source in sorted((ROOT / "templates").iterdir()):
        for publication in publication_modes(source):
            with tempfile.TemporaryDirectory(prefix="premise-go-") as temporary:
                project = Path(temporary) / "project"
                module = "github.com/acme/example/" + source.name.removeprefix("premise-")
                copy_template(source, project, module, publication)
                validate_project(source, project, publication, environment)


if __name__ == "__main__":
    main()
