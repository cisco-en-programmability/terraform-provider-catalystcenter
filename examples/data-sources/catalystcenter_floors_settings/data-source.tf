
data "catalystcenter_floors_settings" "example" {
  provider = catalystcenter
}

output "catalystcenter_floors_settings_example" {
  value = data.catalystcenter_floors_settings.example.item
}
