
resource "catalystcenter_application_visibility_network_devices_enable_cbar" "example" {
    provider = meraki
    parameters = [{
      
      network_devices = [{
        
        exclude_interface_ids = ["string"]
        exclude_wlan_modes = ["string"]
        id = "string"
      }]
    }]
}

output "catalystcenter_application_visibility_network_devices_enable_cbar_example" {
    value = catalystcenter_application_visibility_network_devices_enable_cbar.example
}