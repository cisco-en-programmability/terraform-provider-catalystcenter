
data "catalystcenter_wireless_profile" "example" {
  provider     = catalystcenter
  profile_name = "string"
}

output "catalystcenter_wireless_profile_example" {
  value = data.catalystcenter_wireless_profile.example.items
}
