
resource "catalystcenter_network_devices_delete_with_cleanup" "example" {
    provider = meraki
    parameters = [{
      
      id = "string"
    }]
}

output "catalystcenter_network_devices_delete_with_cleanup_example" {
    value = catalystcenter_network_devices_delete_with_cleanup.example
}