---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_security_threats_type Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Intent API to fetch all threat types defined.
---

# catalystcenter_security_threats_type (Data Source)

It performs read operation on Devices.

- Intent API to fetch all threat types defined.

## Example Usage

```terraform
data "catalystcenter_security_threats_type" "example" {
  provider = catalystcenter
}

output "catalystcenter_security_threats_type_example" {
  value = data.catalystcenter_security_threats_type.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `is_custom` (String)
- `is_deleted` (String)
- `label` (String)
- `name` (String)
- `value` (Number)
