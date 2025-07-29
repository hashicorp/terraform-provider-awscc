# Custom Plan Modifiers

This package provides custom plan modifiers to address shadow drift issues in the AWSCC provider. These replace the framework's `UseStateForUnknown` to avoid false positive drift detection.

**Reference**: https://github.com/hashicorp/terraform-provider-awscc/issues/2726

## File Structure

| File | Types | Functions |
|------|-------|-----------|
| `string_plan_modifiers.go` | String, Bool | `CustomUseStateForUnknownString()`, `CustomUseStateForUnknownBool()` |
| `numeric_plan_modifiers.go` | Int64, Float64, Number | `CustomUseStateForUnknownInt64()`, `CustomUseStateForUnknownFloat64()`, `CustomUseStateForUnknownNumber()` |
| `collection_plan_modifiers.go` | List, Set, Map | `CustomUseStateForUnknownList()`, `CustomUseStateForUnknownSet()`, `CustomUseStateForUnknownMap()` |
| `object_plan_modifiers.go` | Object | `CustomUseStateForUnknownObject()` |

## Usage

```go
import "github.com/hashicorp/terraform-provider-awscc/internal/generic"

// String attribute
"name": schema.StringAttribute{
    Optional: true,
    Computed: true,
    PlanModifiers: []planmodifier.String{
        generic.CustomUseStateForUnknownString(),
    },
},

// Integer attribute
"port": schema.Int64Attribute{
    Optional: true,
    Computed: true,
    PlanModifiers: []planmodifier.Int64{
        generic.CustomUseStateForUnknownInt64(),
    },
},

// Object attribute
"configuration": schema.SingleNestedAttribute{
    Optional: true,
    Computed: true,
    PlanModifiers: []planmodifier.Object{
        generic.CustomUseStateForUnknownObject(),
    },
},
```

## How It Works

These custom modifiers prevent shadow drift by:
1. Using the state value when configuration is null
2. Preventing the framework from marking attributes as "unknown"
3. Avoiding false positive drift detection in computed attributes
