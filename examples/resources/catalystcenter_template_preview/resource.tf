
resource "catalystcenter_template_preview" "example" {
    provider = meraki
    parameters = [{
      
      device_id = "string"
      params = ------
      resource_params = ------
      template_id = "string"
    }]
}

output "catalystcenter_template_preview_example" {
    value = catalystcenter_template_preview.example
}