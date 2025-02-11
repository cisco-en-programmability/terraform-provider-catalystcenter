
data "catalystcenter_sensor" "example" {
    provider = catalystcenter
    site_id = "string"
}

output "catalystcenter_sensor_example" {
    value = data.catalystcenter_sensor.example.items
}
