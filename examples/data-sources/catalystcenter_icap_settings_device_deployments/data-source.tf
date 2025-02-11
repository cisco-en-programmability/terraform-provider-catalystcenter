
data "catalystcenter_icap_settings_device_deployments" "example" {
    provider = catalystcenter
    deploy_activity_id = "string"
    limit = 1
    network_device_ids = "string"
    offset = 1
    order = "string"
    sort_by = "string"
}

output "catalystcenter_icap_settings_device_deployments_example" {
    value = data.catalystcenter_icap_settings_device_deployments.example.items
}
