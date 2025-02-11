
data "catalystcenter_platform_nodes_configuration_summary" "example" {
    provider = catalystcenter
}

output "catalystcenter_platform_nodes_configuration_summary_example" {
    value = data.catalystcenter_platform_nodes_configuration_summary.example.item
}
