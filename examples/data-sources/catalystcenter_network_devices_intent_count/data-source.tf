
data "catalystcenter_network_devices_intent_count" "example" {
  provider            = catalystcenter
  family              = "string"
  id                  = "string"
  management_address  = "string"
  management_state    = "string"
  reachability_status = "string"
  role                = "string"
  serial_number       = "string"
  stack_device        = "string"
  status              = "string"
}

output "catalystcenter_network_devices_intent_count_example" {
  value = data.catalystcenter_network_devices_intent_count.example.item
}
