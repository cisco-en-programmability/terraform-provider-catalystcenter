
data "catalystcenter_wireless_settings_ssids_override_at_sites" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
    site_id = "string"
}

output "catalystcenter_wireless_settings_ssids_override_at_sites_example" {
    value = data.catalystcenter_wireless_settings_ssids_override_at_sites.example.items
}
