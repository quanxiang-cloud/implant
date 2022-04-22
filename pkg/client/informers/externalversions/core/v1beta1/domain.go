/*
Copyright 2022 The OpenFunction Authors.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	corev1beta1 "github.com/openfunction/apis/core/v1beta1"
	versioned "github.com/openfunction/pkg/client/clientset/versioned"
	internalinterfaces "github.com/quanxiang-cloud/implant/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/quanxiang-cloud/implant/pkg/client/listers/core/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DomainInformer provides access to a shared informer and lister for
// Domains.
type DomainInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.DomainLister
}

type domainInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDomainInformer constructs a new informer for Domain type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDomainInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDomainInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDomainInformer constructs a new informer for Domain type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDomainInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().Domains(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().Domains(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1beta1.Domain{},
		resyncPeriod,
		indexers,
	)
}

func (f *domainInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDomainInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *domainInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1beta1.Domain{}, f.defaultInformer)
}

func (f *domainInformer) Lister() v1beta1.DomainLister {
	return v1beta1.NewDomainLister(f.Informer().GetIndexer())
}
