
data "catalystcenter_ipam_server_setting" "example" {
  provider = catalystcenter
}

output "catalystcenter_ipam_server_setting_example" {
  value = data.catalystcenter_ipam_server_setting.example.item
}
