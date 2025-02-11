
resource "catalystcenter_network_devices_delete_without_cleanup" "example" {
    provider = meraki
    parameters = [{
      
      id = "string"
    }]
}

output "catalystcenter_network_devices_delete_without_cleanup_example" {
    value = catalystcenter_network_devices_delete_without_cleanup.example
}