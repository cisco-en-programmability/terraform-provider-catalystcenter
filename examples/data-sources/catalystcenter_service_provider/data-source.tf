
data "catalystcenter_service_provider" "example" {
  provider = catalystcenter
}

output "catalystcenter_service_provider_example" {
  value = data.catalystcenter_service_provider.example.items
}
