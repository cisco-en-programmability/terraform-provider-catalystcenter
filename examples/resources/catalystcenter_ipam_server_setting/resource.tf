
resource "catalystcenter_ipam_server_setting" "example" {
  provider = catalystcenter

  parameters {

    password    = "******"
    provider    = "string"
    server_name = "string"
    server_url  = "string"
    sync_view   = "false"
    user_name   = "string"
    view        = "string"
  }
}

output "catalystcenter_ipam_server_setting_example" {
  value = catalystcenter_ipam_server_setting.example
}