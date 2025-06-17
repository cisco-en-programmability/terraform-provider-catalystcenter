
data "catalystcenter_roles" "example" {
  provider      = catalystcenter
  invoke_source = "string"
}

output "catalystcenter_roles_example" {
  value = data.catalystcenter_roles.example.item
}
