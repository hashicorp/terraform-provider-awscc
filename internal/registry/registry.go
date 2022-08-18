package registry

import (
	"context"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/provider"
)

var dataSourceRegisrationClosed bool
var dataSourceRegistry map[string]func(context.Context) (provider.DataSourceType, error)
var dataSourceRegistryMu sync.Mutex

var resourceRegisrationClosed bool
var resourceRegistry map[string]func(context.Context) (provider.ResourceType, error)
var resourceRegistryMu sync.Mutex

// AddDataSourceTypeFactory registers the specified data source type name and factory.
func AddDataSourceTypeFactory(name string, factory func(context.Context) (provider.DataSourceType, error)) {
	dataSourceRegistryMu.Lock()
	defer dataSourceRegistryMu.Unlock()

	if dataSourceRegisrationClosed {
		panic("Data Source registration is closed")
	}

	if dataSourceRegistry == nil {
		dataSourceRegistry = make(map[string]func(context.Context) (provider.DataSourceType, error))
	}
	dataSourceRegistry[name] = factory
}

// AddResourceTypeFactory registers the specified resource type name and factory.
func AddResourceTypeFactory(name string, factory func(context.Context) (provider.ResourceType, error)) {
	resourceRegistryMu.Lock()
	defer resourceRegistryMu.Unlock()

	if resourceRegisrationClosed {
		panic("Resource registration is closed")
	}

	if resourceRegistry == nil {
		resourceRegistry = make(map[string]func(context.Context) (provider.ResourceType, error))
	}
	resourceRegistry[name] = factory
}

// ResourceFactories returns the registered resource factories.
// Resource registration is closed.
func ResourceFactories() map[string]func(context.Context) (provider.ResourceType, error) {
	resourceRegistryMu.Lock()
	defer resourceRegistryMu.Unlock()

	resourceRegisrationClosed = true

	return resourceRegistry
}

// DataSourceFactories returns the registered data source factories.
// Data Source registration is closed.
func DataSourceFactories() map[string]func(context.Context) (provider.DataSourceType, error) {
	dataSourceRegistryMu.Lock()
	defer dataSourceRegistryMu.Unlock()

	dataSourceRegisrationClosed = true

	return dataSourceRegistry
}
