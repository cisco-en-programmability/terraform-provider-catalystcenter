---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_event_subscription_email Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read and update operations on Event Management.
  Create Email Subscription Endpoint for list of registered events.Update Email Subscription Endpoint for list of registered events
---

# catalystcenter_event_subscription_email (Resource)

It manages create, read and update operations on Event Management.

- Create Email Subscription Endpoint for list of registered events.

- Update Email Subscription Endpoint for list of registered events

## Example Usage

```terraform
resource "catalystcenter_event_subscription_email" "example" {
  provider = catalystcenter

  parameters {
    payload {
      description = "string"
      filter {

        categories = ["string"]
        domains_subdomains {

          domain      = "string"
          sub_domains = ["string"]
        }
        event_ids  = ["string"]
        severities = [1]
        site_ids   = ["string"]
        sources    = ["string"]
        types      = ["string"]
      }
      name = "string"
      subscription_endpoints {

        instance_id = "string"
        subscription_details {

          connector_type     = "string"
          description        = "string"
          from_email_address = "string"
          name               = "string"
          subject            = "string"
          to_email_addresses = ["string"]
        }
      }
      subscription_id = "string"
      version         = "string"
    }
  }
}

output "catalystcenter_event_subscription_email_example" {
  value = catalystcenter_event_subscription_email.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) Array of RequestEventManagementCreateEmailEventSubscription (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `payload` (Block List) Array of RequestEventManagementCreateEmailEventSubscription (see [below for nested schema](#nestedblock--parameters--payload))

<a id="nestedblock--parameters--payload"></a>
### Nested Schema for `parameters.payload`

Optional:

- `description` (String) Description
- `filter` (Block List) (see [below for nested schema](#nestedblock--parameters--payload--filter))
- `name` (String) Name
- `subscription_endpoints` (Block List) (see [below for nested schema](#nestedblock--parameters--payload--subscription_endpoints))
- `subscription_id` (String) Subscription Id (Unique UUID)
- `version` (String) Version

<a id="nestedblock--parameters--payload--filter"></a>
### Nested Schema for `parameters.payload.filter`

Optional:

- `categories` (List of String) Categories
- `domains_subdomains` (Block List) (see [below for nested schema](#nestedblock--parameters--payload--filter--domains_subdomains))
- `event_ids` (List of String) Event Ids
- `severities` (List of Number) Severities
- `site_ids` (List of String) Site Ids
- `sources` (List of String) Sources
- `types` (List of String) Types

<a id="nestedblock--parameters--payload--filter--domains_subdomains"></a>
### Nested Schema for `parameters.payload.filter.domains_subdomains`

Optional:

- `domain` (String) Domain
- `sub_domains` (List of String) Sub Domains



<a id="nestedblock--parameters--payload--subscription_endpoints"></a>
### Nested Schema for `parameters.payload.subscription_endpoints`

Optional:

- `instance_id` (String) (From Get Email Subscription Details --> pick InstanceId if available)
- `subscription_details` (Block List) (see [below for nested schema](#nestedblock--parameters--payload--subscription_endpoints--subscription_details))

<a id="nestedblock--parameters--payload--subscription_endpoints--subscription_details"></a>
### Nested Schema for `parameters.payload.subscription_endpoints.subscription_details`

Optional:

- `connector_type` (String) Connector Type (Must be EMAIL)
- `description` (String) Description
- `from_email_address` (String) Senders Email Address
- `name` (String) Name
- `subject` (String) Email Subject
- `to_email_addresses` (List of String) Recipient's Email Addresses (Comma separated)





<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `description` (String)
- `filter` (List of Object) (see [below for nested schema](#nestedobjatt--item--filter))
- `is_private` (String)
- `name` (String)
- `subscription_endpoints` (List of Object) (see [below for nested schema](#nestedobjatt--item--subscription_endpoints))
- `subscription_id` (String)
- `tenant_id` (String)
- `version` (String)

<a id="nestedobjatt--item--filter"></a>
### Nested Schema for `item.filter`

Read-Only:

- `categories` (List of String)
- `domains_subdomains` (List of Object) (see [below for nested schema](#nestedobjatt--item--filter--domains_subdomains))
- `event_ids` (List of String)
- `others` (List of String)
- `severities` (List of String)
- `site_ids` (List of String)
- `sources` (List of String)
- `types` (List of String)

<a id="nestedobjatt--item--filter--domains_subdomains"></a>
### Nested Schema for `item.filter.domains_subdomains`

Read-Only:

- `domain` (String)
- `sub_domains` (List of String)



<a id="nestedobjatt--item--subscription_endpoints"></a>
### Nested Schema for `item.subscription_endpoints`

Read-Only:

- `connector_type` (String)
- `instance_id` (String)
- `subscription_details` (List of Object) (see [below for nested schema](#nestedobjatt--item--subscription_endpoints--subscription_details))

<a id="nestedobjatt--item--subscription_endpoints--subscription_details"></a>
### Nested Schema for `item.subscription_endpoints.subscription_details`

Read-Only:

- `connector_type` (String)
- `description` (String)
- `from_email_address` (String)
- `instance_id` (String)
- `name` (String)
- `subject` (String)
- `to_email_addresses` (List of String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_event_subscription_email.example "id:=string"
```