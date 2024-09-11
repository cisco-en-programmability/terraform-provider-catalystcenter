
data "catalystcenter_client_proximity" "example" {
  provider        = catalystcenter
  number_days     = 1
  time_resolution = 1.0
  username        = "string"
}

output "catalystcenter_client_proximity_example" {
  value = data.catalystcenter_client_proximity.example.item
}
