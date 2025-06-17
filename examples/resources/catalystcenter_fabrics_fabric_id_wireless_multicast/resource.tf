
resource "catalystcenter_fabrics_fabric_id_wireless_multicast" "example" {
  provider = catalystcenter
 
  parameters {

    fabric_id         = "string"
    multicast_enabled = "false"
  }
}

output "catalystcenter_fabrics_fabric_id_wireless_multicast_example" {
  value = catalystcenter_fabrics_fabric_id_wireless_multicast.example
}