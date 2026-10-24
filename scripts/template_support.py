"""Prepare renamed template copies without changing registry source files."""

from pathlib import Path
import shutil

MODULE_PLACEHOLDER = "example.com/premise-module"


def copy_template(source: Path, destination: Path, module: str) -> None:
    """Copy a source template and apply the questionnaire substitutions."""
    shutil.copytree(
        source,
        destination,
        ignore=shutil.ignore_patterns("dist", ".tools", "node_modules", "bin"),
    )
    for file in destination.rglob("*"):
        if not file.is_file():
            continue
        try:
            text = file.read_text()
        except UnicodeDecodeError:
            continue
        text = text.replace(MODULE_PLACEHOLDER, module)
        text = text.replace(source.name, destination.name)
        file.write_text(text)
