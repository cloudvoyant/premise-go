package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/blang/semver"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/integration"
)

// Public API -----------------------------------------------------------------
func TestProviderSchema(t *testing.T) {
	v, err := provider()
	if err != nil {
		t.Fatal(err)
	}
	server, err := integration.NewServer(context.Background(), "premise-pulumi-component-package", semver.MustParse("0.1.0"), integration.WithProvider(v))
	if err != nil {
		t.Fatal(err)
	}
	schema, err := server.GetSchema(p.GetSchemaRequest{})
	if err != nil {
		t.Fatal(err)
	}
	var decoded struct {
		Resources map[string]struct {
			IsComponent    bool     `json:"isComponent"`
			RequiredInputs []string `json:"requiredInputs"`
		} `json:"resources"`
	}
	if err := json.Unmarshal([]byte(schema.Schema), &decoded); err != nil {
		t.Fatal(err)
	}
	greeting, ok := decoded.Resources["premise-pulumi-component-package:index:Greeting"]
	if !ok || !greeting.IsComponent || len(greeting.RequiredInputs) != 1 || greeting.RequiredInputs[0] != "name" {
		t.Fatalf("unexpected component schema: %#v", decoded)
	}
}
