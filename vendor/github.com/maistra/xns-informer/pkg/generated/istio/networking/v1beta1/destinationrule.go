// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	informers "istio.io/client-go/pkg/informers/externalversions/networking/v1beta1"
	listers "istio.io/client-go/pkg/listers/networking/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type destinationRuleInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.DestinationRuleInformer = &destinationRuleInformer{}

func NewDestinationRuleInformer(f xnsinformers.SharedInformerFactory) informers.DestinationRuleInformer {
	resource := v1beta1.SchemeGroupVersion.WithResource("destinationrules")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1beta1.DestinationRule{},
		&v1beta1.DestinationRuleList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &destinationRuleInformer{informer: informer.Informer()}
}

func (i *destinationRuleInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *destinationRuleInformer) Lister() listers.DestinationRuleLister {
	return listers.NewDestinationRuleLister(i.informer.GetIndexer())
}