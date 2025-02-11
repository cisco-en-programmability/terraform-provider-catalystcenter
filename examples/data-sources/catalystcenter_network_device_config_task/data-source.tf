
data "catalystcenter_network_device_config_task" "example" {
    provider = catalystcenter
    parent_task_id = "string"
}

output "catalystcenter_network_device_config_task_example" {
    value = data.catalystcenter_network_device_config_task.example.items
}
