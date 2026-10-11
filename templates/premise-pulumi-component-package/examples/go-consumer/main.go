package main

import (
	greeting "example.com/premise-module/sdk/go/greeting"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Private helpers ------------------------------------------------------------
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		g, err := greeting.NewGreeting(ctx, "go-consumer", &greeting.GreetingArgs{Name: pulumi.String("Go")})
		if err != nil {
			return err
		}
		ctx.Export("message", g.Message)
		return nil
	})
}
