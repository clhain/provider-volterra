/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AllowedResponseCodesObservation struct {
}

type AllowedResponseCodesParameters struct {

	// +kubebuilder:validation:Required
	ResponseCode []*float64 `json:"responseCode" tf:"response_code,omitempty"`
}

type AnonymizationConfigObservation struct {
}

type AnonymizationConfigParameters struct {

	// +kubebuilder:validation:Optional
	Cookie []CookieParameters `json:"cookie,omitempty" tf:"cookie,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPHeader []HTTPHeaderParameters `json:"httpHeader,omitempty" tf:"http_header,omitempty"`

	// +kubebuilder:validation:Optional
	QueryParameter []QueryParameterParameters `json:"queryParameter,omitempty" tf:"query_parameter,omitempty"`
}

type AppFirewallObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppFirewallParameters struct {

	// +kubebuilder:validation:Optional
	AllowAllResponseCodes *bool `json:"allowAllResponseCodes,omitempty" tf:"allow_all_response_codes,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedResponseCodes []AllowedResponseCodesParameters `json:"allowedResponseCodes,omitempty" tf:"allowed_response_codes,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	Blocking *bool `json:"blocking,omitempty" tf:"blocking,omitempty"`

	// +kubebuilder:validation:Optional
	BlockingPage []BlockingPageParameters `json:"blockingPage,omitempty" tf:"blocking_page,omitempty"`

	// +kubebuilder:validation:Optional
	BotProtectionSetting []BotProtectionSettingParameters `json:"botProtectionSetting,omitempty" tf:"bot_protection_setting,omitempty"`

	// +kubebuilder:validation:Optional
	CustomAnonymization []CustomAnonymizationParameters `json:"customAnonymization,omitempty" tf:"custom_anonymization,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultAnonymization *bool `json:"defaultAnonymization,omitempty" tf:"default_anonymization,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultBotSetting *bool `json:"defaultBotSetting,omitempty" tf:"default_bot_setting,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultDetectionSettings *bool `json:"defaultDetectionSettings,omitempty" tf:"default_detection_settings,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DetectionSettings []DetectionSettingsParameters `json:"detectionSettings,omitempty" tf:"detection_settings,omitempty"`

	// +kubebuilder:validation:Optional
	Disable *bool `json:"disable,omitempty" tf:"disable,omitempty"`

	// +kubebuilder:validation:Optional
	DisableAnonymization *bool `json:"disableAnonymization,omitempty" tf:"disable_anonymization,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Monitoring *bool `json:"monitoring,omitempty" tf:"monitoring,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/clhain/provider-volterra/apis/volterra/v1alpha1.VolterraNamespace
	// +crossplane:generate:reference:extractor=github.com/clhain/provider-volterra/config/common.ExtractResourceName()
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Reference to a VolterraNamespace in volterra to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceRef *v1.Reference `json:"namespaceRef,omitempty" tf:"-"`

	// Selector for a VolterraNamespace in volterra to populate namespace.
	// +kubebuilder:validation:Optional
	NamespaceSelector *v1.Selector `json:"namespaceSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UseDefaultBlockingPage *bool `json:"useDefaultBlockingPage,omitempty" tf:"use_default_blocking_page,omitempty"`

	// +kubebuilder:validation:Optional
	UseLoadbalancerSetting *bool `json:"useLoadbalancerSetting,omitempty" tf:"use_loadbalancer_setting,omitempty"`
}

type AttackTypeSettingsObservation struct {
}

type AttackTypeSettingsParameters struct {

	// +kubebuilder:validation:Required
	DisabledAttackTypes []*string `json:"disabledAttackTypes" tf:"disabled_attack_types,omitempty"`
}

type BlockingPageObservation struct {
}

type BlockingPageParameters struct {

	// +kubebuilder:validation:Optional
	BlockingPage *string `json:"blockingPage,omitempty" tf:"blocking_page,omitempty"`

	// +kubebuilder:validation:Optional
	ResponseCode *string `json:"responseCode,omitempty" tf:"response_code,omitempty"`
}

type BotProtectionSettingObservation struct {
}

type BotProtectionSettingParameters struct {

	// +kubebuilder:validation:Optional
	GoodBotAction *string `json:"goodBotAction,omitempty" tf:"good_bot_action,omitempty"`

	// +kubebuilder:validation:Optional
	MaliciousBotAction *string `json:"maliciousBotAction,omitempty" tf:"malicious_bot_action,omitempty"`

	// +kubebuilder:validation:Optional
	SuspiciousBotAction *string `json:"suspiciousBotAction,omitempty" tf:"suspicious_bot_action,omitempty"`
}

type CookieObservation struct {
}

type CookieParameters struct {

	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`
}

type CustomAnonymizationObservation struct {
}

type CustomAnonymizationParameters struct {

	// +kubebuilder:validation:Optional
	AnonymizationConfig []AnonymizationConfigParameters `json:"anonymizationConfig,omitempty" tf:"anonymization_config,omitempty"`
}

type DetectionSettingsObservation struct {
}

type DetectionSettingsParameters struct {

	// +kubebuilder:validation:Optional
	DefaultViolationSettings *bool `json:"defaultViolationSettings,omitempty" tf:"default_violation_settings,omitempty"`

	// +kubebuilder:validation:Optional
	DisableSuppression *bool `json:"disableSuppression,omitempty" tf:"disable_suppression,omitempty"`

	// +kubebuilder:validation:Optional
	DisableThreatCampaigns *bool `json:"disableThreatCampaigns,omitempty" tf:"disable_threat_campaigns,omitempty"`

	// +kubebuilder:validation:Optional
	EnableSuppression *bool `json:"enableSuppression,omitempty" tf:"enable_suppression,omitempty"`

	// +kubebuilder:validation:Optional
	EnableThreatCampaigns *bool `json:"enableThreatCampaigns,omitempty" tf:"enable_threat_campaigns,omitempty"`

	// +kubebuilder:validation:Optional
	SignatureSelectionSetting []SignatureSelectionSettingParameters `json:"signatureSelectionSetting,omitempty" tf:"signature_selection_setting,omitempty"`

	// +kubebuilder:validation:Optional
	ViolationSettings []ViolationSettingsParameters `json:"violationSettings,omitempty" tf:"violation_settings,omitempty"`
}

type HTTPHeaderObservation struct {
}

type HTTPHeaderParameters struct {

	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`
}

type QueryParameterObservation struct {
}

type QueryParameterParameters struct {

	// +kubebuilder:validation:Optional
	QueryParamName *string `json:"queryParamName,omitempty" tf:"query_param_name,omitempty"`
}

type SignatureSelectionSettingObservation struct {
}

type SignatureSelectionSettingParameters struct {

	// +kubebuilder:validation:Optional
	AttackTypeSettings []AttackTypeSettingsParameters `json:"attackTypeSettings,omitempty" tf:"attack_type_settings,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultAttackTypeSettings *bool `json:"defaultAttackTypeSettings,omitempty" tf:"default_attack_type_settings,omitempty"`

	// +kubebuilder:validation:Optional
	HighMediumAccuracySignatures *bool `json:"highMediumAccuracySignatures,omitempty" tf:"high_medium_accuracy_signatures,omitempty"`

	// +kubebuilder:validation:Optional
	HighMediumLowAccuracySignatures *bool `json:"highMediumLowAccuracySignatures,omitempty" tf:"high_medium_low_accuracy_signatures,omitempty"`

	// +kubebuilder:validation:Optional
	OnlyHighAccuracySignatures *bool `json:"onlyHighAccuracySignatures,omitempty" tf:"only_high_accuracy_signatures,omitempty"`
}

type ViolationSettingsObservation struct {
}

type ViolationSettingsParameters struct {

	// +kubebuilder:validation:Required
	DisabledViolationTypes []*string `json:"disabledViolationTypes" tf:"disabled_violation_types,omitempty"`
}

// AppFirewallSpec defines the desired state of AppFirewall
type AppFirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppFirewallParameters `json:"forProvider"`
}

// AppFirewallStatus defines the observed state of AppFirewall.
type AppFirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppFirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppFirewall is the Schema for the AppFirewalls API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,volterra}
type AppFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppFirewallSpec   `json:"spec"`
	Status            AppFirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppFirewallList contains a list of AppFirewalls
type AppFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppFirewall `json:"items"`
}

// Repository type metadata.
var (
	AppFirewall_Kind             = "AppFirewall"
	AppFirewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppFirewall_Kind}.String()
	AppFirewall_KindAPIVersion   = AppFirewall_Kind + "." + CRDGroupVersion.String()
	AppFirewall_GroupVersionKind = CRDGroupVersion.WithKind(AppFirewall_Kind)
)

func init() {
	SchemeBuilder.Register(&AppFirewall{}, &AppFirewallList{})
}