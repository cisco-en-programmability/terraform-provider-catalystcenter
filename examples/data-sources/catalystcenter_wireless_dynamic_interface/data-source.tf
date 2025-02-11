
data "catalystcenter_wireless_dynamic_interface" "example" {
    provider = catalystcenter
    interface_name = "string"
}

output "catalystcenter_wireless_dynamic_interface_example" {
    value = data.catalystcenter_wireless_dynamic_interface.example.items
}
