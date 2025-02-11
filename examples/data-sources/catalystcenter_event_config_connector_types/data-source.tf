
data "catalystcenter_event_config_connector_types" "example" {
    provider = catalystcenter
}

output "catalystcenter_event_config_connector_types_example" {
    value = data.catalystcenter_event_config_connector_types.example.items
}
