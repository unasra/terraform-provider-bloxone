package flex

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

type FrameworkElementFlExFunc[T any, U any] func(context.Context, T, *diag.Diagnostics) U

func FlattenString(s string) types.String {
	if s == "" {
		return types.StringNull()
	}
	return types.StringValue(s)
}

func FlattenStringPointer(s *string) types.String {
	if s == nil {
		return types.StringNull()
	}
	return FlattenString(*s)
}

func FlattenInt64(i int64) types.Int64 {
	if i == 0 {
		return types.Int64Null()
	}
	return types.Int64Value(i)
}

func FlattenInt64Pointer(i *int64) types.Int64 {
	if i == nil {
		return types.Int64Null()
	}
	return FlattenInt64(*i)
}

func FlattenFloat64(f float64) types.Float64 {
	if f == 0 {
		return types.Float64Null()
	}
	return types.Float64Value(f)
}

func FlattenFloat64Pointer(f *float64) types.Float64 {
	if f == nil {
		return types.Float64Null()
	}
	return FlattenFloat64(*f)
}

func FlattenFrameworkMapString(ctx context.Context, m map[string]interface{}, diags *diag.Diagnostics) types.Map {
	if len(m) == 0 {
		return types.MapNull(types.StringType)
	}
	tfMap, d := types.MapValueFrom(ctx, types.StringType, m)
	diags.Append(d...)
	return tfMap
}

func ExpandFrameworkListString(ctx context.Context, tfList types.List, diags *diag.Diagnostics) []string {
	if tfList.IsNull() || tfList.IsUnknown() {
		return nil
	}
	var data []string
	diags.Append(tfList.ElementsAs(ctx, &data, false)...)
	return data
}

func FlattenFrameworkListString(ctx context.Context, l []string, diags *diag.Diagnostics) types.List {
	if len(l) == 0 {
		return types.ListNull(types.StringType)
	}
	tfList, d := types.ListValueFrom(ctx, types.StringType, l)
	diags.Append(d...)
	return tfList
}

func FlattenFrameworkListNestedBlock[T any, U any](ctx context.Context, data []T, attrTypes map[string]attr.Type, diags *diag.Diagnostics, f FrameworkElementFlExFunc[*T, U]) types.List {
	if len(data) == 0 {
		return types.ListNull(types.ObjectType{AttrTypes: attrTypes})
	}

	tfData := ApplyToAll(data, func(t T) U {
		return f(ctx, &t, diags)
	})

	tfList, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: attrTypes}, tfData)
	diags.Append(d...)
	return tfList
}

func FlattenFrameworkNestedBlock[T any, U any](ctx context.Context, data *T, attrTypes map[string]attr.Type, diags *diag.Diagnostics, f FrameworkElementFlExFunc[*T, U]) types.Object {
	if data == nil {
		return types.ObjectNull(attrTypes)
	}
	u := f(ctx, data, diags)
	t, d := types.ObjectValueFrom(ctx, attrTypes, u)
	diags.Append(d...)
	return t
}

func ExpandTime(_ context.Context, dt timetypes.RFC3339, diags *diag.Diagnostics) time.Time {
	t, d := dt.ValueRFC3339Time()
	diags.Append(d...)
	return t
}

func ExpandTimePointer(_ context.Context, dt timetypes.RFC3339, diags *diag.Diagnostics) *time.Time {
	if dt.IsNull() || dt.IsUnknown() {
		return nil
	}
	t, d := dt.ValueRFC3339Time()
	diags.Append(d...)
	return &t
}

func ExpandFrameworkMapString(ctx context.Context, tfMap types.Map, diags *diag.Diagnostics) map[string]interface{} {
	if tfMap.IsNull() || tfMap.IsUnknown() {
		return nil
	}
	elements := make(map[string]string, len(tfMap.Elements()))
	diags.Append(tfMap.ElementsAs(ctx, &elements, false)...)

	elementsNew := make(map[string]interface{}, len(tfMap.Elements()))
	for k, v := range elements {
		elementsNew[k] = v
	}
	return elementsNew
}

func ExpandFrameworkListNestedBlock[T any, U any](ctx context.Context, tfList types.List, diags *diag.Diagnostics, f FrameworkElementFlExFunc[T, *U]) []U {
	if tfList.IsNull() || tfList.IsUnknown() {
		return nil
	}

	var data []T

	diags.Append(tfList.ElementsAs(ctx, &data, false)...)

	return ApplyToAll(data, func(t T) U {
		return *f(ctx, t, diags)
	})

}

func ExpandFrameworkMapFilterString(ctx context.Context, tfMap types.Map, diags *diag.Diagnostics) string {
	if tfMap.IsNull() || tfMap.IsUnknown() {
		return ""
	}

	elements := make(map[string]string, len(tfMap.Elements()))
	diags.Append(tfMap.ElementsAs(ctx, &elements, false)...)

	var filters []string
	for k, v := range elements {
		filters = append(filters, fmt.Sprintf("%s=='%s'", k, v))
	}
	filterStr := strings.Join(filters, " and ")
	return filterStr
}

// ApplyToAll returns a new slice containing the results of applying the function `f` to each element of the original slice `s`.
func ApplyToAll[T, U any](s []T, f func(T) U) []U {
	v := make([]U, len(s))

	for i, e := range s {
		v[i] = f(e)
	}

	return v
}

func ExpandString(v types.String) string {
	return v.ValueString()
}

func ExpandStringPointer(v types.String) *string {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	return v.ValueStringPointer()
}

func ExpandInt64(v types.Int64) int64 {
	return v.ValueInt64()
}

func ExpandInt64Pointer(v types.Int64) *int64 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	return v.ValueInt64Pointer()
}

func ExpandFloat64(v types.Float64) float64 {
	return v.ValueFloat64()
}

func ExpandFloat64Pointer(v types.Float64) *float64 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	return v.ValueFloat64Pointer()
}

func ExpandFloat32(v types.Float64) float32 {
	return float32(v.ValueFloat64())
}

func ExpandFloat32Pointer(v types.Float64) *float32 {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	return utils.Ptr(float32(v.ValueFloat64()))
}

func ExpandBool(v types.Bool) bool {
	return v.ValueBool()
}

func ExpandBoolPointer(v types.Bool) *bool {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	return v.ValueBoolPointer()
}
