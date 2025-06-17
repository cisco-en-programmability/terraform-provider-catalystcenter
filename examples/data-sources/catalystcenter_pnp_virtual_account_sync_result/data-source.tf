
data "catalystcenter_pnp_virtual_account_sync_result" "example" {
  provider = catalystcenter
  domain   = "string"
  name     = "string"
}

output "catalystcenter_pnp_virtual_account_sync_result_example" {
  value = data.catalystcenter_pnp_virtual_account_sync_result.example.item
}
