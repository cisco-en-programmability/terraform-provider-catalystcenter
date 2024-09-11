
resource "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id" "example" {
  provider = catalystcenter

  parameters {

    fabric_id = "string"
    ssid_details {

      name               = "string"
      security_group_tag = "string"
    }
    vlan_name = "string"
  }
}

output "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_example" {
  value = catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id.example
}