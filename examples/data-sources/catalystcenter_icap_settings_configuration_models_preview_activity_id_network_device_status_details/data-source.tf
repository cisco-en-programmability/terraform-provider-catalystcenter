
data "catalystcenter_icap_settings_configuration_models_preview_activity_id_network_device_status_details" "example" {
    provider = catalystcenter
    preview_activity_id = "string"
}

output "catalystcenter_icap_settings_configuration_models_preview_activity_id_network_device_status_details_example" {
    value = data.catalystcenter_icap_settings_configuration_models_preview_activity_id_network_device_status_details.example.items
}
