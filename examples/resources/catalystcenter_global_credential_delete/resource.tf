
resource "catalystcenter_global_credential_delete" "example" {
  provider             = meraki
  global_credential_id = "string"
  parameters {

  }
}

output "catalystcenter_global_credential_delete_example" {
  value = catalystcenter_global_credential_delete.example
}