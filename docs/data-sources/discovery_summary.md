---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_discovery_summary Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Discovery.
  Returns the devices discovered in the given discovery based on given filters. Discovery ID can be obtained using the
  "Get Discoveries by range" API.
---

# catalystcenter_discovery_summary (Data Source)

It performs read operation on Discovery.

- Returns the devices discovered in the given discovery based on given filters. Discovery ID can be obtained using the
"Get Discoveries by range" API.

## Example Usage

```terraform
data "catalystcenter_discovery_summary" "example" {
  provider       = catalystcenter
  clistatus      = ["string"]
  http_status    = ["string"]
  id             = "string"
  ip_address     = ["string"]
  netconf_status = ["string"]
  ping_status    = ["string"]
  snmp_status    = ["string"]
  sort_by        = "string"
  sort_order     = "string"
  task_id        = "string"
}

output "catalystcenter_discovery_summary_example" {
  value = data.catalystcenter_discovery_summary.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. Discovery ID

### Optional

- `clistatus` (List of String) cliStatus query parameter. CLI status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
- `http_status` (List of String) httpStatus query parameter. HTTP staus for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
- `ip_address` (List of String) ipAddress query parameter. IP Address of the device
- `netconf_status` (List of String) netconfStatus query parameter. NETCONF status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
- `ping_status` (List of String) pingStatus query parameter. Ping status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
- `snmp_status` (List of String) snmpStatus query parameter. SNMP status for the IP during the job run. Available values are 'SUCCESS', 'FAILURE', 'NOT-PROVIDED' and 'NOT-VALIDATED'
- `sort_by` (String) sortBy query parameter. Sort by field. Available values are pingStatus, cliStatus,snmpStatus, httpStatus and netconfStatus
- `sort_order` (String) sortOrder query parameter. Order of sorting based on sortBy. Available values are 'asc' and 'des'
- `task_id` (String) taskId query parameter.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `response` (Number)
- `version` (String)