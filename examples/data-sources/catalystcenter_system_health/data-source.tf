
data "catalystcenter_system_health" "example" {
    provider = catalystcenter
    domain = "string"
    limit = 1
    offset = 1
    subdomain = "string"
    summary = "false"
}

output "catalystcenter_system_health_example" {
    value = data.catalystcenter_system_health.example.item
}
