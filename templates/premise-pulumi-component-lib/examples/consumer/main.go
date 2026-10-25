package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	component "example.com/premise-module/libs/premise-pulumi-component-lib"
)

// Private helpers ------------------------------------------------------------
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		g, err := component.NewGreeting(ctx, "consumer", component.Args{Name: pulumi.String("consumer")})
		if err != nil {
			return err
		}
		ctx.Export("message", g.Message)
		return nil
	})
}
