
data "catalystcenter_security_threats_rogue_allowed_list" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
}

output "catalystcenter_security_threats_rogue_allowed_list_example" {
    value = data.catalystcenter_security_threats_rogue_allowed_list.example.items
}
