
data "catalystcenter_sda_provision_devices_count" "example" {
    provider = catalystcenter
    site_id = "string"
}

output "catalystcenter_sda_provision_devices_count_example" {
    value = data.catalystcenter_sda_provision_devices_count.example.item
}
