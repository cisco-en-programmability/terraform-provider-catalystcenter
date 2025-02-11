
resource "catalystcenter_wireless_controllers_provision" "example" {
    provider = meraki
    device_id = "string"
    parameters = [{
      
      ap_authorization_list_name = "string"
      authorize_mesh_and_non_mesh_access_points = "false"
      interfaces = [{
        
        interface_gateway = "string"
        interface_ipaddress = "string"
        interface_name = "string"
        interface_netmask_in_cid_r = 1
        lag_or_port_number = 1
        vlan_id = 1
      }]
      rolling_ap_upgrade = [{
        
        ap_reboot_percentage = 1
        enable_rolling_ap_upgrade = "false"
      }]
      skip_ap_provision = "false"
    }]
}

output "catalystcenter_wireless_controllers_provision_example" {
    value = catalystcenter_wireless_controllers_provision.example
}