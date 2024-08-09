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

package v1alpha1

import (
	v1alpha1 "k8s.io/api/certificates/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ClusterTrustBundleLister helps list ClusterTrustBundles.
// All objects returned here must be treated as read-only.
type ClusterTrustBundleLister interface {
	// List lists all ClusterTrustBundles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterTrustBundle, err error)
	// Get retrieves the ClusterTrustBundle from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterTrustBundle, error)
	ClusterTrustBundleListerExpansion
}

// clusterTrustBundleLister implements the ClusterTrustBundleLister interface.
type clusterTrustBundleLister struct {
	listers.ResourceIndexer[*v1alpha1.ClusterTrustBundle]
}

// NewClusterTrustBundleLister returns a new ClusterTrustBundleLister.
func NewClusterTrustBundleLister(indexer cache.Indexer) ClusterTrustBundleLister {
	return &clusterTrustBundleLister{listers.New[*v1alpha1.ClusterTrustBundle](indexer, v1alpha1.Resource("clustertrustbundle"))}
}
