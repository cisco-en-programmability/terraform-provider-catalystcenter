
data "catalystcenter_pnp_workflow_count" "example" {
  provider = catalystcenter
  name     = ["string"]
}

output "catalystcenter_pnp_workflow_count_example" {
  value = data.catalystcenter_pnp_workflow_count.example.item
}
