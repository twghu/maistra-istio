// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1alpha1 "k8s.io/api/discovery/v1alpha1"
	informers "k8s.io/client-go/informers/discovery/v1alpha1"
	listers "k8s.io/client-go/listers/discovery/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

type endpointSliceInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.EndpointSliceInformer = &endpointSliceInformer{}

func NewEndpointSliceInformer(f xnsinformers.SharedInformerFactory) informers.EndpointSliceInformer {
	resource := v1alpha1.SchemeGroupVersion.WithResource("endpointslices")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1alpha1.EndpointSlice{},
		&v1alpha1.EndpointSliceList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &endpointSliceInformer{informer: informer.Informer()}
}

func (i *endpointSliceInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *endpointSliceInformer) Lister() listers.EndpointSliceLister {
	return listers.NewEndpointSliceLister(i.informer.GetIndexer())
}