package main

import (
	"context"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"log"
)

// version is replaced by the coordinated release linker flag.
var version = "0.1.0"

type GreetingArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}
type Greeting struct {
	pulumi.ResourceState
	Message pulumi.StringOutput `pulumi:"message"`
}

// Public API -----------------------------------------------------------------

// Annotate fixes the component token independently of the Go module path.
func (*Greeting) Annotate(a infer.Annotator) { a.SetToken("index", "Greeting") }

// NewGreeting implements a cross-language component.
func NewGreeting(ctx *pulumi.Context, name string, args GreetingArgs, opts ...pulumi.ResourceOption) (*Greeting, error) {
	c := &Greeting{}
	if err := ctx.RegisterComponentResource(p.GetTypeToken(ctx.Context()), name, c, opts...); err != nil {
		return nil, err
	}
	c.Message = args.Name.ToStringOutput().ApplyT(func(n string) string { return "Hello, " + n + "!" }).(pulumi.StringOutput)
	return c, ctx.RegisterResourceOutputs(c, pulumi.Map{"message": c.Message})
}

// Private helpers ------------------------------------------------------------
func provider() (p.Provider, error) {
	return infer.NewProviderBuilder().WithNamespace("premise").WithDisplayName("greeting").WithGoImportPath("example.com/premise-module/libs/premise-pulumi-component-package/sdk/go/greeting").WithComponents(infer.ComponentF(NewGreeting)).Build()
}
func main() {
	v, err := provider()
	if err != nil {
		log.Fatal(err)
	}
	if err = v.Run(context.Background(), "premise-pulumi-component-package", version); err != nil {
		log.Fatal(err)
	}
}
