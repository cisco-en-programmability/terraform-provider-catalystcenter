
data "catalystcenter_fabrics_fabric_id_wireless_multicast" "example" {
    provider = catalystcenter
    fabric_id = "string"
}

output "catalystcenter_fabrics_fabric_id_wireless_multicast_example" {
    value = data.catalystcenter_fabrics_fabric_id_wireless_multicast.example.item
}
