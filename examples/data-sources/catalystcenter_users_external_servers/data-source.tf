
data "catalystcenter_users_external_servers" "example" {
  provider      = catalystcenter
  invoke_source = "string"
}

output "catalystcenter_users_external_servers_example" {
  value = data.catalystcenter_users_external_servers.example.item
}
