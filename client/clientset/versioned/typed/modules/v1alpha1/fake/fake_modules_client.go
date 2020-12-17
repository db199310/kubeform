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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/client/clientset/versioned/typed/modules/v1alpha1"
)

type FakeModulesV1alpha1 struct {
	*testing.Fake
}

func (c *FakeModulesV1alpha1) AzureAppServices(namespace string) v1alpha1.AzureAppServiceInterface {
	return &FakeAzureAppServices{c, namespace}
}

func (c *FakeModulesV1alpha1) AzureFnApps(namespace string) v1alpha1.AzureFnAppInterface {
	return &FakeAzureFnApps{c, namespace}
}

func (c *FakeModulesV1alpha1) F4dpAzFnv1s(namespace string) v1alpha1.F4dpAzFnv1Interface {
	return &FakeF4dpAzFnv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) F4dpAzSqls(namespace string) v1alpha1.F4dpAzSqlInterface {
	return &FakeF4dpAzSqls{c, namespace}
}

func (c *FakeModulesV1alpha1) F4dpAzStgv1s(namespace string) v1alpha1.F4dpAzStgv1Interface {
	return &FakeF4dpAzStgv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzADLSv1s(namespace string) v1alpha1.SDPAzADLSv1Interface {
	return &FakeSDPAzADLSv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzAppv1s(namespace string) v1alpha1.SDPAzAppv1Interface {
	return &FakeSDPAzAppv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzFnv1s(namespace string) v1alpha1.SDPAzFnv1Interface {
	return &FakeSDPAzFnv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzFnv2s(namespace string) v1alpha1.SDPAzFnv2Interface {
	return &FakeSDPAzFnv2s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzSqlv1s(namespace string) v1alpha1.SDPAzSqlv1Interface {
	return &FakeSDPAzSqlv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzStgv1s(namespace string) v1alpha1.SDPAzStgv1Interface {
	return &FakeSDPAzStgv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzadfv1s(namespace string) v1alpha1.SDPAzadfv1Interface {
	return &FakeSDPAzadfv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzappserviceplanv1s(namespace string) v1alpha1.SDPAzappserviceplanv1Interface {
	return &FakeSDPAzappserviceplanv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzdbv1s(namespace string) v1alpha1.SDPAzdbv1Interface {
	return &FakeSDPAzdbv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzplatformeventsv1s(namespace string) v1alpha1.SDPAzplatformeventsv1Interface {
	return &FakeSDPAzplatformeventsv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) SDPAzsbv1s(namespace string) v1alpha1.SDPAzsbv1Interface {
	return &FakeSDPAzsbv1s{c, namespace}
}

func (c *FakeModulesV1alpha1) StratosAzStgv1s(namespace string) v1alpha1.StratosAzStgv1Interface {
	return &FakeStratosAzStgv1s{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeModulesV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
