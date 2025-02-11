
data "catalystcenter_wireless_profiles" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
    wireless_profile_name = "string"
}

output "catalystcenter_wireless_profiles_example" {
    value = data.catalystcenter_wireless_profiles.example.items
}

data "catalystcenter_wireless_profiles" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_wireless_profiles_example" {
    value = data.catalystcenter_wireless_profiles.example.item
}
