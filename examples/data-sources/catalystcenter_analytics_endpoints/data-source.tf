
data "catalystcenter_analytics_endpoints" "example" {
  provider = catalystcenter
  ep_id    = "string"
  include  = "string"
}

output "catalystcenter_analytics_endpoints_example" {
  value = data.catalystcenter_analytics_endpoints.example.item
}
