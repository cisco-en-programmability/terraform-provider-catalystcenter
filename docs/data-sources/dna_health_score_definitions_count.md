---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_dna_health_score_definitions_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Get the count of health score definitions based on provided filters. Supported filters are id, name and overall health
  include status. For detailed information about the usage of the API, please refer to the Open API specification document
  https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CECatCenter_Org-
  issueAndHealthDefinitions-1.0.0-resolved.yaml
---

# catalystcenter_dna_health_score_definitions_count (Data Source)

It performs read operation on Devices.

- Get the count of health score definitions based on provided filters. Supported filters are id, name and overall health
include status. For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
issueAndHealthDefinitions-1.0.0-resolved.yaml

## Example Usage

```terraform
data "catalystcenter_dna_health_score_definitions_count" "example" {
  provider                   = catalystcenter
  device_type                = "string"
  id                         = "string"
  include_for_overall_health = "false"
  xca_lle_rid                = "string"
}

output "catalystcenter_dna_health_score_definitions_count_example" {
  value = data.catalystcenter_dna_health_score_definitions_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.

### Optional

- `device_type` (String) deviceType query parameter. These are the device families supported for health score definitions. If no input is made on device family, all device families are considered.
- `id` (String) id query parameter. The definition identifier.
Examples:
id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request)
id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
- `include_for_overall_health` (Boolean) includeForOverallHealth query parameter. The inclusion status of the issue definition, either true or false. true indicates that particular health metric is included in overall health computation, otherwise false. By default it's set to true.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)
