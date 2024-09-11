
resource "catalystcenter_sda_transit_networks" "example" {
  provider = catalystcenter

  parameters {

    id = "string"
    ip_transit_settings {

      autonomous_system_number = "string"
      routing_protocol_name    = "string"
    }
    name = "string"
    sda_transit_settings {

      control_plane_network_device_ids  = ["string"]
      is_multicast_over_transit_enabled = "false"
    }
    type = "string"
  }
}

output "catalystcenter_sda_transit_networks_example" {
  value = catalystcenter_sda_transit_networks.example
}