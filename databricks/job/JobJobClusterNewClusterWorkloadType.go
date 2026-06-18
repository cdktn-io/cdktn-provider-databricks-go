// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package job


type JobJobClusterNewClusterWorkloadType struct {
	// clients block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.118.0/docs/resources/job#clients Job#clients}
	Clients *JobJobClusterNewClusterWorkloadTypeClients `field:"required" json:"clients" yaml:"clients"`
}

