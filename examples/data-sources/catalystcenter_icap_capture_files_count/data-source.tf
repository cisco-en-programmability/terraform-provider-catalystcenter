
data "catalystcenter_icap_capture_files_count" "example" {
  provider    = catalystcenter
  ap_mac      = "string"
  client_mac  = "string"
  end_time    = 1609459200
  start_time  = 1609459200
  type        = "string"
  xca_lle_rid = "string"
}

output "catalystcenter_icap_capture_files_count_example" {
  value = data.catalystcenter_icap_capture_files_count.example.item
}
