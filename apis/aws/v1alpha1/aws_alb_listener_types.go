package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AwsAlbListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAlbListenerSpec   `json:"spec,omitempty"`
	Status            AwsAlbListenerStatus `json:"status,omitempty"`
}

type AwsAlbListenerSpecDefaultActionFixedResponse struct {
	StatusCode  string `json:"status_code"`
	ContentType string `json:"content_type"`
	MessageBody string `json:"message_body"`
}

type AwsAlbListenerSpecDefaultActionAuthenticateCognito struct {
	SessionTimeout                   int               `json:"session_timeout"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
}

type AwsAlbListenerSpecDefaultActionAuthenticateOidc struct {
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `json:"on_unauthenticated_request"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	ClientId                         string            `json:"client_id"`
	ClientSecret                     string            `json:"client_secret"`
	Issuer                           string            `json:"issuer"`
	Scope                            string            `json:"scope"`
	SessionCookieName                string            `json:"session_cookie_name"`
	SessionTimeout                   int               `json:"session_timeout"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
}

type AwsAlbListenerSpecDefaultActionRedirect struct {
	Path       string `json:"path"`
	Port       string `json:"port"`
	Protocol   string `json:"protocol"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
	Host       string `json:"host"`
}

type AwsAlbListenerSpecDefaultAction struct {
	FixedResponse       []AwsAlbListenerSpecDefaultAction `json:"fixed_response"`
	AuthenticateCognito []AwsAlbListenerSpecDefaultAction `json:"authenticate_cognito"`
	AuthenticateOidc    []AwsAlbListenerSpecDefaultAction `json:"authenticate_oidc"`
	Type                string                            `json:"type"`
	Order               int                               `json:"order"`
	TargetGroupArn      string                            `json:"target_group_arn"`
	Redirect            []AwsAlbListenerSpecDefaultAction `json:"redirect"`
}

type AwsAlbListenerSpec struct {
	Arn             string               `json:"arn"`
	LoadBalancerArn string               `json:"load_balancer_arn"`
	Port            int                  `json:"port"`
	Protocol        string               `json:"protocol"`
	SslPolicy       string               `json:"ssl_policy"`
	CertificateArn  string               `json:"certificate_arn"`
	DefaultAction   []AwsAlbListenerSpec `json:"default_action"`
}

type AwsAlbListenerStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAlbListenerList is a list of AwsAlbListeners
type AwsAlbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAlbListener CRD objects
	Items []AwsAlbListener `json:"items,omitempty"`
}
