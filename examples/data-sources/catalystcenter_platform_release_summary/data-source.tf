
data "catalystcenter_platform_release_summary" "example" {
  provider = catalystcenter
}

output "catalystcenter_platform_release_summary_example" {
  value = data.catalystcenter_platform_release_summary.example.item
}
