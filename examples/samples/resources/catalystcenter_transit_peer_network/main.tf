terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}
resource "catalystcenter_transit_peer_network" "example" {
  provider = catalystcenter
  parameters {
    ip_transit_settings {
      autonomous_system_number = "11"
      routing_protocol_name    = "BGP"
    }
    /*sda_transit_settings{
        transit_control_plane_settings{
            device_management_ip_address="string"
            site_name_hierarchy="string"
        }
    }*/
    transit_peer_network_name = "string"
    transit_peer_network_type = "ip_transit"
  }
}

output "catalystcenter_transit_peer_network_example" {
  value = catalystcenter_transit_peer_network.example
}
