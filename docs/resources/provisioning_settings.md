---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_provisioning_settings Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages read and update operations on System Settings.
  Sets provisioning settings
---

# catalystcenter_provisioning_settings (Resource)

It manages read and update operations on System Settings.

- Sets provisioning settings

## Example Usage

```terraform
resource "catalystcenter_provisioning_settings" "example" {
  provider = catalystcenter

  parameters {

    require_itsm_approval = "false"
    require_preview       = "false"
  }
}

output "catalystcenter_provisioning_settings_example" {
  value = catalystcenter_provisioning_settings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `require_itsm_approval` (String) If require ITSM approval is enabled, the planned configurations must be submitted for ITSM approval. Also if enabled, requirePreview will default to enabled.
- `require_preview` (String) If require preview is enabled, the device configurations must be reviewed before deploying them


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `require_itsm_approval` (String)
- `require_preview` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_provisioning_settings.example "id:=string"
```