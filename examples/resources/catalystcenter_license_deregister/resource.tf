
resource "catalystcenter_license_deregister" "example" {
    provider = meraki
}

output "catalystcenter_license_deregister_example" {
    value = catalystcenter_license_deregister.example
}