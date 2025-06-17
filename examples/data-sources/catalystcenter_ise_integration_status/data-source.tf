
data "catalystcenter_ise_integration_status" "example" {
  provider = catalystcenter
}

output "catalystcenter_ise_integration_status_example" {
  value = data.catalystcenter_ise_integration_status.example.item
}
