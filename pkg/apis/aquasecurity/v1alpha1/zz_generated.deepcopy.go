//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Artifact) DeepCopyInto(out *Artifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Artifact.
func (in *Artifact) DeepCopy() *Artifact {
	if in == nil {
		return nil
	}
	out := new(Artifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubeBenchReport) DeepCopyInto(out *CISKubeBenchReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubeBenchReport.
func (in *CISKubeBenchReport) DeepCopy() *CISKubeBenchReport {
	if in == nil {
		return nil
	}
	out := new(CISKubeBenchReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CISKubeBenchReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubeBenchReportData) DeepCopyInto(out *CISKubeBenchReportData) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	out.Scanner = in.Scanner
	out.Summary = in.Summary
	if in.Sections != nil {
		in, out := &in.Sections, &out.Sections
		*out = make([]CISKubeBenchSection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubeBenchReportData.
func (in *CISKubeBenchReportData) DeepCopy() *CISKubeBenchReportData {
	if in == nil {
		return nil
	}
	out := new(CISKubeBenchReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubeBenchReportList) DeepCopyInto(out *CISKubeBenchReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CISKubeBenchReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubeBenchReportList.
func (in *CISKubeBenchReportList) DeepCopy() *CISKubeBenchReportList {
	if in == nil {
		return nil
	}
	out := new(CISKubeBenchReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CISKubeBenchReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubeBenchResult) DeepCopyInto(out *CISKubeBenchResult) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubeBenchResult.
func (in *CISKubeBenchResult) DeepCopy() *CISKubeBenchResult {
	if in == nil {
		return nil
	}
	out := new(CISKubeBenchResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubeBenchSection) DeepCopyInto(out *CISKubeBenchSection) {
	*out = *in
	if in.Tests != nil {
		in, out := &in.Tests, &out.Tests
		*out = make([]CISKubeBenchTests, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubeBenchSection.
func (in *CISKubeBenchSection) DeepCopy() *CISKubeBenchSection {
	if in == nil {
		return nil
	}
	out := new(CISKubeBenchSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubeBenchSummary) DeepCopyInto(out *CISKubeBenchSummary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubeBenchSummary.
func (in *CISKubeBenchSummary) DeepCopy() *CISKubeBenchSummary {
	if in == nil {
		return nil
	}
	out := new(CISKubeBenchSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CISKubeBenchTests) DeepCopyInto(out *CISKubeBenchTests) {
	*out = *in
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]CISKubeBenchResult, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CISKubeBenchTests.
func (in *CISKubeBenchTests) DeepCopy() *CISKubeBenchTests {
	if in == nil {
		return nil
	}
	out := new(CISKubeBenchTests)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Check) DeepCopyInto(out *Check) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(CheckScope)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Check.
func (in *Check) DeepCopy() *Check {
	if in == nil {
		return nil
	}
	out := new(Check)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckScope) DeepCopyInto(out *CheckScope) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckScope.
func (in *CheckScope) DeepCopy() *CheckScope {
	if in == nil {
		return nil
	}
	out := new(CheckScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigAuditReport) DeepCopyInto(out *ClusterConfigAuditReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigAuditReport.
func (in *ClusterConfigAuditReport) DeepCopy() *ClusterConfigAuditReport {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigAuditReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigAuditReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigAuditReportList) DeepCopyInto(out *ClusterConfigAuditReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterConfigAuditReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigAuditReportList.
func (in *ClusterConfigAuditReportList) DeepCopy() *ClusterConfigAuditReportList {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigAuditReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigAuditReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVulnerabilityReport) DeepCopyInto(out *ClusterVulnerabilityReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVulnerabilityReport.
func (in *ClusterVulnerabilityReport) DeepCopy() *ClusterVulnerabilityReport {
	if in == nil {
		return nil
	}
	out := new(ClusterVulnerabilityReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVulnerabilityReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVulnerabilityReportList) DeepCopyInto(out *ClusterVulnerabilityReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterVulnerabilityReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVulnerabilityReportList.
func (in *ClusterVulnerabilityReportList) DeepCopy() *ClusterVulnerabilityReportList {
	if in == nil {
		return nil
	}
	out := new(ClusterVulnerabilityReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVulnerabilityReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditReport) DeepCopyInto(out *ConfigAuditReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditReport.
func (in *ConfigAuditReport) DeepCopy() *ConfigAuditReport {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigAuditReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditReportData) DeepCopyInto(out *ConfigAuditReportData) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	out.Scanner = in.Scanner
	out.Summary = in.Summary
	if in.Checks != nil {
		in, out := &in.Checks, &out.Checks
		*out = make([]Check, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodChecks != nil {
		in, out := &in.PodChecks, &out.PodChecks
		*out = make([]Check, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ContainerChecks != nil {
		in, out := &in.ContainerChecks, &out.ContainerChecks
		*out = make(map[string][]Check, len(*in))
		for key, val := range *in {
			var outVal []Check
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]Check, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditReportData.
func (in *ConfigAuditReportData) DeepCopy() *ConfigAuditReportData {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditReportList) DeepCopyInto(out *ConfigAuditReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigAuditReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditReportList.
func (in *ConfigAuditReportList) DeepCopy() *ConfigAuditReportList {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigAuditReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigAuditSummary) DeepCopyInto(out *ConfigAuditSummary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigAuditSummary.
func (in *ConfigAuditSummary) DeepCopy() *ConfigAuditSummary {
	if in == nil {
		return nil
	}
	out := new(ConfigAuditSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeHunterReport) DeepCopyInto(out *KubeHunterReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeHunterReport.
func (in *KubeHunterReport) DeepCopy() *KubeHunterReport {
	if in == nil {
		return nil
	}
	out := new(KubeHunterReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeHunterReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeHunterReportData) DeepCopyInto(out *KubeHunterReportData) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	out.Scanner = in.Scanner
	out.Summary = in.Summary
	if in.Vulnerabilities != nil {
		in, out := &in.Vulnerabilities, &out.Vulnerabilities
		*out = make([]KubeHunterVulnerability, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeHunterReportData.
func (in *KubeHunterReportData) DeepCopy() *KubeHunterReportData {
	if in == nil {
		return nil
	}
	out := new(KubeHunterReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeHunterReportList) DeepCopyInto(out *KubeHunterReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeHunterReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeHunterReportList.
func (in *KubeHunterReportList) DeepCopy() *KubeHunterReportList {
	if in == nil {
		return nil
	}
	out := new(KubeHunterReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeHunterReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeHunterSummary) DeepCopyInto(out *KubeHunterSummary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeHunterSummary.
func (in *KubeHunterSummary) DeepCopy() *KubeHunterSummary {
	if in == nil {
		return nil
	}
	out := new(KubeHunterSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeHunterVulnerability) DeepCopyInto(out *KubeHunterVulnerability) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeHunterVulnerability.
func (in *KubeHunterVulnerability) DeepCopy() *KubeHunterVulnerability {
	if in == nil {
		return nil
	}
	out := new(KubeHunterVulnerability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scanner) DeepCopyInto(out *Scanner) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scanner.
func (in *Scanner) DeepCopy() *Scanner {
	if in == nil {
		return nil
	}
	out := new(Scanner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vulnerability) DeepCopyInto(out *Vulnerability) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Score != nil {
		in, out := &in.Score, &out.Score
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vulnerability.
func (in *Vulnerability) DeepCopy() *Vulnerability {
	if in == nil {
		return nil
	}
	out := new(Vulnerability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityReport) DeepCopyInto(out *VulnerabilityReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityReport.
func (in *VulnerabilityReport) DeepCopy() *VulnerabilityReport {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VulnerabilityReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityReportData) DeepCopyInto(out *VulnerabilityReportData) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	out.Scanner = in.Scanner
	out.Registry = in.Registry
	out.Artifact = in.Artifact
	out.Summary = in.Summary
	if in.Vulnerabilities != nil {
		in, out := &in.Vulnerabilities, &out.Vulnerabilities
		*out = make([]Vulnerability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityReportData.
func (in *VulnerabilityReportData) DeepCopy() *VulnerabilityReportData {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityReportData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityReportList) DeepCopyInto(out *VulnerabilityReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VulnerabilityReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityReportList.
func (in *VulnerabilityReportList) DeepCopy() *VulnerabilityReportList {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VulnerabilityReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilitySummary) DeepCopyInto(out *VulnerabilitySummary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilitySummary.
func (in *VulnerabilitySummary) DeepCopy() *VulnerabilitySummary {
	if in == nil {
		return nil
	}
	out := new(VulnerabilitySummary)
	in.DeepCopyInto(out)
	return out
}
