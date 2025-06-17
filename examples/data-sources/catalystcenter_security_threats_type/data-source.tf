
data "catalystcenter_security_threats_type" "example" {
  provider = catalystcenter
}

output "catalystcenter_security_threats_type_example" {
  value = data.catalystcenter_security_threats_type.example.items
}
