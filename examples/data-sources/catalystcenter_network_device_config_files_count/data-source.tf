
data "catalystcenter_network_device_config_files_count" "example" {
    provider = catalystcenter
    file_type = "string"
    id = "string"
    network_device_id = "string"
}

output "catalystcenter_network_device_config_files_count_example" {
    value = data.catalystcenter_network_device_config_files_count.example.item
}
