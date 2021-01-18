// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha3

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	informers "istio.io/client-go/pkg/informers/externalversions/networking/v1alpha3"
	listers "istio.io/client-go/pkg/listers/networking/v1alpha3"
	"k8s.io/client-go/tools/cache"
)

type gatewayInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.GatewayInformer = &gatewayInformer{}

func NewGatewayInformer(f xnsinformers.SharedInformerFactory) informers.GatewayInformer {
	resource := v1alpha3.SchemeGroupVersion.WithResource("gateways")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1alpha3.Gateway{},
		&v1alpha3.GatewayList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &gatewayInformer{informer: informer.Informer()}
}

func (i *gatewayInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *gatewayInformer) Lister() listers.GatewayLister {
	return listers.NewGatewayLister(i.informer.GetIndexer())
}