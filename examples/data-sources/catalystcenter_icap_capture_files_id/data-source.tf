
data "catalystcenter_icap_capture_files_id" "example" {
  provider    = catalystcenter
  id          = "string"
  xca_lle_rid = "string"
}

output "catalystcenter_icap_capture_files_id_example" {
  value = data.catalystcenter_icap_capture_files_id.example.item
}
