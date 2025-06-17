
data "catalystcenter_topology_site" "example" {
  provider = catalystcenter
}

output "catalystcenter_topology_site_example" {
  value = data.catalystcenter_topology_site.example.item
}
