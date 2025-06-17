
resource "catalystcenter_global_credential_delete" "example" {
  provider             = catalystcenter
  global_credential_id = "string"
  parameters = [{

  }]
}

output "catalystcenter_global_credential_delete_example" {
  value = catalystcenter_global_credential_delete.example
}