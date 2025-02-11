
data "catalystcenter_site_membership" "example" {
    provider = catalystcenter
    device_family = "string"
    limit = 1
    offset = 1
    serial_number = "string"
    site_id = "string"
}

output "catalystcenter_site_membership_example" {
    value = data.catalystcenter_site_membership.example.item
}
