package v1beta1

import (
	opnicorev1 "github.com/rancher/opni/pkg/apis/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
type BootstrapToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              *opnicorev1.BootstrapToken `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
type BootstrapTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BootstrapToken `json:"items"`
}

// +kubebuilder:object:root=true
type Keyring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Data              []byte `json:"data,omitempty"`
}

// +kubebuilder:object:root=true
type KeyringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Keyring `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&BootstrapToken{}, &BootstrapTokenList{},
		&Keyring{}, &KeyringList{},
	)
}
