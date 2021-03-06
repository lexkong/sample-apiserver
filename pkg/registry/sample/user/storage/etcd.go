package storage

import (
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
)

// REST implements the REST endpoint for usertoken
type etcdREST struct {
	*genericregistry.Store
}

// // NewREST returns a RESTStorage object that will work with testtype.
// func NewREST(opts generic.RESTOptions) *REST {
// 	prefix := "/" + opts.ResourcePrefix
// 	newListFunc := func() runtime.Object { return &api.UserList{} }

// 	storageConfig := opts.StorageConfig
// 	storageInterface, dFunc := generic.NewRawStorage(storageConfig)

// 	store := &registry.Store{
// 		NewFunc:     func() runtime.Object { return &api.User{} },
// 		NewListFunc: newListFunc,
// 		KeyRootFunc: func(ctx freezerapi.Context) string {
// 			return prefix
// 		},
// 		KeyFunc: func(ctx freezerapi.Context, name string) (string, error) {
// 			return registry.NoNamespaceKeyFunc(ctx, prefix, name)
// 		},
// 		ObjectNameFunc: func(obj runtime.Object) (string, error) {
// 			return obj.(*api.User).Name, nil
// 		},
// 		PredicateFunc:           user.MatchUser,
// 		QualifiedResource:       api.Resource("users"),
// 		EnableGarbageCollection: opts.EnableGarbageCollection,
// 		DeleteCollectionWorkers: opts.DeleteCollectionWorkers,

// 		CreateStrategy:      user.Strategy,
// 		UpdateStrategy:      user.Strategy,
// 		DeleteStrategy:      user.Strategy,
// 		ReturnDeletedObject: true,
// 		// AfterCreate:         node.PadObj,
// 		// AfterUpdate:         node.PadObj,
// 		// AfterDelete:         node.PadObj,

// 		Storage:     storageInterface,
// 		DestroyFunc: dFunc,
// 	}

// 	return &REST{
// 		etcdregistry.NewStore(*store),
// 	}

// }
