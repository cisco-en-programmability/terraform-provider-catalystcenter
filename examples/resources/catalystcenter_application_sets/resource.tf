
resource "catalystcenter_application_sets" "example" {
  provider = catalystcenter

  parameters {

    name = "string"
  }
}

output "catalystcenter_application_sets_example" {
  value = catalystcenter_application_sets.example
}