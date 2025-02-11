
data "catalystcenter_license_setting" "example" {
    provider = catalystcenter
}

output "catalystcenter_license_setting_example" {
    value = data.catalystcenter_license_setting.example.item
}
