// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mlflowexperiment


type MlflowExperimentTraceLocation struct {
	// uc_trace_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/databricks/databricks/1.126.0/docs/resources/mlflow_experiment#uc_trace_location MlflowExperiment#uc_trace_location}
	UcTraceLocation *MlflowExperimentTraceLocationUcTraceLocation `field:"optional" json:"ucTraceLocation" yaml:"ucTraceLocation"`
}

