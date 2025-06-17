
resource "catalystcenter_image_distribution" "example" {
  provider = catalystcenter
  parameters = [{

    device_uuid = "string"
    image_uuid  = "string"
  }]
}

output "catalystcenter_image_distribution_example" {
  value = catalystcenter_image_distribution.example
}