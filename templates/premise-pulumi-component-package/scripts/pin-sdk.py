from pathlib import Path
import json
p = Path("sdk/nodejs/package.json")
v = json.loads(p.read_text())
v["dependencies"]["@pulumi/pulumi"] = "3.200.0"
v["devDependencies"] = {"@types/node": "22.10.2", "typescript": "5.7.3"}
v["scripts"]["build"] = "tsc && cp package.json bin/package.json"
p.write_text(json.dumps(v, indent=2) + "\n")
