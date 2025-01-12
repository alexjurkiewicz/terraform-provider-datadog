---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datadog_dashboard_list Data Source - terraform-provider-datadog"
subcategory: ""
description: |-
  Use this data source to retrieve information about an existing dashboard list, for use in other resources. In particular, it can be used in a dashboard to register it in the list.
---

# datadog_dashboard_list (Data Source)

Use this data source to retrieve information about an existing dashboard list, for use in other resources. In particular, it can be used in a dashboard to register it in the list.

## Example Usage

```terraform
data "datadog_dashboard_list" "test" {
  name = "My super list"
}

# Create a dashboard and register it in the list above.
resource "datadog_dashboard" "time" {
  title           = "TF Test Layout Dashboard"
  description     = "Created using the Datadog provider in Terraform"
  dashboard_lists = ["${data.datadog_dashboard_list.test.id}"]
  layout_type     = "ordered"
  is_read_only    = true
  widget {
    alert_graph_definition {
      alert_id  = "1234"
      viz_type  = "timeseries"
      title     = "Widget Title"
      live_span = "1h"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) A dashboard list name to limit the search.

### Read-Only

- **id** (String) The ID of this resource.


