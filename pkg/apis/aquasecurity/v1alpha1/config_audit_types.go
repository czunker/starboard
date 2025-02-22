package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ConfigAuditReportCRName    = "configauditreports.aquasecurity.github.io"
	ConfigAuditReportCRVersion = "v1alpha1"
	ConfigAuditReportKind      = "ConfigAuditReport"
	ConfigAuditReportListKind  = "ConfigAuditReportList"

	ClusterConfigAuditReportCRName = "clusterconfigauditreports.aquasecurity.github.io"
)

const (
	ConfigAuditSeverityDanger  = "danger"
	ConfigAuditSeverityWarning = "warning"
)

type ConfigAuditSummary struct {
	PassCount    int `json:"passCount"`
	DangerCount  int `json:"dangerCount"`
	WarningCount int `json:"warningCount"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigAuditReport is a specification for the ConfigAuditReport resource.
type ConfigAuditReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Report ConfigAuditReportData `json:"report"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigAuditReportList is a list of AuditConfig resources.
type ConfigAuditReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ConfigAuditReport `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterConfigAuditReport is a specification for the ClusterConfigAuditReport resource.
type ClusterConfigAuditReport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Report ConfigAuditReportData `json:"report"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterConfigAuditReportList is a list of ClusterConfigAuditReport resources.
type ClusterConfigAuditReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterConfigAuditReport `json:"items"`
}

type ConfigAuditReportData struct {
	UpdateTimestamp metav1.Time        `json:"updateTimestamp"`
	Scanner         Scanner            `json:"scanner"`
	Summary         ConfigAuditSummary `json:"summary"`

	// Checks provides results of conducting audit steps.
	Checks []Check `json:"checks"`

	// Deprecated in 0.12+ use Checks with CheckScope instead
	PodChecks []Check `json:"podChecks"`
	// Deprecated in 0.12+ use Checks with CheckScope instead
	ContainerChecks map[string][]Check `json:"containerChecks"`
}

// CheckScope has Type and Value fields to further identify a given Check.
// For example, we can use `Container` as Type and `nginx` as Value to indicate
// that a particular check is relevant to the nginx container. Alternatively,
// Type may be `JSONPath` and the Value would be JSONPath expression, e.g.
// `.spec.container[0].securityContext.allowPrivilegeEscalation`.
//
// Another use case for CheckScope is to inspect a ConfigMap with many keys and
// indicate a troublesome key. In this case the Type would be `ConfigMapKey`
// and the Value will hold the name of a key, e.g. `myawsprivatekey`.
type CheckScope struct {

	// Type indicates type of this scope, e.g. Container, ConfigMapKey or JSONPath.
	Type string `json:"type"`

	// Value indicates value of this scope that depends on Type, e.g. container name, ConfigMap key or JSONPath expression
	Value string `json:"value"`
}

// Check provides the result of conducting a single audit step.
type Check struct {
	ID      string `json:"checkID"`
	Message string `json:"message"`

	// Remediation provides description or links to external resources to remediate failing check.
	// +optional
	Remediation string `json:"remediation,omitempty"`

	Success  bool   `json:"success"`
	Severity string `json:"severity"`
	Category string `json:"category"`

	// Scope indicates the section of config that was audited.
	// +optional
	Scope *CheckScope `json:"scope,omitempty"`
}
