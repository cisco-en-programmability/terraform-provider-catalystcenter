
data "catalystcenter_sites_ntp_settings" "example" {
    provider = catalystcenter
    id = "string"
    inherited = "false"
}

output "catalystcenter_sites_ntp_settings_example" {
    value = data.catalystcenter_sites_ntp_settings.example.item
}
