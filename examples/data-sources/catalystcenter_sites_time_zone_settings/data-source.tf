
data "catalystcenter_sites_time_zone_settings" "example" {
    provider = catalystcenter
    id = "string"
    inherited = "false"
}

output "catalystcenter_sites_time_zone_settings_example" {
    value = data.catalystcenter_sites_time_zone_settings.example.item
}
