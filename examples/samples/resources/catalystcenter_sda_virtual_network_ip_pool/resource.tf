terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # source = "cisco-en-programmability/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}
resource "catalystcenter_sda_virtual_network_ip_pool" "example" {
  provider = catalystcenter
  parameters {

    # auto_generate_vlan_name  = "false"
    ip_pool_name = "21Network"
    # is_common_pool           = "false"
    # is_ip_directed_broadcast = "false"
    # is_l2_flooding_enabled   = "false"
    # is_layer2_only           = "false"
    # is_this_critical_pool    = "false"
    is_wireless_pool = "false"
    # pool_type                = "string"
    # scalable_group_name      = "string"
    # traffic_type             = "string"
    site_name_hierarchy  = "Global/New Jersey/MurrayHill/test/TestFloor"
    virtual_network_name = "ANSIBLE800"
    # vlan_id                  = "string"
    # vlan_name                = "string"
  }
}

output "catalystcenter_sda_virtual_network_ip_pool_example" {
  value = catalystcenter_sda_virtual_network_ip_pool.example
}
