
resource "catalystcenter_pnp_virtual_account_deregister" "example" {
  provider = catalystcenter
  domain   = "string"
  name     = "string"
}

output "catalystcenter_pnp_virtual_account_deregister_example" {
  value = catalystcenter_pnp_virtual_account_deregister.example
}