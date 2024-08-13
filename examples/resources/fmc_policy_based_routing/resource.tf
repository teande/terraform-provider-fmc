resource "fmc_policy_based_routing" "example" {
  device_id   = "76d24097-41c4-4558-a4d0-a8c07ac08470"
  name        = "PBR-req-test (dummy)"
  description = "By provider"
  ingress_interfaces = [
    {
      id = "76d24097-41c4-4558-a4d0-a8c07ac08470"
    }
  ]
  forwarding_actions = [
    {
      forwarding_action_type = "SET_EGRESS_INTF_BY_PRIORITY"
      match_criteria_name    = "box-acl"
      match_criteria_id      = "000C29A7-0BEE-0ed3-0000-008589983100"
      egress_interfaces = [
        {
          id = ""
        }
      ]
      default_interface = true
    }
  ]
}
