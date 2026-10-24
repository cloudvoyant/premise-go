"""Exercise real Premise generation for all templates in one root Go module."""

from pathlib import Path
import os
import shutil
import subprocess
import tempfile
import tomllib

ROOT = Path(__file__).resolve().parents[1]
PROJECTS = {
    "premise-cobra-cli": "apps/cli",
    "premise-connect-grpc-service": "apps/service",
    "premise-htmx-app": "apps/web",
    "premise-go-lib": "libs/math",
    "premise-pulumi-component-lib": "libs/components",
    "premise-pulumi-component-package": "libs/provider",
}


def main() -> None:
    requested = os.environ.get("PREMISE_BIN")
    if requested:
        resolved = shutil.which(requested)
        if resolved is None:
            raise RuntimeError(f"PREMISE_BIN does not resolve to an executable: {requested}")
    else:
        configuration = tomllib.loads((ROOT / "mise.toml").read_text())
        tool = "go:github.com/cloudvoyant/premise"
        version = configuration["tools"][tool]
        # Resolve the exact tool so another premise on PATH cannot shadow the pin.
        resolved = subprocess.check_output(
            ["mise", "which", "premise", "--tool", f"{tool}@{version}"],
            cwd=ROOT,
            text=True,
        ).strip()
    binary = str(Path(resolved).resolve())
    environment = dict(os.environ, GOWORK="off", PREMISE_TEMPLATE_TEST="1", MISE_YES="1")
    with tempfile.TemporaryDirectory(prefix="premise-go-workspace-") as temporary:
        workspace = Path(temporary)
        subprocess.run([binary, "init"], cwd=workspace, env=environment, check=True)
        for template, relative in PROJECTS.items():
            name = Path(relative).name
            answers = f"name={name},module=github.com/acme/example"
            if relative.startswith("libs/"):
                answers += ",package.public=true"
            print(f"Generate {relative} with the actual CLI", flush=True)
            subprocess.run(
                [binary, "generate", f"{ROOT}:{template}", "--answer", answers],
                cwd=workspace,
                env=environment,
                check=True,
            )
        subprocess.run(["go", "test", "./..."], cwd=workspace, env=environment, check=True)
        subprocess.run(["go", "vet", "./..."], cwd=workspace, env=environment, check=True)
        assert list(workspace.rglob("go.mod")) == [workspace / "go.mod"]
        assert not (workspace / "go.work").exists()
        assert not list(workspace.rglob(".premise-stage-*"))
        print("All six projects pass actual generation with one root module and no go.work")


if __name__ == "__main__":
    main()
