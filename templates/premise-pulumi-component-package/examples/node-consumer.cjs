const assert = require("node:assert/strict");
const pulumi = require("../sdk/nodejs/node_modules/@pulumi/pulumi");
pulumi.runtime.setMocks({
  newResource(args) {
    assert.equal(args.type, "premise-pulumi-component-package:index:Greeting");
    return { id: args.name + "_id", state: { ...args.inputs, message: "Hello, " + args.inputs.name + "!" } };
  },
  call(args) { return args.inputs; }
}, "consumer", "test", false);
const { Greeting } = require("../sdk/nodejs/bin");
(async () => {
  const g = new Greeting("consumer", { name: "Node" });
  const message = await new Promise((resolve) => g.message.apply(resolve));
  assert.equal(message, "Hello, Node!");
  console.log("Generated Node SDK consumer passed with mocks");
})().catch((err) => { console.error(err); process.exitCode = 1; });
