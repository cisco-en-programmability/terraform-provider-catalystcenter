
data "catalystcenter_aaa_services_id" "example" {
    provider = catalystcenter
    end_time = 1609459200
    id = "string"
    start_time = 1609459200
    xca_lle_rid = "string"
}

output "catalystcenter_aaa_services_id_example" {
    value = data.catalystcenter_aaa_services_id.example.item
}
