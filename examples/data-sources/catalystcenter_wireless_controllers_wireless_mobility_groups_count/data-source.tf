
data "catalystcenter_wireless_controllers_wireless_mobility_groups_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_wireless_controllers_wireless_mobility_groups_count_example" {
    value = data.catalystcenter_wireless_controllers_wireless_mobility_groups_count.example.item
}
