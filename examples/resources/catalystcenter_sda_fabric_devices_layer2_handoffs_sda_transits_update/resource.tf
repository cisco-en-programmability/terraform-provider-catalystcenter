
resource "catalystcenter_sda_fabric_devices_layer2_handoffs_sda_transits_update" "example" {
  provider = meraki
  parameters {

    affinity_id_decider               = 1
    affinity_id_prime                 = 1
    connected_to_internet             = "false"
    fabric_id                         = "string"
    is_multicast_over_transit_enabled = "false"
    network_device_id                 = "string"
    transit_network_id                = "string"
  }
}

output "catalystcenter_sda_fabric_devices_layer2_handoffs_sda_transits_update_example" {
  value = catalystcenter_sda_fabric_devices_layer2_handoffs_sda_transits_update.example
}