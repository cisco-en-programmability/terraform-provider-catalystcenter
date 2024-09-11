
data "catalystcenter_floors" "example" {
  provider         = catalystcenter
  id               = "string"
  units_of_measure = "string"
}

output "catalystcenter_floors_example" {
  value = data.catalystcenter_floors.example.item
}
