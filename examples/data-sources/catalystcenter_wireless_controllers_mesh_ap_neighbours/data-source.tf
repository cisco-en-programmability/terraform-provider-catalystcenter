
data "catalystcenter_wireless_controllers_mesh_ap_neighbours" "example" {
    provider = catalystcenter
    ap_name = "string"
    ethernet_mac_address = "string"
    wlc_ip_address = "string"
}

output "catalystcenter_wireless_controllers_mesh_ap_neighbours_example" {
    value = data.catalystcenter_wireless_controllers_mesh_ap_neighbours.example.item
}
