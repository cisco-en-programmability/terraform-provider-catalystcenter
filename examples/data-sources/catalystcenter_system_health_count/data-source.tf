
data "catalystcenter_system_health_count" "example" {
    provider = catalystcenter
    domain = "string"
    subdomain = "string"
}

output "catalystcenter_system_health_count_example" {
    value = data.catalystcenter_system_health_count.example.item
}
