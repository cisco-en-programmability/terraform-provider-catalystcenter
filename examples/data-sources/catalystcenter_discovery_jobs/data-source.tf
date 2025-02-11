
data "catalystcenter_discovery_jobs" "example" {
    provider = catalystcenter
    ip_address = "string"
    limit = 1
    name = "string"
    offset = 1
}

output "catalystcenter_discovery_jobs_example" {
    value = data.catalystcenter_discovery_jobs.example.items
}
