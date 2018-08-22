/*

Copyright 2018 This Project Authors.

Author:  seanchann <seanchann@foxmail.com>

See docs/ for more information about the  project.

*/

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	time "time"

	sample "github.com/seanchann/sample-apiserver/pkg/apis/sample"
	clientsetinternalversion "github.com/seanchann/sample-apiserver/pkg/client/clientset/internalversion"
	internalinterfaces "github.com/seanchann/sample-apiserver/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/seanchann/sample-apiserver/pkg/client/listers/sample/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TestInformer provides access to a shared informer and lister for
// Tests.
type TestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.TestLister
}

type testInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTestInformer constructs a new informer for Test type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTestInformer constructs a new informer for Test type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Sample().Tests(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Sample().Tests(namespace).Watch(options)
			},
		},
		&sample.Test{},
		resyncPeriod,
		indexers,
	)
}

func (f *testInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *testInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sample.Test{}, f.defaultInformer)
}

func (f *testInformer) Lister() internalversion.TestLister {
	return internalversion.NewTestLister(f.Informer().GetIndexer())
}