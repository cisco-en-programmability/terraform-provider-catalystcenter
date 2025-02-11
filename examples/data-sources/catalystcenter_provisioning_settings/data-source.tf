
data "catalystcenter_provisioning_settings" "example" {
    provider = catalystcenter
}

output "catalystcenter_provisioning_settings_example" {
    value = data.catalystcenter_provisioning_settings.example.item
}
