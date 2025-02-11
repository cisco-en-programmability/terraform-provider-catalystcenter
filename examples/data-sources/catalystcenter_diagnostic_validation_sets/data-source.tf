
data "catalystcenter_diagnostic_validation_sets" "example" {
    provider = catalystcenter
    view = "string"
}

output "catalystcenter_diagnostic_validation_sets_example" {
    value = data.catalystcenter_diagnostic_validation_sets.example.items
}

data "catalystcenter_diagnostic_validation_sets" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_diagnostic_validation_sets_example" {
    value = data.catalystcenter_diagnostic_validation_sets.example.item
}
