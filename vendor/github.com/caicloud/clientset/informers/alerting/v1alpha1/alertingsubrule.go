/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1alpha1 "github.com/caicloud/clientset/listers/alerting/v1alpha1"
	alertingv1alpha1 "github.com/caicloud/clientset/pkg/apis/alerting/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// AlertingSubRuleInformer provides access to a shared informer and lister for
// AlertingSubRules.
type AlertingSubRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AlertingSubRuleLister
}

type alertingSubRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAlertingSubRuleInformer constructs a new informer for AlertingSubRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAlertingSubRuleInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAlertingSubRuleInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAlertingSubRuleInformer constructs a new informer for AlertingSubRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAlertingSubRuleInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertingV1alpha1().AlertingSubRules().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AlertingV1alpha1().AlertingSubRules().Watch(options)
			},
		},
		&alertingv1alpha1.AlertingSubRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *alertingSubRuleInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAlertingSubRuleInformer(client.(kubernetes.Interface), resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *alertingSubRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&alertingv1alpha1.AlertingSubRule{}, f.defaultInformer)
}

func (f *alertingSubRuleInformer) Lister() v1alpha1.AlertingSubRuleLister {
	return v1alpha1.NewAlertingSubRuleLister(f.Informer().GetIndexer())
}
