
data "catalystcenter_pnp_smart_account_domains" "example" {
  provider = catalystcenter
}

output "catalystcenter_pnp_smart_account_domains_example" {
  value = data.catalystcenter_pnp_smart_account_domains.example.items
}
