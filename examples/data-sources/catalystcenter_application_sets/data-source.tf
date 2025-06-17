
data "catalystcenter_application_sets" "example" {
  provider = catalystcenter
  limit    = 1
  name     = "string"
  offset   = 1
}

output "catalystcenter_application_sets_example" {
  value = data.catalystcenter_application_sets.example.items
}
