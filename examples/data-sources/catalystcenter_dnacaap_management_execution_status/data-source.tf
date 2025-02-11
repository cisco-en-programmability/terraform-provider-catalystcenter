
data "catalystcenter_dnacaap_management_execution_status" "example" {
    provider = catalystcenter
    execution_id = "string"
}

output "catalystcenter_dnacaap_management_execution_status_example" {
    value = data.catalystcenter_dnacaap_management_execution_status.example.item
}
