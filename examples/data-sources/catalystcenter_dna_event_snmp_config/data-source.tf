
data "catalystcenter_dna_event_snmp_config" "example" {
    provider = catalystcenter
    config_id = "string"
    limit = 1
    offset = 1
    order = "string"
    sort_by = "string"
}

output "catalystcenter_dna_event_snmp_config_example" {
    value = data.catalystcenter_dna_event_snmp_config.example.items
}
