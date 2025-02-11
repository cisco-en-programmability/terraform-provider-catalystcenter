
resource "catalystcenter_network_device_config_files_id_download_unmasked" "example" {
    provider = meraki
    id = "string"
    parameters = [{
      
      password = "******"
    }]
}

output "catalystcenter_network_device_config_files_id_download_unmasked_example" {
    value = catalystcenter_network_device_config_files_id_download_unmasked.example
}