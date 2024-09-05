// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type VirtualRoutingInterface struct {
	Id          types.String                        `tfsdk:"id"`
	Domain      types.String                        `tfsdk:"domain"`
	DeviceId    types.String                        `tfsdk:"device_id"`
	Name        types.String                        `tfsdk:"name"`
	Description types.String                        `tfsdk:"description"`
	Interfaces  []VirtualRoutingInterfaceInterfaces `tfsdk:"interfaces"`
}

type VirtualRoutingInterfaceInterfaces struct {
	Id     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Ifname types.String `tfsdk:"ifname"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VirtualRoutingInterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/virtualrouters", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VirtualRoutingInterface) toBody(ctx context.Context, state VirtualRoutingInterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, "interfaces", []interface{}{})
		for _, item := range data.Interfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Ifname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ifname", item.Ifname.ValueString())
			}
			body, _ = sjson.SetRaw(body, "interfaces.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VirtualRoutingInterface) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("interfaces"); value.Exists() {
		data.Interfaces = make([]VirtualRoutingInterfaceInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VirtualRoutingInterfaceInterfaces{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("ifname"); cValue.Exists() {
				item.Ifname = types.StringValue(cValue.String())
			} else {
				item.Ifname = types.StringNull()
			}
			data.Interfaces = append(data.Interfaces, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VirtualRoutingInterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := 0; i < len(data.Interfaces); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Interfaces[i].Id.ValueString()}

		var r gjson.Result
		res.Get("interfaces").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if !r.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing data.Interfaces[%d] = %+v",
				i,
				data.Interfaces[i],
			))
			data.Interfaces = slices.Delete(data.Interfaces, i, i+1)
			i--

			continue
		}
		if value := r.Get("id"); value.Exists() && !data.Interfaces[i].Id.IsNull() {
			data.Interfaces[i].Id = types.StringValue(value.String())
		} else {
			data.Interfaces[i].Id = types.StringNull()
		}
		if value := r.Get("name"); value.Exists() && !data.Interfaces[i].Name.IsNull() {
			data.Interfaces[i].Name = types.StringValue(value.String())
		} else {
			data.Interfaces[i].Name = types.StringNull()
		}
		if value := r.Get("ifname"); value.Exists() && !data.Interfaces[i].Ifname.IsNull() {
			data.Interfaces[i].Ifname = types.StringValue(value.String())
		} else {
			data.Interfaces[i].Ifname = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VirtualRoutingInterface) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
