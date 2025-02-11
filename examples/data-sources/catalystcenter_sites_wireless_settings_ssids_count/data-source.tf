
data "catalystcenter_sites_wireless_settings_ssids_count" "example" {
    provider = catalystcenter
    inherited = "false"
    site_id = "string"
}

output "catalystcenter_sites_wireless_settings_ssids_count_example" {
    value = data.catalystcenter_sites_wireless_settings_ssids_count.example.item
}
