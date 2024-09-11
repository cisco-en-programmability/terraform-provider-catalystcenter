
data "catalystcenter_authentication_policy_servers" "example" {
  provider       = catalystcenter
  is_ise_enabled = "false"
  role           = "string"
  state          = "string"
}

output "catalystcenter_authentication_policy_servers_example" {
  value = data.catalystcenter_authentication_policy_servers.example.items
}
