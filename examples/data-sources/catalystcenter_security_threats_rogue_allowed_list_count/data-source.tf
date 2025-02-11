
data "catalystcenter_security_threats_rogue_allowed_list_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_security_threats_rogue_allowed_list_count_example" {
    value = data.catalystcenter_security_threats_rogue_allowed_list_count.example.item
}
