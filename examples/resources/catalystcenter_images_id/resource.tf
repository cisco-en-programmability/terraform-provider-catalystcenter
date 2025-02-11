
resource "catalystcenter_images_id" "example" {
    provider = meraki
    id = "string"
    parameters = [{
      
    }]
}

output "catalystcenter_images_id_example" {
    value = catalystcenter_images_id.example
}