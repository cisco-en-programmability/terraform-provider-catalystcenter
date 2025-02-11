
data "catalystcenter_itsm_cmdb_sync_status" "example" {
    provider = catalystcenter
    date = "string"
    status = "string"
}

output "catalystcenter_itsm_cmdb_sync_status_example" {
    value = data.catalystcenter_itsm_cmdb_sync_status.example.items
}
