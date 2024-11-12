
resource "catalystcenter_pnp_virtual_account_deregister" "example" {
  provider = meraki
  domain   = "string"
  name     = "string"
}

output "catalystcenter_pnp_virtual_account_deregister_example" {
  value = catalystcenter_pnp_virtual_account_deregister.example
}