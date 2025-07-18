terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_sda_fabric_border_device" "zh_fiab_1" {
  provider = catalystcenter
  #   depends_on = [ catalystcenter_sda_provision_device.zh_fiab_1, catalystcenter_transit_peer_network.Customer_C_Transit ]
  parameters {
    payload {
      device_management_ip_address       = "10.121.1.1"
      device_role                        = ["Border_Node", "Control_Plane_Node", "Edge_Node"]
      border_session_type                = "EXTERNAL"
      border_with_external_connectivity  = "true"
      connected_to_internet              = "false"
      external_connectivity_ip_pool_name = "Customer_C_ZH-Transit-Pool"
      external_connectivity_settings {

        external_autonomou_system_number = "65535"
        interface_name                   = "TenGigabitEthernet1/1/1"
        interface_description            = "Uplink"
        l3_handoff {
          virtual_network {

            virtual_network_name = "C_Guest"
            vlan_id              = "3003"
          }
        }
        l3_handoff {
          virtual_network {

            virtual_network_name = "INFRA_VN"
            vlan_id              = "3001"
          }
        }
        l3_handoff {
          virtual_network {

            virtual_network_name = "C_Campus"
            vlan_id              = "3002"
          }
        }
      }


      external_domain_routing_protocol_name = "BGP"
      internal_autonomou_system_number      = "65534"
      site_name_hierarchy                   = "Global/San Francisco"
    }
  }
}
