
data "catalystcenter_security_threats_level" "example" {
    provider = catalystcenter
}

output "catalystcenter_security_threats_level_example" {
    value = data.catalystcenter_security_threats_level.example.items
}
