
data "catalystcenter_discovery_job_by_id" "example" {
    provider = catalystcenter
    id = "string"
    ip_address = "string"
    limit = 1
    offset = 1
}

output "catalystcenter_discovery_job_by_id_example" {
    value = data.catalystcenter_discovery_job_by_id.example.items
}
