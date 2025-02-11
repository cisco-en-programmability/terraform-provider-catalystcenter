
data "catalystcenter_wireless_controllers_mesh_ap_neighbours_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_wireless_controllers_mesh_ap_neighbours_count_example" {
    value = data.catalystcenter_wireless_controllers_mesh_ap_neighbours_count.example.item
}
