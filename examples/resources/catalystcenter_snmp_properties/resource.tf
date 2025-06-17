
resource "catalystcenter_snmp_properties" "example" {
  provider = catalystcenter
 
  parameters {

    id                   = "string"
    instance_tenant_id   = "string"
    instance_uuid        = "string"
    int_value            = 1
    system_property_name = "string"
  }
}

output "catalystcenter_snmp_properties_example" {
  value = catalystcenter_snmp_properties.example
}