
data "catalystcenter_security_advisories" "example" {
    provider = catalystcenter
}

output "catalystcenter_security_advisories_example" {
    value = data.catalystcenter_security_advisories.example.items
}
