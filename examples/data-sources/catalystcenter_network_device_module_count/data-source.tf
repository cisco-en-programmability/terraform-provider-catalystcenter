
data "catalystcenter_network_device_module_count" "example" {
  provider                    = catalystcenter
  device_id                   = "string"
  name_list                   = ["string"]
  operational_state_code_list = ["string"]
  part_number_list            = ["string"]
  vendor_equipment_type_list  = ["string"]
}

output "catalystcenter_network_device_module_count_example" {
  value = data.catalystcenter_network_device_module_count.example.item
}
