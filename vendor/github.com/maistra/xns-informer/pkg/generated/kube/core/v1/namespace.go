// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/core/v1"
	informers "k8s.io/client-go/informers/core/v1"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type namespaceInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.NamespaceInformer = &namespaceInformer{}

func NewNamespaceInformer(f xnsinformers.SharedInformerFactory) informers.NamespaceInformer {
	resource := v1.SchemeGroupVersion.WithResource("namespaces")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1.Namespace{},
		&v1.NamespaceList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      true,
		ListWatchConverter: converter,
	})

	return &namespaceInformer{informer: informer.Informer()}
}

func (i *namespaceInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *namespaceInformer) Lister() listers.NamespaceLister {
	return listers.NewNamespaceLister(i.informer.GetIndexer())
}