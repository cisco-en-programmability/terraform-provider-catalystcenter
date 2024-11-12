
resource "catalystcenter_interface_operation_create" "example" {
  provider        = meraki
  deployment_mode = "string"
  interface_uuid  = "string"
  parameters {

    operation = "string"
    payload   = "------"
  }
}

output "catalystcenter_interface_operation_create_example" {
  value = catalystcenter_interface_operation_create.example
}