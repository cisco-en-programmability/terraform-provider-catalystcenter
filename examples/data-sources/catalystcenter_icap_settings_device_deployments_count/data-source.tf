
data "catalystcenter_icap_settings_device_deployments_count" "example" {
    provider = catalystcenter
    deploy_activity_id = "string"
    network_device_ids = "string"
}

output "catalystcenter_icap_settings_device_deployments_count_example" {
    value = data.catalystcenter_icap_settings_device_deployments_count.example.item
}
