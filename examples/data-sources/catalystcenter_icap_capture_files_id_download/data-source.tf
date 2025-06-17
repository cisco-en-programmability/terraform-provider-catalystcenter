
data "catalystcenter_icap_capture_files_id_download" "example" {
  provider    = catalystcenter
  id          = "string"
  xca_lle_rid = "string"
}

output "catalystcenter_icap_capture_files_id_download_example" {
  value = data.catalystcenter_icap_capture_files_id_download.example.item
}
