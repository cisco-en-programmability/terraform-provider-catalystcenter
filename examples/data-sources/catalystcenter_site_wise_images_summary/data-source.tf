
data "catalystcenter_site_wise_images_summary" "example" {
  provider = catalystcenter
  site_id  = "string"
}

output "catalystcenter_site_wise_images_summary_example" {
  value = data.catalystcenter_site_wise_images_summary.example.item
}
