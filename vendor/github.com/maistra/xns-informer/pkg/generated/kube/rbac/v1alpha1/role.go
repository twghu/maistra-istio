// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1alpha1 "k8s.io/api/rbac/v1alpha1"
	informers "k8s.io/client-go/informers/rbac/v1alpha1"
	listers "k8s.io/client-go/listers/rbac/v1alpha1"
	"k8s.io/client-go/tools/cache"
)

type roleInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.RoleInformer = &roleInformer{}

func NewRoleInformer(f xnsinformers.SharedInformerFactory) informers.RoleInformer {
	resource := v1alpha1.SchemeGroupVersion.WithResource("roles")
	converter := xnsinformers.NewListWatchConverter(
		f.GetScheme(),
		&v1alpha1.Role{},
		&v1alpha1.RoleList{},
	)

	informer := f.ForResource(resource, xnsinformers.ResourceOptions{
		ClusterScoped:      false,
		ListWatchConverter: converter,
	})

	return &roleInformer{informer: informer.Informer()}
}

func (i *roleInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *roleInformer) Lister() listers.RoleLister {
	return listers.NewRoleLister(i.informer.GetIndexer())
}