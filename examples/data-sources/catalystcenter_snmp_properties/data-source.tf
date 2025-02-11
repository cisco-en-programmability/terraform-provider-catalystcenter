
data "catalystcenter_snmp_properties" "example" {
    provider = catalystcenter
}

output "catalystcenter_snmp_properties_example" {
    value = data.catalystcenter_snmp_properties.example.items
}
