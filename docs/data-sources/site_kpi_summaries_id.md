---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_site_kpi_summaries_id Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Sites.
  Returns site analytics for the given site. For detailed information about the usage of the API, please refer to the
  Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
  specs/blob/main/Assurance/CECatCenter_Org-SiteKpiSummaries-1.0.0-resolved.yaml
---

# catalystcenter_site_kpi_summaries_id (Data Source)

It performs read operation on Sites.

- Returns site analytics for the given site. For detailed information about the usage of the API, please refer to the
Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml

## Example Usage

```terraform
data "catalystcenter_site_kpi_summaries_id" "example" {
  provider         = catalystcenter
  attribute        = "string"
  band             = "string"
  end_time         = 1609459200
  failure_category = "string"
  failure_reason   = "string"
  id               = "string"
  ssid             = "string"
  start_time       = 1609459200
  task_id          = "string"
  view             = "string"
  xca_lle_rid      = "string"
}

output "catalystcenter_site_kpi_summaries_id_example" {
  value = data.catalystcenter_site_kpi_summaries_id.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. The Site UUID
- `xca_lle_rid` (String) X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.

### Optional

- `attribute` (String) attribute query parameter. List of attributes related to site analytics. If these are provided, then only those attributes will be part of response along with the default attributes. Examples: **attribute=coverageAverage** (single attribute requested) **attribute=coverageFailureMetrics&attribute=coverageTotalCount** (multiple attributes requested)
- `band` (String) band query parameter. WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz GHz Examples: **band=5** (single band requested) **band=2.4&band=6** (multiple band requested)
- `end_time` (Number) endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
- `failure_category` (String) failureCategory query parameter. Category of failure when a client fails to meet the threshold. Examples: **failureCategory=AUTH** (single failure category requested) **failureCategory=AUTH&failureCategory=DHCP** (multiple failure categories requested)
- `failure_reason` (String) failureReason query parameter. Reason for failure when a client fails to meet the threshold. Examples: **failureReason=MOBILITY_FAILURE** (single ssid requested) **failureReason=REASON_IPLEARN_CONNECT_TIMEOUT&failureReason=ST_EAP_TIMEOUT**   (multiple ssid requested)
- `ssid` (String) ssid query parameter. SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID Wireless Local Area Network Identifier. Examples: **ssid=Alpha** (single ssid requested) **ssid=Alpha&ssid=Guest** (multiple ssid requested)
- `start_time` (Number) startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
- `task_id` (String) taskId query parameter. used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
- `view` (String) view query parameter. 
The name of the View. Each view represents a specific data set. Please refer to the 
SiteAnalyticsView
 Model for supported views. View is predefined set of attributes supported by the API. Only the attributes related to the given view will be part of the API response along with default attributes. If multiple views are provided, then response will contain attributes from all those views. If no views are specified, all attributes will be returned.
View Name
Included Attributes
coverage
coverageAverage, coverageSuccessPercentage, coverageSuccessCount, coverageTotalCount, coverageFailureCount, coverageClientCount, coverageImpactedEntities, coverageFailureImpactedEntities, coverageFailureMetrics
onboardingAttempts
onboardingAttemptsSuccessPercentage, onboardingAttemptsSuccessCount, onboardingAttemptsTotalCount, onboardingAttemptsFailureCount, onboardingAttemptsClientCount, onboardingAttemptsImpactedEntities, onboardingAttemptsFailureImpactedEntities, onboardingAttemptsFailureMetrics
onboardingDuration
onboardingDurationAverage, onboardingDurationSuccessPercentage, onboardingDurationSuccessCount, onboardingDurationTotalCount, onboardingDurationFailureCount, onboardingDurationClientCount, onboardingDurationImpactedEntities, onboardingDurationFailureImpactedEntities, onboardingDurationFailureMetrics
roamingAttempts
roamingAttemptsSuccessPercentage, roamingAttemptsSuccessCount, roamingAttemptsTotalCount, roamingAttemptsFailureCount, roamingAttemptsClientCount, roamingAttemptsImpactedEntities, roamingAttemptsFailureImpactedEntities, roamingAttemptsFailureMetrics
roamingDuration
roamingDurationAverage, roamingDurationSuccessPercentage, roamingDurationSuccessCount, roamingDurationTotalCount, roamingDurationFailureCount, roamingDurationClientCount, roamingDurationImpactedEntities, roamingDurationFailureImpactedEntities, roamingDurationFailureMetrics
connectionSpeed
connectionSpeedAverage, connectionSpeedSuccessPercentage, connectionSpeedSuccessCount, connectionSpeedTotalCount, connectionSpeedFailureCount, connectionSpeedClientCount, connectionSpeedImpactedEntities, connectionSpeedFailureImpactedEntities, connectionSpeedFailureMetrics
Examples: 
view=connectionSpeed
 (single view requested) 
view=roamingDuration&view=roamingAttempts
 (multiple views requested)

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `ap_count` (Number)
- `connection_speed_average` (Number)
- `connection_speed_client_count` (Number)
- `connection_speed_failure_count` (Number)
- `connection_speed_failure_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--connection_speed_failure_impacted_entities))
- `connection_speed_failure_metrics` (List of Object) (see [below for nested schema](#nestedobjatt--item--connection_speed_failure_metrics))
- `connection_speed_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--connection_speed_impacted_entities))
- `connection_speed_success_count` (Number)
- `connection_speed_success_percentage` (Number)
- `connection_speed_total_count` (Number)
- `coverage_average` (Number)
- `coverage_client_count` (Number)
- `coverage_failure_count` (Number)
- `coverage_failure_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--coverage_failure_impacted_entities))
- `coverage_failure_metrics` (List of Object) (see [below for nested schema](#nestedobjatt--item--coverage_failure_metrics))
- `coverage_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--coverage_impacted_entities))
- `coverage_success_count` (Number)
- `coverage_success_percentage` (Number)
- `coverage_total_count` (Number)
- `id` (String)
- `onboarding_attempts_client_count` (Number)
- `onboarding_attempts_failure_count` (Number)
- `onboarding_attempts_failure_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--onboarding_attempts_failure_impacted_entities))
- `onboarding_attempts_failure_metrics` (List of Object) (see [below for nested schema](#nestedobjatt--item--onboarding_attempts_failure_metrics))
- `onboarding_attempts_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--onboarding_attempts_impacted_entities))
- `onboarding_attempts_success_count` (Number)
- `onboarding_attempts_success_percentage` (Number)
- `onboarding_attempts_total_count` (Number)
- `onboarding_duration_average` (Number)
- `onboarding_duration_client_count` (Number)
- `onboarding_duration_failure_count` (Number)
- `onboarding_duration_failure_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--onboarding_duration_failure_impacted_entities))
- `onboarding_duration_failure_metrics` (List of Object) (see [below for nested schema](#nestedobjatt--item--onboarding_duration_failure_metrics))
- `onboarding_duration_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--onboarding_duration_impacted_entities))
- `onboarding_duration_success_count` (Number)
- `onboarding_duration_success_percentage` (Number)
- `onboarding_duration_total_count` (Number)
- `roaming_attempts_client_count` (Number)
- `roaming_attempts_failure_count` (Number)
- `roaming_attempts_failure_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--roaming_attempts_failure_impacted_entities))
- `roaming_attempts_failure_metrics` (List of Object) (see [below for nested schema](#nestedobjatt--item--roaming_attempts_failure_metrics))
- `roaming_attempts_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--roaming_attempts_impacted_entities))
- `roaming_attempts_success_count` (Number)
- `roaming_attempts_success_percentage` (Number)
- `roaming_attempts_total_count` (Number)
- `roaming_duration_average` (Number)
- `roaming_duration_client_count` (Number)
- `roaming_duration_failure_count` (Number)
- `roaming_duration_failure_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--roaming_duration_failure_impacted_entities))
- `roaming_duration_failure_metrics` (List of Object) (see [below for nested schema](#nestedobjatt--item--roaming_duration_failure_metrics))
- `roaming_duration_impacted_entities` (List of Object) (see [below for nested schema](#nestedobjatt--item--roaming_duration_impacted_entities))
- `roaming_duration_success_count` (Number)
- `roaming_duration_success_percentage` (Number)
- `roaming_duration_total_count` (Number)
- `site_hierarchy` (String)
- `site_hierarchy_id` (String)
- `site_id` (String)
- `site_type` (String)

<a id="nestedobjatt--item--connection_speed_failure_impacted_entities"></a>
### Nested Schema for `item.connection_speed_failure_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--connection_speed_failure_metrics"></a>
### Nested Schema for `item.connection_speed_failure_metrics`

Read-Only:

- `failure_ap_count` (Number)
- `failure_client_count` (Number)
- `failure_percentage` (Number)


<a id="nestedobjatt--item--connection_speed_impacted_entities"></a>
### Nested Schema for `item.connection_speed_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--coverage_failure_impacted_entities"></a>
### Nested Schema for `item.coverage_failure_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--coverage_failure_metrics"></a>
### Nested Schema for `item.coverage_failure_metrics`

Read-Only:

- `failure_ap_count` (Number)
- `failure_client_count` (Number)
- `failure_percentage` (Number)


<a id="nestedobjatt--item--coverage_impacted_entities"></a>
### Nested Schema for `item.coverage_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--onboarding_attempts_failure_impacted_entities"></a>
### Nested Schema for `item.onboarding_attempts_failure_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--onboarding_attempts_failure_metrics"></a>
### Nested Schema for `item.onboarding_attempts_failure_metrics`

Read-Only:

- `failure_ap_count` (Number)
- `failure_client_count` (Number)
- `failure_percentage` (Number)


<a id="nestedobjatt--item--onboarding_attempts_impacted_entities"></a>
### Nested Schema for `item.onboarding_attempts_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--onboarding_duration_failure_impacted_entities"></a>
### Nested Schema for `item.onboarding_duration_failure_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--onboarding_duration_failure_metrics"></a>
### Nested Schema for `item.onboarding_duration_failure_metrics`

Read-Only:

- `failure_ap_count` (Number)
- `failure_client_count` (Number)
- `failure_percentage` (Number)


<a id="nestedobjatt--item--onboarding_duration_impacted_entities"></a>
### Nested Schema for `item.onboarding_duration_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--roaming_attempts_failure_impacted_entities"></a>
### Nested Schema for `item.roaming_attempts_failure_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--roaming_attempts_failure_metrics"></a>
### Nested Schema for `item.roaming_attempts_failure_metrics`

Read-Only:

- `failure_ap_count` (Number)
- `failure_client_count` (Number)
- `failure_percentage` (Number)


<a id="nestedobjatt--item--roaming_attempts_impacted_entities"></a>
### Nested Schema for `item.roaming_attempts_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--roaming_duration_failure_impacted_entities"></a>
### Nested Schema for `item.roaming_duration_failure_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)


<a id="nestedobjatt--item--roaming_duration_failure_metrics"></a>
### Nested Schema for `item.roaming_duration_failure_metrics`

Read-Only:

- `failure_ap_count` (Number)
- `failure_client_count` (Number)
- `failure_percentage` (Number)


<a id="nestedobjatt--item--roaming_duration_impacted_entities"></a>
### Nested Schema for `item.roaming_duration_impacted_entities`

Read-Only:

- `ap_count` (Number)
- `building_count` (Number)
- `floor_count` (Number)
- `sites_count` (Number)
