
data "catalystcenter_network_device_module" "example" {
    provider = catalystcenter
    device_id = "string"
    limit = 1
    name_list = ["string"]
    offset = 1
    operational_state_code_list = ["string"]
    part_number_list = ["string"]
    vendor_equipment_type_list = ["string"]
}

output "catalystcenter_network_device_module_example" {
    value = data.catalystcenter_network_device_module.example.items
}

data "catalystcenter_network_device_module" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_network_device_module_example" {
    value = data.catalystcenter_network_device_module.example.item
}
