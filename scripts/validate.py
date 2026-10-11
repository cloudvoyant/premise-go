"""Validate isolated generated modules and their complete Mise contracts."""
from pathlib import Path
import os, shutil, subprocess, tempfile, tomllib
ROOT = Path(__file__).resolve().parents[1]
COMMON = ["install", "build", "test", "lint", "lint:fix", "format", "format:check", "env-pull", "publish:rc", "publish", "clean"]
env = dict(os.environ, GOWORK="off", PREMISE_TEMPLATE_TEST="1", MISE_YES="1")
for source in sorted((ROOT / "templates").iterdir()):
    for publication in (["public", "private-local"] if "premise-publication" in (source / ".premise-go.toml").read_text() else ["private-local"]):
        with tempfile.TemporaryDirectory(prefix="premise-go-") as tmp:
            project = Path(tmp) / "project"
            shutil.copytree(source, project, ignore=shutil.ignore_patterns("dist", ".tools", "node_modules", "bin"))
            for file in project.rglob("*"):
                if file.is_file():
                    try:
                        text = file.read_text()
                    except UnicodeDecodeError:
                        continue
                    module = "github.com/acme/example/" + source.name.removeprefix("premise-")
                    file.write_text(text.replace("premise-publication", publication).replace("example.com/premise-module", module))
            config = tomllib.loads((project / "mise.toml").read_text())
            metadata = tomllib.loads((project / ".premise-go.toml").read_text())
            tasks = COMMON + (["run", "dev", "deploy", "e2e"] if "run" in config["tasks"] else [])
            subprocess.run(["mise", "trust", str(project / "mise.toml")], check=True, env=env)
            for task in tasks:
                assert task in config["tasks"], (source.name, task)
                print(f"{source.name} ({publication}): {task}", flush=True)
                subprocess.run(["mise", "run", task], cwd=project, env=env, check=True)
            if "test:sdk" in config["tasks"]:
                subprocess.run(["mise", "run", "test:sdk"], cwd=project, env=env, check=True)
