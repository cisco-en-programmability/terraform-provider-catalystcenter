
data "catalystcenter_pnp_workflow" "example" {
  provider   = catalystcenter
  limit      = 1
  name       = ["string"]
  offset     = 1
  sort       = ["string"]
  sort_order = "string"
  type       = ["string"]
}

output "catalystcenter_pnp_workflow_example" {
  value = data.catalystcenter_pnp_workflow.example.items
}

data "catalystcenter_pnp_workflow" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_pnp_workflow_example" {
  value = data.catalystcenter_pnp_workflow.example.item
}
