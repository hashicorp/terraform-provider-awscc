// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registry

import (
	"context"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var dataSourceRegistrationClosed bool
var dataSourceRegistry map[string]func(context.Context) (datasource.DataSource, error)
var dataSourceRegistryMu sync.Mutex

var resourceRegistrationClosed bool
var resourceRegistry map[string]func(context.Context) (resource.Resource, error)
var resourceRegistryMu sync.Mutex

// AddDataSourceFactory registers the specified data source type name and factory.
func AddDataSourceFactory(name string, factory func(context.Context) (datasource.DataSource, error)) {
	dataSourceRegistryMu.Lock()
	defer dataSourceRegistryMu.Unlock()

	if dataSourceRegistrationClosed {
		panic("Data Source registration is closed")
	}

	if dataSourceRegistry == nil {
		dataSourceRegistry = make(map[string]func(context.Context) (datasource.DataSource, error))
	}
	dataSourceRegistry[name] = factory
}

// AddResourceFactory registers the specified resource type name and factory.
func AddResourceFactory(name string, factory func(context.Context) (resource.Resource, error)) {
	resourceRegistryMu.Lock()
	defer resourceRegistryMu.Unlock()

	if resourceRegistrationClosed {
		panic("Resource registration is closed")
	}

	if resourceRegistry == nil {
		resourceRegistry = make(map[string]func(context.Context) (resource.Resource, error))
	}
	resourceRegistry[name] = factory
}

// DataSourceFactories returns the registered data source factories.
// Data Source registration is closed.
func DataSourceFactories() map[string]func(context.Context) (datasource.DataSource, error) {
	dataSourceRegistryMu.Lock()
	defer dataSourceRegistryMu.Unlock()

	dataSourceRegistrationClosed = true

	return dataSourceRegistry
}

// ResourceFactories returns the registered resource factories.
// Resource registration is closed.
func ResourceFactories() map[string]func(context.Context) (resource.Resource, error) {
	resourceRegistryMu.Lock()
	defer resourceRegistryMu.Unlock()

	resourceRegistrationClosed = true

	return resourceRegistry
}
