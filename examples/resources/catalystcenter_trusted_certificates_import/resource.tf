
resource "catalystcenter_trusted_certificates_import" "example" {
  provider = catalystcenter
}

output "catalystcenter_trusted_certificates_import_example" {
  value = catalystcenter_trusted_certificates_import.example
}