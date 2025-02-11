
data "catalystcenter_discovery" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_discovery_example" {
    value = data.catalystcenter_discovery.example.item
}
