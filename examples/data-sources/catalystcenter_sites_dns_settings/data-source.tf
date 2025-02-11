
data "catalystcenter_sites_dns_settings" "example" {
    provider = catalystcenter
    id = "string"
    inherited = "false"
}

output "catalystcenter_sites_dns_settings_example" {
    value = data.catalystcenter_sites_dns_settings.example.item
}
