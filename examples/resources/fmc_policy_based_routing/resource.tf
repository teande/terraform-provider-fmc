resource "fmc_policy_based_routing" "example" {
  device_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name        = "PBR-req-test"
  description = "By provider"
  ingress_interfaces = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  forwarding_actions = [
    {
      forwarding_action_type = "SET_EGRESS_INTF_BY_PRIORITY"
      match_criteria_name    = "my-acl"
      match_criteria_id      = "0AFFE3D6-879D-0ed3-0000-004295026024"
      egress_interfaces = [
        {
          id = ""
        }
      ]
      default_interface = true
    }
  ]
}
