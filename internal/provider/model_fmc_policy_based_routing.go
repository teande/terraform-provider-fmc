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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type PolicyBasedRouting struct {
	Id                types.String                          `tfsdk:"id"`
	Domain            types.String                          `tfsdk:"domain"`
	DeviceId          types.String                          `tfsdk:"device_id"`
	Name              types.String                          `tfsdk:"name"`
	Description       types.String                          `tfsdk:"description"`
	IngressInterfaces []PolicyBasedRoutingIngressInterfaces `tfsdk:"ingress_interfaces"`
	ForwardingActions []PolicyBasedRoutingForwardingActions `tfsdk:"forwarding_actions"`
}

type PolicyBasedRoutingIngressInterfaces struct {
	Id     types.String `tfsdk:"id"`
	Ifname types.String `tfsdk:"ifname"`
	Name   types.String `tfsdk:"name"`
}

type PolicyBasedRoutingForwardingActions struct {
	ForwardingActionType types.String                                          `tfsdk:"forwarding_action_type"`
	MatchCriteriaName    types.String                                          `tfsdk:"match_criteria_name"`
	MatchCriteriaId      types.String                                          `tfsdk:"match_criteria_id"`
	EgressInterfaces     []PolicyBasedRoutingForwardingActionsEgressInterfaces `tfsdk:"egress_interfaces"`
	DefaultInterface     types.Bool                                            `tfsdk:"default_interface"`
}

type PolicyBasedRoutingForwardingActionsEgressInterfaces struct {
	Id     types.String `tfsdk:"id"`
	Ifname types.String `tfsdk:"ifname"`
	Name   types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data PolicyBasedRouting) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/policybasedroutes", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data PolicyBasedRouting) toBody(ctx context.Context, state PolicyBasedRouting) string {
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
	if len(data.IngressInterfaces) > 0 {
		body, _ = sjson.Set(body, "ingressInterfaces", []interface{}{})
		for _, item := range data.IngressInterfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Ifname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ifname", item.Ifname.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ingressInterfaces.-1", itemBody)
		}
	}
	if len(data.ForwardingActions) > 0 {
		body, _ = sjson.Set(body, "forwardingActions", []interface{}{})
		for _, item := range data.ForwardingActions {
			itemBody := ""
			if !item.ForwardingActionType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "forwardingActionType", item.ForwardingActionType.ValueString())
			}
			if !item.MatchCriteriaName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchCriteria.name", item.MatchCriteriaName.ValueString())
			}
			if !item.MatchCriteriaId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchCriteria.id", item.MatchCriteriaId.ValueString())
			}
			if len(item.EgressInterfaces) > 0 {
				itemBody, _ = sjson.Set(itemBody, "egressInterfaces", []interface{}{})
				for _, childItem := range item.EgressInterfaces {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Ifname.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ifname", childItem.Ifname.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "egressInterfaces.-1", itemChildBody)
				}
			}
			if !item.DefaultInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "defaultInterface", item.DefaultInterface.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "forwardingActions.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *PolicyBasedRouting) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("ingressInterfaces"); value.Exists() {
		data.IngressInterfaces = make([]PolicyBasedRoutingIngressInterfaces, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyBasedRoutingIngressInterfaces{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("ifname"); cValue.Exists() {
				item.Ifname = types.StringValue(cValue.String())
			} else {
				item.Ifname = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			data.IngressInterfaces = append(data.IngressInterfaces, item)
			return true
		})
	}
	if value := res.Get("forwardingActions"); value.Exists() {
		data.ForwardingActions = make([]PolicyBasedRoutingForwardingActions, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PolicyBasedRoutingForwardingActions{}
			if cValue := v.Get("forwardingActionType"); cValue.Exists() {
				item.ForwardingActionType = types.StringValue(cValue.String())
			} else {
				item.ForwardingActionType = types.StringNull()
			}
			if cValue := v.Get("matchCriteria.name"); cValue.Exists() {
				item.MatchCriteriaName = types.StringValue(cValue.String())
			} else {
				item.MatchCriteriaName = types.StringNull()
			}
			if cValue := v.Get("matchCriteria.id"); cValue.Exists() {
				item.MatchCriteriaId = types.StringValue(cValue.String())
			} else {
				item.MatchCriteriaId = types.StringNull()
			}
			if cValue := v.Get("egressInterfaces"); cValue.Exists() {
				item.EgressInterfaces = make([]PolicyBasedRoutingForwardingActionsEgressInterfaces, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := PolicyBasedRoutingForwardingActionsEgressInterfaces{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("ifname"); ccValue.Exists() {
						cItem.Ifname = types.StringValue(ccValue.String())
					} else {
						cItem.Ifname = types.StringNull()
					}
					if ccValue := cv.Get("name"); ccValue.Exists() {
						cItem.Name = types.StringValue(ccValue.String())
					} else {
						cItem.Name = types.StringNull()
					}
					item.EgressInterfaces = append(item.EgressInterfaces, cItem)
					return true
				})
			}
			if cValue := v.Get("defaultInterface"); cValue.Exists() {
				item.DefaultInterface = types.BoolValue(cValue.Bool())
			} else {
				item.DefaultInterface = types.BoolNull()
			}
			data.ForwardingActions = append(data.ForwardingActions, item)
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
func (data *PolicyBasedRouting) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	for i := 0; i < len(data.IngressInterfaces); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.IngressInterfaces[i].Id.ValueString()}

		var r gjson.Result
		res.Get("ingressInterfaces").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing data.IngressInterfaces[%d] = %+v",
				i,
				data.IngressInterfaces[i],
			))
			data.IngressInterfaces = slices.Delete(data.IngressInterfaces, i, i+1)
			i--

			continue
		}
		if value := r.Get("id"); value.Exists() && !data.IngressInterfaces[i].Id.IsNull() {
			data.IngressInterfaces[i].Id = types.StringValue(value.String())
		} else {
			data.IngressInterfaces[i].Id = types.StringNull()
		}
		if value := r.Get("ifname"); value.Exists() && !data.IngressInterfaces[i].Ifname.IsNull() {
			data.IngressInterfaces[i].Ifname = types.StringValue(value.String())
		} else {
			data.IngressInterfaces[i].Ifname = types.StringNull()
		}
		if value := r.Get("name"); value.Exists() && !data.IngressInterfaces[i].Name.IsNull() {
			data.IngressInterfaces[i].Name = types.StringValue(value.String())
		} else {
			data.IngressInterfaces[i].Name = types.StringNull()
		}
	}
	for i := 0; i < len(data.ForwardingActions); i++ {
		keys := [...]string{"forwardingActionType", "matchCriteria.name", "matchCriteria.id", "defaultInterface"}
		keyValues := [...]string{data.ForwardingActions[i].ForwardingActionType.ValueString(), data.ForwardingActions[i].MatchCriteriaName.ValueString(), data.ForwardingActions[i].MatchCriteriaId.ValueString(), strconv.FormatBool(data.ForwardingActions[i].DefaultInterface.ValueBool())}

		var r gjson.Result
		res.Get("forwardingActions").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing data.ForwardingActions[%d] = %+v",
				i,
				data.ForwardingActions[i],
			))
			data.ForwardingActions = slices.Delete(data.ForwardingActions, i, i+1)
			i--

			continue
		}
		if value := r.Get("forwardingActionType"); value.Exists() && !data.ForwardingActions[i].ForwardingActionType.IsNull() {
			data.ForwardingActions[i].ForwardingActionType = types.StringValue(value.String())
		} else {
			data.ForwardingActions[i].ForwardingActionType = types.StringNull()
		}
		if value := r.Get("matchCriteria.name"); value.Exists() && !data.ForwardingActions[i].MatchCriteriaName.IsNull() {
			data.ForwardingActions[i].MatchCriteriaName = types.StringValue(value.String())
		} else {
			data.ForwardingActions[i].MatchCriteriaName = types.StringNull()
		}
		if value := r.Get("matchCriteria.id"); value.Exists() && !data.ForwardingActions[i].MatchCriteriaId.IsNull() {
			data.ForwardingActions[i].MatchCriteriaId = types.StringValue(value.String())
		} else {
			data.ForwardingActions[i].MatchCriteriaId = types.StringNull()
		}
		for ci := 0; ci < len(data.ForwardingActions[i].EgressInterfaces); ci++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.ForwardingActions[i].EgressInterfaces[ci].Id.ValueString()}

			var cr gjson.Result
			r.Get("egressInterfaces").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.ForwardingActions[i].EgressInterfaces[%d] = %+v",
					ci,
					data.ForwardingActions[i].EgressInterfaces[ci],
				))
				data.ForwardingActions[i].EgressInterfaces = slices.Delete(data.ForwardingActions[i].EgressInterfaces, ci, ci+1)
				ci--

				continue
			}
			if value := cr.Get("id"); value.Exists() && !data.ForwardingActions[i].EgressInterfaces[ci].Id.IsNull() {
				data.ForwardingActions[i].EgressInterfaces[ci].Id = types.StringValue(value.String())
			} else {
				data.ForwardingActions[i].EgressInterfaces[ci].Id = types.StringNull()
			}
			if value := cr.Get("ifname"); value.Exists() && !data.ForwardingActions[i].EgressInterfaces[ci].Ifname.IsNull() {
				data.ForwardingActions[i].EgressInterfaces[ci].Ifname = types.StringValue(value.String())
			} else {
				data.ForwardingActions[i].EgressInterfaces[ci].Ifname = types.StringNull()
			}
			if value := cr.Get("name"); value.Exists() && !data.ForwardingActions[i].EgressInterfaces[ci].Name.IsNull() {
				data.ForwardingActions[i].EgressInterfaces[ci].Name = types.StringValue(value.String())
			} else {
				data.ForwardingActions[i].EgressInterfaces[ci].Name = types.StringNull()
			}
		}
		if value := r.Get("defaultInterface"); value.Exists() && !data.ForwardingActions[i].DefaultInterface.IsNull() {
			data.ForwardingActions[i].DefaultInterface = types.BoolValue(value.Bool())
		} else {
			data.ForwardingActions[i].DefaultInterface = types.BoolNull()
		}
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *PolicyBasedRouting) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
