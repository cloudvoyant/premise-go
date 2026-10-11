package component

import (
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"testing"
)

// Public API -----------------------------------------------------------------
func TestComponent(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		c, err := NewGreeting(ctx, "test", Args{Name: pulumi.String("Ada")})
		if err != nil {
			return err
		}
		c.Message.ApplyT(func(s string) string {
			if s != "Hello, Ada!" {
				t.Errorf("message %q", s)
			}
			return s
		})
		return nil
	}, pulumi.WithMocks("example", "test", mocks{t}))
	if err != nil {
		t.Fatal(err)
	}
}

// Private helpers ------------------------------------------------------------
type mocks struct{ t *testing.T }

func (m mocks) NewResource(a pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	if a.TypeToken == "premise:index:Message" {
		if a.RegisterRPC.GetParent() == "" {
			m.t.Error("child has no parent")
		}
	}
	return a.Name + "_id", a.Inputs, nil
}
func (m mocks) Call(a pulumi.MockCallArgs) (resource.PropertyMap, error) { return a.Args, nil }
