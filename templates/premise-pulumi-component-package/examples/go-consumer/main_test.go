package main

import (
	greeting "example.com/premise-module/sdk/go/greeting"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"testing"
)

// Public API -----------------------------------------------------------------
func TestGeneratedConsumer(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		g, err := greeting.NewGreeting(ctx, "consumer", &greeting.GreetingArgs{Name: pulumi.String("Go")})
		if err != nil {
			return err
		}
		g.Message.ApplyT(func(message string) string {
			if message != "Hello, Go!" {
				t.Errorf("message %q", message)
			}
			return message
		})
		return nil
	}, pulumi.WithMocks("consumer", "test", mocks{t}))
	if err != nil {
		t.Fatal(err)
	}
}

// Private helpers ------------------------------------------------------------
type mocks struct{ t *testing.T }

func (m mocks) NewResource(a pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	if a.TypeToken != "premise-pulumi-component-package:index:Greeting" {
		m.t.Errorf("unexpected token %s", a.TypeToken)
	}
	a.Inputs["message"] = resource.NewStringProperty("Hello, " + a.Inputs["name"].StringValue() + "!")
	return a.Name + "_id", a.Inputs, nil
}
func (m mocks) Call(a pulumi.MockCallArgs) (resource.PropertyMap, error) { return a.Args, nil }
