
resource "catalystcenter_interface_operation_create" "example" {
  provider        = catalystcenter
  deployment_mode = "string"
  interface_uuid  = "string"
  parameters {

    operation = "string"
    payload   = "string"
  }
}

output "catalystcenter_interface_operation_create_example" {
  value = catalystcenter_interface_operation_create.example
}