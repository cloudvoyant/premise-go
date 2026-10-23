"""Validate all generated projects under one Go module and dependency set."""

from pathlib import Path
import json
import os
import subprocess
import tempfile

from template_support import copy_template

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
    environment = dict(os.environ, GOWORK="off")
    with tempfile.TemporaryDirectory(prefix="premise-go-workspace-") as temporary:
        workspace = Path(temporary)
        requirements = {}
        checksums = set()
        for template, relative in PROJECTS.items():
            project = workspace / relative
            copy_template(ROOT / "templates" / template, project, "github.com/acme/example", "public")
            data = subprocess.check_output(["go", "mod", "edit", "-json"], cwd=project, env=environment, text=True)
            manifest = json.loads(data)
            for requirement in manifest.get("Require") or []:
                requirements[requirement["Path"]] = requirement["Version"]
            sums = project / "go.sum"
            if sums.exists():
                checksums.update(sums.read_text().splitlines())
                sums.unlink()
            (project / "go.mod").unlink()
        (workspace / "go.mod").write_text("module github.com/acme/example\n\ngo 1.25.0\n")
        for path, version in requirements.items():
            subprocess.run(["go", "mod", "edit", f"-require={path}@{version}"], cwd=workspace, env=environment, check=True)
        (workspace / "go.sum").write_text("\n".join(sorted(checksums)) + "\n")
        subprocess.run(["go", "mod", "tidy"], cwd=workspace, env=environment, check=True)
        for relative in PROJECTS.values():
            print(f"Shared root module: {relative}", flush=True)
            subprocess.run(["go", "test", "./" + relative + "/..."], cwd=workspace, env=environment, check=True)
            subprocess.run(["go", "vet", "./" + relative + "/..."], cwd=workspace, env=environment, check=True)
        assert list(workspace.rglob("go.mod")) == [workspace / "go.mod"]
        assert not (workspace / "go.work").exists()
        print("All six projects pass with one root module and no go.work")


if __name__ == "__main__":
    main()
