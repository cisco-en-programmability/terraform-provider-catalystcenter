
data "catalystcenter_floors_settings_v2" "example" {
    provider = catalystcenter
}

output "catalystcenter_floors_settings_v2_example" {
    value = data.catalystcenter_floors_settings_v2.example.item
}
