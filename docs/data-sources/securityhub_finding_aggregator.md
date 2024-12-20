---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_securityhub_finding_aggregator Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::SecurityHub::FindingAggregator
---

# awscc_securityhub_finding_aggregator (Data Source)

Data Source schema for AWS::SecurityHub::FindingAggregator



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `finding_aggregation_region` (String)
- `finding_aggregator_arn` (String)
- `region_linking_mode` (String) Indicates whether to aggregate findings from all of the available Regions in the current partition. Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.
 The selected option also determines how to use the Regions provided in the Regions list.
 The options are as follows:
  +   ``ALL_REGIONS`` - Aggregates findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them. 
  +   ``ALL_REGIONS_EXCEPT_SPECIFIED`` - Aggregates findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the ``Regions`` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them. 
  +   ``SPECIFIED_REGIONS`` - Aggregates findings only from the Regions listed in the ``Regions`` parameter. Security Hub does not automatically aggregate findings from new Regions. 
  +   ``NO_REGIONS`` - Aggregates no data because no Regions are selected as linked Regions.
- `regions` (Set of String) If ``RegionLinkingMode`` is ``ALL_REGIONS_EXCEPT_SPECIFIED``, then this is a space-separated list of Regions that don't replicate and send findings to the home Region.
 If ``RegionLinkingMode`` is ``SPECIFIED_REGIONS``, then this is a space-separated list of Regions that do replicate and send findings to the home Region. 
 An ``InvalidInputException`` error results if you populate this field while ``RegionLinkingMode`` is ``NO_REGIONS``.
