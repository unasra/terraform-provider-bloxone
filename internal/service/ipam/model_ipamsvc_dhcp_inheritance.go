package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcDHCPInheritanceModel struct {
	AsmConfig                  types.Object `tfsdk:"asm_config"`
	DdnsClientUpdate           types.Object `tfsdk:"ddns_client_update"`
	DdnsConflictResolutionMode types.Object `tfsdk:"ddns_conflict_resolution_mode"`
	DdnsEnabled                types.Object `tfsdk:"ddns_enabled"`
	DdnsHostnameBlock          types.Object `tfsdk:"ddns_hostname_block"`
	DdnsTtlPercent             types.Object `tfsdk:"ddns_ttl_percent"`
	DdnsUpdateBlock            types.Object `tfsdk:"ddns_update_block"`
	DdnsUpdateOnRenew          types.Object `tfsdk:"ddns_update_on_renew"`
	DdnsUseConflictResolution  types.Object `tfsdk:"ddns_use_conflict_resolution"`
	DhcpConfig                 types.Object `tfsdk:"dhcp_config"`
	DhcpOptions                types.Object `tfsdk:"dhcp_options"`
	HeaderOptionFilename       types.Object `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress  types.Object `tfsdk:"header_option_server_address"`
	HeaderOptionServerName     types.Object `tfsdk:"header_option_server_name"`
	HostnameRewriteBlock       types.Object `tfsdk:"hostname_rewrite_block"`
}

var IpamsvcDHCPInheritanceAttrTypes = map[string]attr.Type{
	"asm_config":                    types.ObjectType{AttrTypes: IpamsvcInheritedASMConfigAttrTypes},
	"ddns_client_update":            types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"ddns_conflict_resolution_mode": types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"ddns_enabled":                  types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"ddns_hostname_block":           types.ObjectType{AttrTypes: IpamsvcInheritedDDNSHostnameBlockAttrTypes},
	"ddns_ttl_percent":              types.ObjectType{AttrTypes: InheritanceInheritedFloatAttrTypes},
	"ddns_update_block":             types.ObjectType{AttrTypes: IpamsvcInheritedDDNSUpdateBlockAttrTypes},
	"ddns_update_on_renew":          types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"ddns_use_conflict_resolution":  types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"dhcp_config":                   types.ObjectType{AttrTypes: IpamsvcInheritedDHCPConfigAttrTypes},
	"dhcp_options":                  types.ObjectType{AttrTypes: IpamsvcInheritedDHCPOptionListAttrTypes},
	"header_option_filename":        types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_address":  types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_name":     types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"hostname_rewrite_block":        types.ObjectType{AttrTypes: IpamsvcInheritedHostnameRewriteBlockAttrTypes},
}

var IpamsvcDHCPInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"asm_config": schema.SingleNestedAttribute{
		Attributes: IpamsvcInheritedASMConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"ddns_client_update": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
	"ddns_conflict_resolution_mode": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
	"ddns_enabled": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:   true,
	},
	"ddns_hostname_block": schema.SingleNestedAttribute{
		Attributes: IpamsvcInheritedDDNSHostnameBlockResourceSchemaAttributes,
		Optional:   true,
	},
	"ddns_ttl_percent": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedFloatResourceSchemaAttributes,
		Optional:   true,
	},
	"ddns_update_block": schema.SingleNestedAttribute{
		Attributes: IpamsvcInheritedDDNSUpdateBlockResourceSchemaAttributes,
		Optional:   true,
	},
	"ddns_update_on_renew": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:   true,
	},
	"ddns_use_conflict_resolution": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:   true,
	},
	"dhcp_config": schema.SingleNestedAttribute{
		Attributes: IpamsvcInheritedDHCPConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"dhcp_options": schema.SingleNestedAttribute{
		Attributes: IpamsvcInheritedDHCPOptionListResourceSchemaAttributes,
		Optional:   true,
	},
	"header_option_filename": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
	"header_option_server_address": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
	"header_option_server_name": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
	"hostname_rewrite_block": schema.SingleNestedAttribute{
		Attributes: IpamsvcInheritedHostnameRewriteBlockResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandIpamsvcDHCPInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDHCPInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcDHCPInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcDHCPInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDHCPInheritance {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcDHCPInheritance{
		AsmConfig:                  ExpandIpamsvcInheritedASMConfig(ctx, m.AsmConfig, diags),
		DdnsClientUpdate:           ExpandInheritanceInheritedString(ctx, m.DdnsClientUpdate, diags),
		DdnsConflictResolutionMode: ExpandInheritanceInheritedString(ctx, m.DdnsConflictResolutionMode, diags),
		DdnsEnabled:                ExpandInheritanceInheritedBool(ctx, m.DdnsEnabled, diags),
		DdnsHostnameBlock:          ExpandIpamsvcInheritedDDNSHostnameBlock(ctx, m.DdnsHostnameBlock, diags),
		DdnsTtlPercent:             ExpandInheritanceInheritedFloat(ctx, m.DdnsTtlPercent, diags),
		DdnsUpdateBlock:            ExpandIpamsvcInheritedDDNSUpdateBlock(ctx, m.DdnsUpdateBlock, diags),
		DdnsUpdateOnRenew:          ExpandInheritanceInheritedBool(ctx, m.DdnsUpdateOnRenew, diags),
		DdnsUseConflictResolution:  ExpandInheritanceInheritedBool(ctx, m.DdnsUseConflictResolution, diags),
		DhcpConfig:                 ExpandIpamsvcInheritedDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                ExpandIpamsvcInheritedDHCPOptionList(ctx, m.DhcpOptions, diags),
		HeaderOptionFilename:       ExpandInheritanceInheritedString(ctx, m.HeaderOptionFilename, diags),
		HeaderOptionServerAddress:  ExpandInheritanceInheritedString(ctx, m.HeaderOptionServerAddress, diags),
		HeaderOptionServerName:     ExpandInheritanceInheritedString(ctx, m.HeaderOptionServerName, diags),
		HostnameRewriteBlock:       ExpandIpamsvcInheritedHostnameRewriteBlock(ctx, m.HostnameRewriteBlock, diags),
	}
	return to
}

func FlattenIpamsvcDHCPInheritance(ctx context.Context, from *ipam.IpamsvcDHCPInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDHCPInheritanceAttrTypes)
	}
	m := IpamsvcDHCPInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDHCPInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDHCPInheritanceModel) Flatten(ctx context.Context, from *ipam.IpamsvcDHCPInheritance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDHCPInheritanceModel{}
	}
	m.AsmConfig = FlattenIpamsvcInheritedASMConfig(ctx, from.AsmConfig, diags)
	m.DdnsClientUpdate = FlattenInheritanceInheritedString(ctx, from.DdnsClientUpdate, diags)
	m.DdnsConflictResolutionMode = FlattenInheritanceInheritedString(ctx, from.DdnsConflictResolutionMode, diags)
	m.DdnsEnabled = FlattenInheritanceInheritedBool(ctx, from.DdnsEnabled, diags)
	m.DdnsHostnameBlock = FlattenIpamsvcInheritedDDNSHostnameBlock(ctx, from.DdnsHostnameBlock, diags)
	m.DdnsTtlPercent = FlattenInheritanceInheritedFloat(ctx, from.DdnsTtlPercent, diags)
	m.DdnsUpdateBlock = FlattenIpamsvcInheritedDDNSUpdateBlock(ctx, from.DdnsUpdateBlock, diags)
	m.DdnsUpdateOnRenew = FlattenInheritanceInheritedBool(ctx, from.DdnsUpdateOnRenew, diags)
	m.DdnsUseConflictResolution = FlattenInheritanceInheritedBool(ctx, from.DdnsUseConflictResolution, diags)
	m.DhcpConfig = FlattenIpamsvcInheritedDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = FlattenIpamsvcInheritedDHCPOptionList(ctx, from.DhcpOptions, diags)
	m.HeaderOptionFilename = FlattenInheritanceInheritedString(ctx, from.HeaderOptionFilename, diags)
	m.HeaderOptionServerAddress = FlattenInheritanceInheritedString(ctx, from.HeaderOptionServerAddress, diags)
	m.HeaderOptionServerName = FlattenInheritanceInheritedString(ctx, from.HeaderOptionServerName, diags)
	m.HostnameRewriteBlock = FlattenIpamsvcInheritedHostnameRewriteBlock(ctx, from.HostnameRewriteBlock, diags)
}
