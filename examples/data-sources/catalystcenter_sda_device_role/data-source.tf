
data "catalystcenter_sda_device_role" "example" {
    provider = catalystcenter
    device_management_ip_address = "string"
}

output "catalystcenter_sda_device_role_example" {
    value = data.catalystcenter_sda_device_role.example.item
}
