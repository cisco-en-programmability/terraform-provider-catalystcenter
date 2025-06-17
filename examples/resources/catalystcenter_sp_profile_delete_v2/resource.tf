
resource "catalystcenter_sp_profile_delete_v2" "example" {
  provider        = catalystcenter
  sp_profile_name = "string"
  parameters = [{

  }]
}

output "catalystcenter_sp_profile_delete_v2_example" {
  value = catalystcenter_sp_profile_delete_v2.example
}