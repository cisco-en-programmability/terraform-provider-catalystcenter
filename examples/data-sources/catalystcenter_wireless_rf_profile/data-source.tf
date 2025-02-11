
data "catalystcenter_wireless_rf_profile" "example" {
    provider = catalystcenter
    rf_profile_name = "string"
}

output "catalystcenter_wireless_rf_profile_example" {
    value = data.catalystcenter_wireless_rf_profile.example.item
}
