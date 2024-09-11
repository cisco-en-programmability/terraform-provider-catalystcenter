
data "catalystcenter_service_provider_v2" "example" {
  provider = catalystcenter
}

output "catalystcenter_service_provider_v2_example" {
  value = data.catalystcenter_service_provider_v2.example.items
}
