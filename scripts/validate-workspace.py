"""Catch dependency selection failures across generated Go project modules."""
from pathlib import Path
import os, shutil, subprocess, tempfile
ROOT = Path(__file__).resolve().parents[1]
with tempfile.TemporaryDirectory(prefix="premise-go-workspace-") as temporary:
    workspace = Path(temporary)
    directories = []
    for source in sorted((ROOT / "templates").iterdir()):
        destination = workspace / "projects" / source.name.removeprefix("premise-")
        shutil.copytree(source, destination, ignore=shutil.ignore_patterns("dist", ".tools", "node_modules", "bin"))
        module = "github.com/acme/example/projects/" + destination.name
        for file in destination.rglob("*"):
            if file.is_file():
                file.write_text(file.read_text().replace("example.com/premise-module", module).replace("premise-publication", "public"))
        directories.append(destination)
    env = {key: value for key, value in os.environ.items() if key != "GOWORK"}
    subprocess.run(["go", "work", "init", *[str(path.relative_to(workspace)) for path in directories]], cwd=workspace, env=env, check=True)
    for directory in directories:
        print(f"Shared go.work: {directory.name}", flush=True)
        subprocess.run(["go", "test", "./..."], cwd=directory, env=env, check=True)
        subprocess.run(["go", "vet", "./..."], cwd=directory, env=env, check=True)
    print("All six renamed Go modules pass with shared workspace dependency selection")
