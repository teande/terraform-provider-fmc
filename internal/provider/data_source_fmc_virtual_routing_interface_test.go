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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcVirtualRoutingInterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" && os.Getenv("TF_VAR_outside_int_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id or TF_VAR_outside_int_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_virtual_routing_interface.test", "name", "vrf_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_virtual_routing_interface.test", "description", "My virtual router"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVirtualRoutingInterfacePrerequisitesConfig + testAccDataSourceFmcVirtualRoutingInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcVirtualRoutingInterfacePrerequisitesConfig + testAccNamedDataSourceFmcVirtualRoutingInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcVirtualRoutingInterfacePrerequisitesConfig = `
variable device_id {}
variable outside_int_id {}

data "fmc_device_physical_interface" "example" {
  id        = var.outside_int_id 
  device_id = var.device_id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVirtualRoutingInterfaceConfig() string {
	config := `resource "fmc_virtual_routing_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	name = "vrf_1"` + "\n"
	config += `	description = "My virtual router"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		id = var.outside_int_id` + "\n"
	config += `		name = data.fmc_device_physical_interface.example.name` + "\n"
	config += `		ifname = data.fmc_device_physical_interface.example.logical_name` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_virtual_routing_interface" "test" {
			id = fmc_virtual_routing_interface.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcVirtualRoutingInterfaceConfig() string {
	config := `resource "fmc_virtual_routing_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	name = "vrf_1"` + "\n"
	config += `	description = "My virtual router"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		id = var.outside_int_id` + "\n"
	config += `		name = data.fmc_device_physical_interface.example.name` + "\n"
	config += `		ifname = data.fmc_device_physical_interface.example.logical_name` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_virtual_routing_interface" "test" {
			name = fmc_virtual_routing_interface.test.name
			device_id = var.device_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
