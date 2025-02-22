package matcher

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"

	"fmt"

	"github.com/aquasecurity/starboard/pkg/apis/aquasecurity/v1alpha1"
	"github.com/aquasecurity/starboard/pkg/starboard"
	"github.com/onsi/gomega/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

var (
	trivyScanner = v1alpha1.Scanner{
		Name:    "Trivy",
		Vendor:  "Aqua Security",
		Version: "0.19.2",
	}
	polarisScanner = v1alpha1.Scanner{
		Name:    "Polaris",
		Vendor:  "Fairwinds Ops",
		Version: "4.0",
	}
	conftestScanner = v1alpha1.Scanner{
		Name:    "Conftest",
		Vendor:  "Open Policy Agent",
		Version: "v0.25.0",
	}
)

// IsVulnerabilityReportForContainerOwnedBy succeeds if a v1alpha1.VulnerabilityReport has a valid structure,
// corresponds to the given container and is owned by the specified client.Object.
//
// Note: This matcher is not suitable for unit tests because it does not perform a strict validation
// of the actual v1alpha1.VulnerabilityReport.
func IsVulnerabilityReportForContainerOwnedBy(containerName string, owner client.Object) types.GomegaMatcher {
	return &vulnerabilityReportMatcher{
		containerName: containerName,
		owner:         owner,
	}
}

type vulnerabilityReportMatcher struct {
	owner                 client.Object
	containerName         string
	failureMessage        string
	negatedFailureMessage string
}

func (m *vulnerabilityReportMatcher) Match(actual interface{}) (bool, error) {
	_, ok := actual.(v1alpha1.VulnerabilityReport)
	if !ok {
		return false, fmt.Errorf("%T expects a %T", vulnerabilityReportMatcher{}, v1alpha1.VulnerabilityReport{})
	}
	gvk, err := apiutil.GVKForObject(m.owner, starboard.NewScheme())
	if err != nil {
		return false, err
	}

	matcher := MatchFields(IgnoreExtras, Fields{
		"ObjectMeta": MatchFields(IgnoreExtras, Fields{
			"Labels": MatchKeys(IgnoreExtras, Keys{
				starboard.LabelContainerName:     Equal(m.containerName),
				starboard.LabelResourceKind:      Equal(gvk.Kind),
				starboard.LabelResourceName:      Equal(m.owner.GetName()),
				starboard.LabelResourceNamespace: Equal(m.owner.GetNamespace()),
			}),
			"OwnerReferences": ConsistOf(metav1.OwnerReference{
				APIVersion:         gvk.GroupVersion().Identifier(),
				Kind:               gvk.Kind,
				Name:               m.owner.GetName(),
				UID:                m.owner.GetUID(),
				Controller:         pointer.BoolPtr(true),
				BlockOwnerDeletion: pointer.BoolPtr(false),
			}),
		}),
		"Report": MatchFields(IgnoreExtras, Fields{
			"Scanner":         Equal(trivyScanner),
			"Vulnerabilities": Not(BeNil()),
		}),
	})

	success, err := matcher.Match(actual)
	if err != nil {
		return false, err
	}
	m.failureMessage = matcher.FailureMessage(actual)
	m.negatedFailureMessage = matcher.NegatedFailureMessage(actual)
	return success, nil
}

func (m *vulnerabilityReportMatcher) FailureMessage(_ interface{}) string {
	// TODO Add more descriptive message rather than rely on composed matchers' defaults
	return m.failureMessage
}

func (m *vulnerabilityReportMatcher) NegatedFailureMessage(_ interface{}) string {
	return m.negatedFailureMessage
}

// IsConfigAuditReportOwnedBy succeeds if a v1alpha1.ConfigAuditReport has a valid structure,
// and is owned by the specified client.Object.
//
// Note: This matcher is not suitable for unit tests because it does not perform a strict validation
// of the actual v1alpha1.ConfigAuditReport.
func IsConfigAuditReportOwnedBy(owner client.Object, scanner starboard.Scanner) types.GomegaMatcher {
	return &configAuditReportMatcher{
		owner:   owner,
		scanner: scanner,
	}
}

type configAuditReportMatcher struct {
	owner                 client.Object
	scanner               starboard.Scanner
	failureMessage        string
	negatedFailureMessage string
}

func (m *configAuditReportMatcher) Match(actual interface{}) (bool, error) {
	_, ok := actual.(v1alpha1.ConfigAuditReport)
	if !ok {
		return false, fmt.Errorf("%T expects a %T", configAuditReportMatcher{}, v1alpha1.ConfigAuditReport{})
	}
	gvk, err := apiutil.GVKForObject(m.owner, starboard.NewScheme())
	if err != nil {
		return false, err
	}

	scanner := polarisScanner
	if m.scanner == starboard.Conftest {
		scanner = conftestScanner
	}

	matcher := MatchFields(IgnoreExtras, Fields{
		"ObjectMeta": MatchFields(IgnoreExtras, Fields{
			"Labels": MatchKeys(IgnoreExtras, Keys{
				starboard.LabelResourceKind:      Equal(gvk.Kind),
				starboard.LabelResourceName:      Equal(m.owner.GetName()),
				starboard.LabelResourceNamespace: Equal(m.owner.GetNamespace()),
			}),
			"OwnerReferences": ConsistOf(metav1.OwnerReference{
				APIVersion:         gvk.GroupVersion().Identifier(),
				Kind:               gvk.Kind,
				Name:               m.owner.GetName(),
				UID:                m.owner.GetUID(),
				Controller:         pointer.BoolPtr(true),
				BlockOwnerDeletion: pointer.BoolPtr(false),
			}),
		}),
		"Report": MatchFields(IgnoreExtras, Fields{
			"Scanner": Equal(scanner),
		}),
	})
	success, err := matcher.Match(actual)
	if err != nil {
		return false, err
	}
	m.failureMessage = matcher.FailureMessage(actual)
	m.negatedFailureMessage = matcher.NegatedFailureMessage(actual)
	return success, nil
}

func (m *configAuditReportMatcher) FailureMessage(_ interface{}) string {
	// TODO Add more descriptive message rather than rely on composed matchers' defaults
	return m.failureMessage
}

func (m *configAuditReportMatcher) NegatedFailureMessage(_ interface{}) string {
	return m.negatedFailureMessage
}
