/*

Copyright 2018 This Project Authors.

Author:  seanchann <seanchann@foxmail.com>

See docs/ for more information about the  project.

*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/seanchann/sample-apiserver/pkg/client/clientset/versioned/typed/sample/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSampleV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSampleV1alpha1) Tests(namespace string) v1alpha1.TestInterface {
	return &FakeTests{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSampleV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}