package registry

import (
	"context"
	"sync"

	tfsdk "github.com/hashicorp/terraform-plugin-framework"
)

var resourceRegisrationClosed bool
var resourceRegistry map[string]func(context.Context) (tfsdk.ResourceType, error)
var resourceRegistryMu sync.Mutex

// AddResourceTypeFactory registers the specified resource type name and factory.
func AddResourceTypeFactory(name string, factory func(context.Context) (tfsdk.ResourceType, error)) {
	resourceRegistryMu.Lock()
	defer resourceRegistryMu.Unlock()

	if resourceRegisrationClosed {
		panic("Resource registration is closed")
	}

	if resourceRegistry == nil {
		resourceRegistry = make(map[string]func(context.Context) (tfsdk.ResourceType, error))
	}
	resourceRegistry[name] = factory
}

// ResourceFactories returns the registered resource factories.
// Resource registration is closed.
func ResourceFactories() map[string]func(context.Context) (tfsdk.ResourceType, error) {
	resourceRegistryMu.Lock()
	defer resourceRegistryMu.Unlock()

	resourceRegisrationClosed = true

	return resourceRegistry
}
