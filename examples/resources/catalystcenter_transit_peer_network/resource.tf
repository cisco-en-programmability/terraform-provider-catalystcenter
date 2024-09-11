
resource "catalystcenter_transit_peer_network" "example" {
  provider = catalystcenter

  parameters {

    ip_transit_settings {

      autonomous_system_number = "string"
      routing_protocol_name    = "string"
    }
    sda_transit_settings {

      transit_control_plane_settings {

        device_management_ip_address = "string"
        site_name_hierarchy          = "string"
      }
    }
    transit_peer_network_name = "string"
    transit_peer_network_type = "string"
  }
}

output "catalystcenter_transit_peer_network_example" {
  value = catalystcenter_transit_peer_network.example
}