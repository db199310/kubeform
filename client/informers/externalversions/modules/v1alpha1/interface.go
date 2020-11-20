/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AzureAppServices returns a AzureAppServiceInformer.
	AzureAppServices() AzureAppServiceInformer
	// AzureFnApps returns a AzureFnAppInformer.
	AzureFnApps() AzureFnAppInformer
	// F4dpAzFnv1s returns a F4dpAzFnv1Informer.
	F4dpAzFnv1s() F4dpAzFnv1Informer
	// F4dpAzSqls returns a F4dpAzSqlInformer.
	F4dpAzSqls() F4dpAzSqlInformer
	// F4dpAzStgv1s returns a F4dpAzStgv1Informer.
	F4dpAzStgv1s() F4dpAzStgv1Informer
	// SDPAzAppv1s returns a SDPAzAppv1Informer.
	SDPAzAppv1s() SDPAzAppv1Informer
	// SDPAzFnv1s returns a SDPAzFnv1Informer.
	SDPAzFnv1s() SDPAzFnv1Informer
	// SDPAzSqlv1s returns a SDPAzSqlv1Informer.
	SDPAzSqlv1s() SDPAzSqlv1Informer
	// SDPAzStgv1s returns a SDPAzStgv1Informer.
	SDPAzStgv1s() SDPAzStgv1Informer
	// SDPAzadfv1s returns a SDPAzadfv1Informer.
	SDPAzadfv1s() SDPAzadfv1Informer
	// SDPAzappserviceplanv1s returns a SDPAzappserviceplanv1Informer.
	SDPAzappserviceplanv1s() SDPAzappserviceplanv1Informer
	// SDPAzsbv1s returns a SDPAzsbv1Informer.
	SDPAzsbv1s() SDPAzsbv1Informer
	// StratosAzStgv1s returns a StratosAzStgv1Informer.
	StratosAzStgv1s() StratosAzStgv1Informer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AzureAppServices returns a AzureAppServiceInformer.
func (v *version) AzureAppServices() AzureAppServiceInformer {
	return &azureAppServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AzureFnApps returns a AzureFnAppInformer.
func (v *version) AzureFnApps() AzureFnAppInformer {
	return &azureFnAppInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// F4dpAzFnv1s returns a F4dpAzFnv1Informer.
func (v *version) F4dpAzFnv1s() F4dpAzFnv1Informer {
	return &f4dpAzFnv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// F4dpAzSqls returns a F4dpAzSqlInformer.
func (v *version) F4dpAzSqls() F4dpAzSqlInformer {
	return &f4dpAzSqlInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// F4dpAzStgv1s returns a F4dpAzStgv1Informer.
func (v *version) F4dpAzStgv1s() F4dpAzStgv1Informer {
	return &f4dpAzStgv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDPAzAppv1s returns a SDPAzAppv1Informer.
func (v *version) SDPAzAppv1s() SDPAzAppv1Informer {
	return &sDPAzAppv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDPAzFnv1s returns a SDPAzFnv1Informer.
func (v *version) SDPAzFnv1s() SDPAzFnv1Informer {
	return &sDPAzFnv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDPAzSqlv1s returns a SDPAzSqlv1Informer.
func (v *version) SDPAzSqlv1s() SDPAzSqlv1Informer {
	return &sDPAzSqlv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDPAzStgv1s returns a SDPAzStgv1Informer.
func (v *version) SDPAzStgv1s() SDPAzStgv1Informer {
	return &sDPAzStgv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDPAzadfv1s returns a SDPAzadfv1Informer.
func (v *version) SDPAzadfv1s() SDPAzadfv1Informer {
	return &sDPAzadfv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDPAzappserviceplanv1s returns a SDPAzappserviceplanv1Informer.
func (v *version) SDPAzappserviceplanv1s() SDPAzappserviceplanv1Informer {
	return &sDPAzappserviceplanv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SDPAzsbv1s returns a SDPAzsbv1Informer.
func (v *version) SDPAzsbv1s() SDPAzsbv1Informer {
	return &sDPAzsbv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StratosAzStgv1s returns a StratosAzStgv1Informer.
func (v *version) StratosAzStgv1s() StratosAzStgv1Informer {
	return &stratosAzStgv1Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
