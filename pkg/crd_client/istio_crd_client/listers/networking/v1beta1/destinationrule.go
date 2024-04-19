/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DestinationRuleLister helps list DestinationRules.
// All objects returned here must be treated as read-only.
type DestinationRuleLister interface {
	// List lists all DestinationRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DestinationRule, err error)
	// DestinationRules returns an object that can list and get DestinationRules.
	DestinationRules(namespace string) DestinationRuleNamespaceLister
	DestinationRuleListerExpansion
}

// destinationRuleLister implements the DestinationRuleLister interface.
type destinationRuleLister struct {
	indexer cache.Indexer
}

// NewDestinationRuleLister returns a new DestinationRuleLister.
func NewDestinationRuleLister(indexer cache.Indexer) DestinationRuleLister {
	return &destinationRuleLister{indexer: indexer}
}

// List lists all DestinationRules in the indexer.
func (s *destinationRuleLister) List(selector labels.Selector) (ret []*v1beta1.DestinationRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DestinationRule))
	})
	return ret, err
}

// DestinationRules returns an object that can list and get DestinationRules.
func (s *destinationRuleLister) DestinationRules(namespace string) DestinationRuleNamespaceLister {
	return destinationRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DestinationRuleNamespaceLister helps list and get DestinationRules.
// All objects returned here must be treated as read-only.
type DestinationRuleNamespaceLister interface {
	// List lists all DestinationRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DestinationRule, err error)
	// Get retrieves the DestinationRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.DestinationRule, error)
	DestinationRuleNamespaceListerExpansion
}

// destinationRuleNamespaceLister implements the DestinationRuleNamespaceLister
// interface.
type destinationRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DestinationRules in the indexer for a given namespace.
func (s destinationRuleNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.DestinationRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DestinationRule))
	})
	return ret, err
}

// Get retrieves the DestinationRule from the indexer for a given namespace and name.
func (s destinationRuleNamespaceLister) Get(name string) (*v1beta1.DestinationRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("destinationrule"), name)
	}
	return obj.(*v1beta1.DestinationRule), nil
}
