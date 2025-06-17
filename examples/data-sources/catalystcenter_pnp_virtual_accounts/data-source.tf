
data "catalystcenter_pnp_virtual_accounts" "example" {
  provider = catalystcenter
  domain   = "string"
}

output "catalystcenter_pnp_virtual_accounts_example" {
  value = data.catalystcenter_pnp_virtual_accounts.example.items
}
