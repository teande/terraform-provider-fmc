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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcPolicyBasedRouting(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" && os.Getenv("TF_VAR_outside_int_id") == "" && os.Getenv("TF_VAR_outside2_int_id") == "" && os.Getenv("TF_VAR_inside_int_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id or TF_VAR_outside_int_id or TF_VAR_outside2_int_id or TF_VAR_inside_int_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_policy_based_routing.test", "forwarding_actions.0.forwarding_action_type", "SET_EGRESS_INTF_BY_PRIORITY"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcPolicyBasedRoutingPrerequisitesConfig + testAccFmcPolicyBasedRoutingConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcPolicyBasedRoutingPrerequisitesConfig + testAccFmcPolicyBasedRoutingConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcPolicyBasedRoutingPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "outside_int_id" { default = null } // tests will set $TF_VAR_outside_int_id
variable "outside2_int_id" { default = null } // tests will set $TF_VAR_outside2_int_id
variable "inside_int_id" { default = null } // tests will set $TF_VAR_inside_int_id

resource "fmc_port" "test" {
  name = "myport2"
  protocol = "TCP"
  port = "65000"
}

resource "fmc_extended_acl" "_acl_test" {
  name        = "extended_acl_1"
  description = "My Extended Access Control List"
  entries = [
  {
      action               = "DENY"
      log_level            = "WARNING"
      logging              = "PER_ACCESS_LIST_ENTRY"
      log_interval_seconds = 120
      source_network_literals = [
        {
          value = "10.1.1.0/24"
          type  = "Network"
        }
      ]
      destination_network_literals = [
        {
          value = "10.2.2.2"
          type  = "Host"
        }
      ]
      source_port_objects = [
        {
          id = fmc_port.test.id
        }
      ]
      destination_port_objects = [
        {
          id = fmc_port.test.id
        }
      ]
    }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcPolicyBasedRoutingConfig_minimum() string {
	config := `resource "fmc_policy_based_routing" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	ingress_interfaces = [{` + "\n"
	config += `		id = var.inside_int_id` + "\n"
	config += `	}]` + "\n"
	config += `	forwarding_actions = [{` + "\n"
	config += `		forwarding_action_type = "SET_EGRESS_INTF_BY_PRIORITY"` + "\n"
	config += `		match_criteria_id = fmc_extended_acl._acl_test.id` + "\n"
	config += `		egress_interfaces = [{` + "\n"
	config += `		id = var.outside_int_id` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcPolicyBasedRoutingConfig_all() string {
	config := `resource "fmc_policy_based_routing" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	ingress_interfaces = [{` + "\n"
	config += `		id = var.inside_int_id` + "\n"
	config += `	}]` + "\n"
	config += `	forwarding_actions = [{` + "\n"
	config += `		forwarding_action_type = "SET_EGRESS_INTF_BY_PRIORITY"` + "\n"
	config += `		match_criteria_name = fmc_extended_acl._acl_test.name` + "\n"
	config += `		match_criteria_id = fmc_extended_acl._acl_test.id` + "\n"
	config += `		egress_interfaces = [{` + "\n"
	config += `			id = var.outside_int_id` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
