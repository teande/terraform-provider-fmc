resource "fmc_virtual_routing_interface" "example" {
  device_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name        = "vrf_1"
  description = "My virtual router"
  interfaces = [
    {
      id     = "76d24097-41c4-4558-a4d0-a8c07ac08470"
      name   = "TenGigabitEthernet0/1"
      ifname = "Outside"
    }
  ]
}
