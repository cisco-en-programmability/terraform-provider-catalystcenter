
data "catalystcenter_device_reboot_apreboot" "example" {
    provider = catalystcenter
    parent_task_id = "string"
}

output "catalystcenter_device_reboot_apreboot_example" {
    value = data.catalystcenter_device_reboot_apreboot.example.items
}
