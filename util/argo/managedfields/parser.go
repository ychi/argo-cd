package managedfields

import (
	"sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func StaticParser() *typed.Parser {
	parserOnce.Do(func() {
		parser, _ = typed.NewParser(schemaYAML)
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: io.k8s.api.admissionregistration.v1.MutatingWebhook
  map:
    fields:
    - name: admissionReviewVersions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: clientConfig
      type:
        namedType: io.k8s.api.admissionregistration.v1.WebhookClientConfig
      default: {}
    - name: failurePolicy
      type:
        scalar: string
    - name: matchPolicy
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespaceSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: objectSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: reinvocationPolicy
      type:
        scalar: string
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.admissionregistration.v1.RuleWithOperations
          elementRelationship: atomic
    - name: sideEffects
      type:
        scalar: string
    - name: timeoutSeconds
      type:
        scalar: numeric
- name: io.k8s.api.admissionregistration.v1.MutatingWebhookConfiguration
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: webhooks
      type:
        list:
          elementType:
            namedType: io.k8s.api.admissionregistration.v1.MutatingWebhook
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.api.admissionregistration.v1.RuleWithOperations
  map:
    fields:
    - name: apiGroups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: apiVersions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: operations
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resources
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: scope
      type:
        scalar: string
- name: io.k8s.api.admissionregistration.v1.ServiceReference
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
    - name: path
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
- name: io.k8s.api.admissionregistration.v1.ValidatingWebhook
  map:
    fields:
    - name: admissionReviewVersions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: clientConfig
      type:
        namedType: io.k8s.api.admissionregistration.v1.WebhookClientConfig
      default: {}
    - name: failurePolicy
      type:
        scalar: string
    - name: matchPolicy
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespaceSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: objectSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.admissionregistration.v1.RuleWithOperations
          elementRelationship: atomic
    - name: sideEffects
      type:
        scalar: string
    - name: timeoutSeconds
      type:
        scalar: numeric
- name: io.k8s.api.admissionregistration.v1.ValidatingWebhookConfiguration
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: webhooks
      type:
        list:
          elementType:
            namedType: io.k8s.api.admissionregistration.v1.ValidatingWebhook
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.api.admissionregistration.v1.WebhookClientConfig
  map:
    fields:
    - name: caBundle
      type:
        scalar: string
    - name: service
      type:
        namedType: io.k8s.api.admissionregistration.v1.ServiceReference
    - name: url
      type:
        scalar: string
- name: io.k8s.api.admissionregistration.v1beta1.MutatingWebhook
  map:
    fields:
    - name: admissionReviewVersions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: clientConfig
      type:
        namedType: io.k8s.api.admissionregistration.v1beta1.WebhookClientConfig
      default: {}
    - name: failurePolicy
      type:
        scalar: string
    - name: matchPolicy
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespaceSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: objectSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: reinvocationPolicy
      type:
        scalar: string
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.admissionregistration.v1beta1.RuleWithOperations
          elementRelationship: atomic
    - name: sideEffects
      type:
        scalar: string
    - name: timeoutSeconds
      type:
        scalar: numeric
- name: io.k8s.api.admissionregistration.v1beta1.MutatingWebhookConfiguration
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: webhooks
      type:
        list:
          elementType:
            namedType: io.k8s.api.admissionregistration.v1beta1.MutatingWebhook
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.api.admissionregistration.v1beta1.RuleWithOperations
  map:
    fields:
    - name: apiGroups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: apiVersions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: operations
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resources
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: scope
      type:
        scalar: string
- name: io.k8s.api.admissionregistration.v1beta1.ServiceReference
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
    - name: path
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
- name: io.k8s.api.admissionregistration.v1beta1.ValidatingWebhook
  map:
    fields:
    - name: admissionReviewVersions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: clientConfig
      type:
        namedType: io.k8s.api.admissionregistration.v1beta1.WebhookClientConfig
      default: {}
    - name: failurePolicy
      type:
        scalar: string
    - name: matchPolicy
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespaceSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: objectSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.admissionregistration.v1beta1.RuleWithOperations
          elementRelationship: atomic
    - name: sideEffects
      type:
        scalar: string
    - name: timeoutSeconds
      type:
        scalar: numeric
- name: io.k8s.api.admissionregistration.v1beta1.ValidatingWebhookConfiguration
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: webhooks
      type:
        list:
          elementType:
            namedType: io.k8s.api.admissionregistration.v1beta1.ValidatingWebhook
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.api.admissionregistration.v1beta1.WebhookClientConfig
  map:
    fields:
    - name: caBundle
      type:
        scalar: string
    - name: service
      type:
        namedType: io.k8s.api.admissionregistration.v1beta1.ServiceReference
    - name: url
      type:
        scalar: string
- name: io.k8s.api.apiserverinternal.v1alpha1.ServerStorageVersion
  map:
    fields:
    - name: apiServerID
      type:
        scalar: string
    - name: decodableVersions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: encodingVersion
      type:
        scalar: string
- name: io.k8s.api.apiserverinternal.v1alpha1.StorageVersion
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apiserverinternal.v1alpha1.StorageVersionSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apiserverinternal.v1alpha1.StorageVersionStatus
      default: {}
- name: io.k8s.api.apiserverinternal.v1alpha1.StorageVersionCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: observedGeneration
      type:
        scalar: numeric
    - name: reason
      type:
        scalar: string
      default: ""
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apiserverinternal.v1alpha1.StorageVersionSpec
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.api.apiserverinternal.v1alpha1.StorageVersionStatus
  map:
    fields:
    - name: commonEncodingVersion
      type:
        scalar: string
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apiserverinternal.v1alpha1.StorageVersionCondition
          elementRelationship: associative
          keys:
          - type
    - name: storageVersions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apiserverinternal.v1alpha1.ServerStorageVersion
          elementRelationship: associative
          keys:
          - apiServerID
- name: io.k8s.api.apps.v1.ControllerRevision
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: data
      type:
        namedType: __untyped_atomic_
      default: {}
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: revision
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.apps.v1.DaemonSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1.DaemonSetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1.DaemonSetStatus
      default: {}
- name: io.k8s.api.apps.v1.DaemonSetCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1.DaemonSetSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
    - name: updateStrategy
      type:
        namedType: io.k8s.api.apps.v1.DaemonSetUpdateStrategy
      default: {}
- name: io.k8s.api.apps.v1.DaemonSetStatus
  map:
    fields:
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1.DaemonSetCondition
          elementRelationship: associative
          keys:
          - type
    - name: currentNumberScheduled
      type:
        scalar: numeric
      default: 0
    - name: desiredNumberScheduled
      type:
        scalar: numeric
      default: 0
    - name: numberAvailable
      type:
        scalar: numeric
    - name: numberMisscheduled
      type:
        scalar: numeric
      default: 0
    - name: numberReady
      type:
        scalar: numeric
      default: 0
    - name: numberUnavailable
      type:
        scalar: numeric
    - name: observedGeneration
      type:
        scalar: numeric
    - name: updatedNumberScheduled
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1.DaemonSetUpdateStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.apps.v1.RollingUpdateDaemonSet
    - name: type
      type:
        scalar: string
- name: io.k8s.api.apps.v1.Deployment
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1.DeploymentSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1.DeploymentStatus
      default: {}
- name: io.k8s.api.apps.v1.DeploymentCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastUpdateTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1.DeploymentSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: paused
      type:
        scalar: boolean
    - name: progressDeadlineSeconds
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: strategy
      type:
        namedType: io.k8s.api.apps.v1.DeploymentStrategy
      default: {}
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
- name: io.k8s.api.apps.v1.DeploymentStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1.DeploymentCondition
          elementRelationship: associative
          keys:
          - type
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: unavailableReplicas
      type:
        scalar: numeric
    - name: updatedReplicas
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1.DeploymentStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.apps.v1.RollingUpdateDeployment
    - name: type
      type:
        scalar: string
- name: io.k8s.api.apps.v1.ReplicaSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1.ReplicaSetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1.ReplicaSetStatus
      default: {}
- name: io.k8s.api.apps.v1.ReplicaSetCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1.ReplicaSetSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
- name: io.k8s.api.apps.v1.ReplicaSetStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1.ReplicaSetCondition
          elementRelationship: associative
          keys:
          - type
    - name: fullyLabeledReplicas
      type:
        scalar: numeric
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.apps.v1.RollingUpdateDaemonSet
  map:
    fields:
    - name: maxSurge
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
- name: io.k8s.api.apps.v1.RollingUpdateDeployment
  map:
    fields:
    - name: maxSurge
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
- name: io.k8s.api.apps.v1.RollingUpdateStatefulSetStrategy
  map:
    fields:
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: partition
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1.StatefulSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1.StatefulSetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1.StatefulSetStatus
      default: {}
- name: io.k8s.api.apps.v1.StatefulSetCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1.StatefulSetPersistentVolumeClaimRetentionPolicy
  map:
    fields:
    - name: whenDeleted
      type:
        scalar: string
    - name: whenScaled
      type:
        scalar: string
- name: io.k8s.api.apps.v1.StatefulSetSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: persistentVolumeClaimRetentionPolicy
      type:
        namedType: io.k8s.api.apps.v1.StatefulSetPersistentVolumeClaimRetentionPolicy
    - name: podManagementPolicy
      type:
        scalar: string
    - name: replicas
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: serviceName
      type:
        scalar: string
      default: ""
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
    - name: updateStrategy
      type:
        namedType: io.k8s.api.apps.v1.StatefulSetUpdateStrategy
      default: {}
    - name: volumeClaimTemplates
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PersistentVolumeClaim
          elementRelationship: atomic
- name: io.k8s.api.apps.v1.StatefulSetStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
      default: 0
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1.StatefulSetCondition
          elementRelationship: associative
          keys:
          - type
    - name: currentReplicas
      type:
        scalar: numeric
    - name: currentRevision
      type:
        scalar: string
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
      default: 0
    - name: updateRevision
      type:
        scalar: string
    - name: updatedReplicas
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1.StatefulSetUpdateStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.apps.v1.RollingUpdateStatefulSetStrategy
    - name: type
      type:
        scalar: string
- name: io.k8s.api.apps.v1beta1.ControllerRevision
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: data
      type:
        namedType: __untyped_atomic_
      default: {}
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: revision
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.apps.v1beta1.Deployment
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1beta1.DeploymentSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1beta1.DeploymentStatus
      default: {}
- name: io.k8s.api.apps.v1beta1.DeploymentCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastUpdateTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1beta1.DeploymentSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: paused
      type:
        scalar: boolean
    - name: progressDeadlineSeconds
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: rollbackTo
      type:
        namedType: io.k8s.api.apps.v1beta1.RollbackConfig
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: strategy
      type:
        namedType: io.k8s.api.apps.v1beta1.DeploymentStrategy
      default: {}
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
- name: io.k8s.api.apps.v1beta1.DeploymentStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1beta1.DeploymentCondition
          elementRelationship: associative
          keys:
          - type
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: unavailableReplicas
      type:
        scalar: numeric
    - name: updatedReplicas
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1beta1.DeploymentStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.apps.v1beta1.RollingUpdateDeployment
    - name: type
      type:
        scalar: string
- name: io.k8s.api.apps.v1beta1.RollbackConfig
  map:
    fields:
    - name: revision
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1beta1.RollingUpdateDeployment
  map:
    fields:
    - name: maxSurge
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
- name: io.k8s.api.apps.v1beta1.RollingUpdateStatefulSetStrategy
  map:
    fields:
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: partition
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1beta1.StatefulSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1beta1.StatefulSetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1beta1.StatefulSetStatus
      default: {}
- name: io.k8s.api.apps.v1beta1.StatefulSetCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1beta1.StatefulSetPersistentVolumeClaimRetentionPolicy
  map:
    fields:
    - name: whenDeleted
      type:
        scalar: string
    - name: whenScaled
      type:
        scalar: string
- name: io.k8s.api.apps.v1beta1.StatefulSetSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: persistentVolumeClaimRetentionPolicy
      type:
        namedType: io.k8s.api.apps.v1beta1.StatefulSetPersistentVolumeClaimRetentionPolicy
    - name: podManagementPolicy
      type:
        scalar: string
    - name: replicas
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: serviceName
      type:
        scalar: string
      default: ""
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
    - name: updateStrategy
      type:
        namedType: io.k8s.api.apps.v1beta1.StatefulSetUpdateStrategy
      default: {}
    - name: volumeClaimTemplates
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PersistentVolumeClaim
          elementRelationship: atomic
- name: io.k8s.api.apps.v1beta1.StatefulSetStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
      default: 0
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1beta1.StatefulSetCondition
          elementRelationship: associative
          keys:
          - type
    - name: currentReplicas
      type:
        scalar: numeric
    - name: currentRevision
      type:
        scalar: string
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
      default: 0
    - name: updateRevision
      type:
        scalar: string
    - name: updatedReplicas
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1beta1.StatefulSetUpdateStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.apps.v1beta1.RollingUpdateStatefulSetStrategy
    - name: type
      type:
        scalar: string
- name: io.k8s.api.apps.v1beta2.ControllerRevision
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: data
      type:
        namedType: __untyped_atomic_
      default: {}
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: revision
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.apps.v1beta2.DaemonSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1beta2.DaemonSetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1beta2.DaemonSetStatus
      default: {}
- name: io.k8s.api.apps.v1beta2.DaemonSetCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1beta2.DaemonSetSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
    - name: updateStrategy
      type:
        namedType: io.k8s.api.apps.v1beta2.DaemonSetUpdateStrategy
      default: {}
- name: io.k8s.api.apps.v1beta2.DaemonSetStatus
  map:
    fields:
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1beta2.DaemonSetCondition
          elementRelationship: associative
          keys:
          - type
    - name: currentNumberScheduled
      type:
        scalar: numeric
      default: 0
    - name: desiredNumberScheduled
      type:
        scalar: numeric
      default: 0
    - name: numberAvailable
      type:
        scalar: numeric
    - name: numberMisscheduled
      type:
        scalar: numeric
      default: 0
    - name: numberReady
      type:
        scalar: numeric
      default: 0
    - name: numberUnavailable
      type:
        scalar: numeric
    - name: observedGeneration
      type:
        scalar: numeric
    - name: updatedNumberScheduled
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1beta2.DaemonSetUpdateStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.apps.v1beta2.RollingUpdateDaemonSet
    - name: type
      type:
        scalar: string
- name: io.k8s.api.apps.v1beta2.Deployment
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1beta2.DeploymentSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1beta2.DeploymentStatus
      default: {}
- name: io.k8s.api.apps.v1beta2.DeploymentCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastUpdateTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1beta2.DeploymentSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: paused
      type:
        scalar: boolean
    - name: progressDeadlineSeconds
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: strategy
      type:
        namedType: io.k8s.api.apps.v1beta2.DeploymentStrategy
      default: {}
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
- name: io.k8s.api.apps.v1beta2.DeploymentStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1beta2.DeploymentCondition
          elementRelationship: associative
          keys:
          - type
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: unavailableReplicas
      type:
        scalar: numeric
    - name: updatedReplicas
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1beta2.DeploymentStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.apps.v1beta2.RollingUpdateDeployment
    - name: type
      type:
        scalar: string
- name: io.k8s.api.apps.v1beta2.ReplicaSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1beta2.ReplicaSetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1beta2.ReplicaSetStatus
      default: {}
- name: io.k8s.api.apps.v1beta2.ReplicaSetCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1beta2.ReplicaSetSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
- name: io.k8s.api.apps.v1beta2.ReplicaSetStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1beta2.ReplicaSetCondition
          elementRelationship: associative
          keys:
          - type
    - name: fullyLabeledReplicas
      type:
        scalar: numeric
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.apps.v1beta2.RollingUpdateDaemonSet
  map:
    fields:
    - name: maxSurge
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
- name: io.k8s.api.apps.v1beta2.RollingUpdateDeployment
  map:
    fields:
    - name: maxSurge
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
- name: io.k8s.api.apps.v1beta2.RollingUpdateStatefulSetStrategy
  map:
    fields:
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: partition
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1beta2.StatefulSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.apps.v1beta2.StatefulSetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.apps.v1beta2.StatefulSetStatus
      default: {}
- name: io.k8s.api.apps.v1beta2.StatefulSetCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.apps.v1beta2.StatefulSetPersistentVolumeClaimRetentionPolicy
  map:
    fields:
    - name: whenDeleted
      type:
        scalar: string
    - name: whenScaled
      type:
        scalar: string
- name: io.k8s.api.apps.v1beta2.StatefulSetSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: persistentVolumeClaimRetentionPolicy
      type:
        namedType: io.k8s.api.apps.v1beta2.StatefulSetPersistentVolumeClaimRetentionPolicy
    - name: podManagementPolicy
      type:
        scalar: string
    - name: replicas
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: serviceName
      type:
        scalar: string
      default: ""
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
    - name: updateStrategy
      type:
        namedType: io.k8s.api.apps.v1beta2.StatefulSetUpdateStrategy
      default: {}
    - name: volumeClaimTemplates
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PersistentVolumeClaim
          elementRelationship: atomic
- name: io.k8s.api.apps.v1beta2.StatefulSetStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
      default: 0
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.apps.v1beta2.StatefulSetCondition
          elementRelationship: associative
          keys:
          - type
    - name: currentReplicas
      type:
        scalar: numeric
    - name: currentRevision
      type:
        scalar: string
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
      default: 0
    - name: updateRevision
      type:
        scalar: string
    - name: updatedReplicas
      type:
        scalar: numeric
- name: io.k8s.api.apps.v1beta2.StatefulSetUpdateStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.apps.v1beta2.RollingUpdateStatefulSetStrategy
    - name: type
      type:
        scalar: string
- name: io.k8s.api.autoscaling.v1.CrossVersionObjectReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.api.autoscaling.v1.HorizontalPodAutoscaler
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.autoscaling.v1.HorizontalPodAutoscalerSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.autoscaling.v1.HorizontalPodAutoscalerStatus
      default: {}
- name: io.k8s.api.autoscaling.v1.HorizontalPodAutoscalerSpec
  map:
    fields:
    - name: maxReplicas
      type:
        scalar: numeric
      default: 0
    - name: minReplicas
      type:
        scalar: numeric
    - name: scaleTargetRef
      type:
        namedType: io.k8s.api.autoscaling.v1.CrossVersionObjectReference
      default: {}
    - name: targetCPUUtilizationPercentage
      type:
        scalar: numeric
- name: io.k8s.api.autoscaling.v1.HorizontalPodAutoscalerStatus
  map:
    fields:
    - name: currentCPUUtilizationPercentage
      type:
        scalar: numeric
    - name: currentReplicas
      type:
        scalar: numeric
      default: 0
    - name: desiredReplicas
      type:
        scalar: numeric
      default: 0
    - name: lastScaleTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: observedGeneration
      type:
        scalar: numeric
- name: io.k8s.api.autoscaling.v2.ContainerResourceMetricSource
  map:
    fields:
    - name: container
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2.ContainerResourceMetricStatus
  map:
    fields:
    - name: container
      type:
        scalar: string
      default: ""
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricValueStatus
      default: {}
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2.CrossVersionObjectReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2.ExternalMetricSource
  map:
    fields:
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricIdentifier
      default: {}
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2.ExternalMetricStatus
  map:
    fields:
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricValueStatus
      default: {}
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricIdentifier
      default: {}
- name: io.k8s.api.autoscaling.v2.HPAScalingPolicy
  map:
    fields:
    - name: periodSeconds
      type:
        scalar: numeric
      default: 0
    - name: type
      type:
        scalar: string
      default: ""
    - name: value
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.autoscaling.v2.HPAScalingRules
  map:
    fields:
    - name: policies
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2.HPAScalingPolicy
          elementRelationship: atomic
    - name: selectPolicy
      type:
        scalar: string
    - name: stabilizationWindowSeconds
      type:
        scalar: numeric
- name: io.k8s.api.autoscaling.v2.HorizontalPodAutoscaler
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerStatus
      default: {}
- name: io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerBehavior
  map:
    fields:
    - name: scaleDown
      type:
        namedType: io.k8s.api.autoscaling.v2.HPAScalingRules
    - name: scaleUp
      type:
        namedType: io.k8s.api.autoscaling.v2.HPAScalingRules
- name: io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerSpec
  map:
    fields:
    - name: behavior
      type:
        namedType: io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerBehavior
    - name: maxReplicas
      type:
        scalar: numeric
      default: 0
    - name: metrics
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2.MetricSpec
          elementRelationship: atomic
    - name: minReplicas
      type:
        scalar: numeric
    - name: scaleTargetRef
      type:
        namedType: io.k8s.api.autoscaling.v2.CrossVersionObjectReference
      default: {}
- name: io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerCondition
          elementRelationship: associative
          keys:
          - type
    - name: currentMetrics
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2.MetricStatus
          elementRelationship: atomic
    - name: currentReplicas
      type:
        scalar: numeric
    - name: desiredReplicas
      type:
        scalar: numeric
      default: 0
    - name: lastScaleTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: observedGeneration
      type:
        scalar: numeric
- name: io.k8s.api.autoscaling.v2.MetricIdentifier
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
- name: io.k8s.api.autoscaling.v2.MetricSpec
  map:
    fields:
    - name: containerResource
      type:
        namedType: io.k8s.api.autoscaling.v2.ContainerResourceMetricSource
    - name: external
      type:
        namedType: io.k8s.api.autoscaling.v2.ExternalMetricSource
    - name: object
      type:
        namedType: io.k8s.api.autoscaling.v2.ObjectMetricSource
    - name: pods
      type:
        namedType: io.k8s.api.autoscaling.v2.PodsMetricSource
    - name: resource
      type:
        namedType: io.k8s.api.autoscaling.v2.ResourceMetricSource
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2.MetricStatus
  map:
    fields:
    - name: containerResource
      type:
        namedType: io.k8s.api.autoscaling.v2.ContainerResourceMetricStatus
    - name: external
      type:
        namedType: io.k8s.api.autoscaling.v2.ExternalMetricStatus
    - name: object
      type:
        namedType: io.k8s.api.autoscaling.v2.ObjectMetricStatus
    - name: pods
      type:
        namedType: io.k8s.api.autoscaling.v2.PodsMetricStatus
    - name: resource
      type:
        namedType: io.k8s.api.autoscaling.v2.ResourceMetricStatus
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2.MetricTarget
  map:
    fields:
    - name: averageUtilization
      type:
        scalar: numeric
    - name: averageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: type
      type:
        scalar: string
      default: ""
    - name: value
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.autoscaling.v2.MetricValueStatus
  map:
    fields:
    - name: averageUtilization
      type:
        scalar: numeric
    - name: averageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: value
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.autoscaling.v2.ObjectMetricSource
  map:
    fields:
    - name: describedObject
      type:
        namedType: io.k8s.api.autoscaling.v2.CrossVersionObjectReference
      default: {}
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricIdentifier
      default: {}
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2.ObjectMetricStatus
  map:
    fields:
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricValueStatus
      default: {}
    - name: describedObject
      type:
        namedType: io.k8s.api.autoscaling.v2.CrossVersionObjectReference
      default: {}
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricIdentifier
      default: {}
- name: io.k8s.api.autoscaling.v2.PodsMetricSource
  map:
    fields:
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricIdentifier
      default: {}
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2.PodsMetricStatus
  map:
    fields:
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricValueStatus
      default: {}
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricIdentifier
      default: {}
- name: io.k8s.api.autoscaling.v2.ResourceMetricSource
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2.ResourceMetricStatus
  map:
    fields:
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2.MetricValueStatus
      default: {}
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta1.ContainerResourceMetricSource
  map:
    fields:
    - name: container
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: targetAverageUtilization
      type:
        scalar: numeric
    - name: targetAverageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.autoscaling.v2beta1.ContainerResourceMetricStatus
  map:
    fields:
    - name: container
      type:
        scalar: string
      default: ""
    - name: currentAverageUtilization
      type:
        scalar: numeric
    - name: currentAverageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
      default: {}
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta1.CrossVersionObjectReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta1.ExternalMetricSource
  map:
    fields:
    - name: metricName
      type:
        scalar: string
      default: ""
    - name: metricSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: targetAverageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: targetValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.autoscaling.v2beta1.ExternalMetricStatus
  map:
    fields:
    - name: currentAverageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: currentValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
      default: {}
    - name: metricName
      type:
        scalar: string
      default: ""
    - name: metricSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
- name: io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscaler
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscalerSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscalerStatus
      default: {}
- name: io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscalerCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscalerSpec
  map:
    fields:
    - name: maxReplicas
      type:
        scalar: numeric
      default: 0
    - name: metrics
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2beta1.MetricSpec
          elementRelationship: atomic
    - name: minReplicas
      type:
        scalar: numeric
    - name: scaleTargetRef
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.CrossVersionObjectReference
      default: {}
- name: io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscalerStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2beta1.HorizontalPodAutoscalerCondition
          elementRelationship: atomic
    - name: currentMetrics
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2beta1.MetricStatus
          elementRelationship: atomic
    - name: currentReplicas
      type:
        scalar: numeric
      default: 0
    - name: desiredReplicas
      type:
        scalar: numeric
      default: 0
    - name: lastScaleTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: observedGeneration
      type:
        scalar: numeric
- name: io.k8s.api.autoscaling.v2beta1.MetricSpec
  map:
    fields:
    - name: containerResource
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.ContainerResourceMetricSource
    - name: external
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.ExternalMetricSource
    - name: object
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.ObjectMetricSource
    - name: pods
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.PodsMetricSource
    - name: resource
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.ResourceMetricSource
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta1.MetricStatus
  map:
    fields:
    - name: containerResource
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.ContainerResourceMetricStatus
    - name: external
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.ExternalMetricStatus
    - name: object
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.ObjectMetricStatus
    - name: pods
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.PodsMetricStatus
    - name: resource
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.ResourceMetricStatus
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta1.ObjectMetricSource
  map:
    fields:
    - name: averageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: metricName
      type:
        scalar: string
      default: ""
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.CrossVersionObjectReference
      default: {}
    - name: targetValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
      default: {}
- name: io.k8s.api.autoscaling.v2beta1.ObjectMetricStatus
  map:
    fields:
    - name: averageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: currentValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
      default: {}
    - name: metricName
      type:
        scalar: string
      default: ""
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2beta1.CrossVersionObjectReference
      default: {}
- name: io.k8s.api.autoscaling.v2beta1.PodsMetricSource
  map:
    fields:
    - name: metricName
      type:
        scalar: string
      default: ""
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: targetAverageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
      default: {}
- name: io.k8s.api.autoscaling.v2beta1.PodsMetricStatus
  map:
    fields:
    - name: currentAverageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
      default: {}
    - name: metricName
      type:
        scalar: string
      default: ""
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
- name: io.k8s.api.autoscaling.v2beta1.ResourceMetricSource
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: targetAverageUtilization
      type:
        scalar: numeric
    - name: targetAverageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.autoscaling.v2beta1.ResourceMetricStatus
  map:
    fields:
    - name: currentAverageUtilization
      type:
        scalar: numeric
    - name: currentAverageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
      default: {}
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta2.ContainerResourceMetricSource
  map:
    fields:
    - name: container
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.ContainerResourceMetricStatus
  map:
    fields:
    - name: container
      type:
        scalar: string
      default: ""
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricValueStatus
      default: {}
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta2.CrossVersionObjectReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta2.ExternalMetricSource
  map:
    fields:
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricIdentifier
      default: {}
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.ExternalMetricStatus
  map:
    fields:
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricValueStatus
      default: {}
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricIdentifier
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.HPAScalingPolicy
  map:
    fields:
    - name: periodSeconds
      type:
        scalar: numeric
      default: 0
    - name: type
      type:
        scalar: string
      default: ""
    - name: value
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.autoscaling.v2beta2.HPAScalingRules
  map:
    fields:
    - name: policies
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2beta2.HPAScalingPolicy
          elementRelationship: atomic
    - name: selectPolicy
      type:
        scalar: string
    - name: stabilizationWindowSeconds
      type:
        scalar: numeric
- name: io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscaler
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerStatus
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerBehavior
  map:
    fields:
    - name: scaleDown
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.HPAScalingRules
    - name: scaleUp
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.HPAScalingRules
- name: io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerSpec
  map:
    fields:
    - name: behavior
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerBehavior
    - name: maxReplicas
      type:
        scalar: numeric
      default: 0
    - name: metrics
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2beta2.MetricSpec
          elementRelationship: atomic
    - name: minReplicas
      type:
        scalar: numeric
    - name: scaleTargetRef
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.CrossVersionObjectReference
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2beta2.HorizontalPodAutoscalerCondition
          elementRelationship: atomic
    - name: currentMetrics
      type:
        list:
          elementType:
            namedType: io.k8s.api.autoscaling.v2beta2.MetricStatus
          elementRelationship: atomic
    - name: currentReplicas
      type:
        scalar: numeric
      default: 0
    - name: desiredReplicas
      type:
        scalar: numeric
      default: 0
    - name: lastScaleTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: observedGeneration
      type:
        scalar: numeric
- name: io.k8s.api.autoscaling.v2beta2.MetricIdentifier
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
- name: io.k8s.api.autoscaling.v2beta2.MetricSpec
  map:
    fields:
    - name: containerResource
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.ContainerResourceMetricSource
    - name: external
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.ExternalMetricSource
    - name: object
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.ObjectMetricSource
    - name: pods
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.PodsMetricSource
    - name: resource
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.ResourceMetricSource
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta2.MetricStatus
  map:
    fields:
    - name: containerResource
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.ContainerResourceMetricStatus
    - name: external
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.ExternalMetricStatus
    - name: object
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.ObjectMetricStatus
    - name: pods
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.PodsMetricStatus
    - name: resource
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.ResourceMetricStatus
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.autoscaling.v2beta2.MetricTarget
  map:
    fields:
    - name: averageUtilization
      type:
        scalar: numeric
    - name: averageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: type
      type:
        scalar: string
      default: ""
    - name: value
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.autoscaling.v2beta2.MetricValueStatus
  map:
    fields:
    - name: averageUtilization
      type:
        scalar: numeric
    - name: averageValue
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: value
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.autoscaling.v2beta2.ObjectMetricSource
  map:
    fields:
    - name: describedObject
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.CrossVersionObjectReference
      default: {}
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricIdentifier
      default: {}
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.ObjectMetricStatus
  map:
    fields:
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricValueStatus
      default: {}
    - name: describedObject
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.CrossVersionObjectReference
      default: {}
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricIdentifier
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.PodsMetricSource
  map:
    fields:
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricIdentifier
      default: {}
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.PodsMetricStatus
  map:
    fields:
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricValueStatus
      default: {}
    - name: metric
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricIdentifier
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.ResourceMetricSource
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: target
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricTarget
      default: {}
- name: io.k8s.api.autoscaling.v2beta2.ResourceMetricStatus
  map:
    fields:
    - name: current
      type:
        namedType: io.k8s.api.autoscaling.v2beta2.MetricValueStatus
      default: {}
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.batch.v1.CronJob
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.batch.v1.CronJobSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.batch.v1.CronJobStatus
      default: {}
- name: io.k8s.api.batch.v1.CronJobSpec
  map:
    fields:
    - name: concurrencyPolicy
      type:
        scalar: string
    - name: failedJobsHistoryLimit
      type:
        scalar: numeric
    - name: jobTemplate
      type:
        namedType: io.k8s.api.batch.v1.JobTemplateSpec
      default: {}
    - name: schedule
      type:
        scalar: string
      default: ""
    - name: startingDeadlineSeconds
      type:
        scalar: numeric
    - name: successfulJobsHistoryLimit
      type:
        scalar: numeric
    - name: suspend
      type:
        scalar: boolean
    - name: timeZone
      type:
        scalar: string
- name: io.k8s.api.batch.v1.CronJobStatus
  map:
    fields:
    - name: active
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ObjectReference
          elementRelationship: atomic
    - name: lastScheduleTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: lastSuccessfulTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.api.batch.v1.Job
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.batch.v1.JobSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.batch.v1.JobStatus
      default: {}
- name: io.k8s.api.batch.v1.JobCondition
  map:
    fields:
    - name: lastProbeTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.batch.v1.JobSpec
  map:
    fields:
    - name: activeDeadlineSeconds
      type:
        scalar: numeric
    - name: backoffLimit
      type:
        scalar: numeric
    - name: completionMode
      type:
        scalar: string
    - name: completions
      type:
        scalar: numeric
    - name: manualSelector
      type:
        scalar: boolean
    - name: parallelism
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: suspend
      type:
        scalar: boolean
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
    - name: ttlSecondsAfterFinished
      type:
        scalar: numeric
- name: io.k8s.api.batch.v1.JobStatus
  map:
    fields:
    - name: active
      type:
        scalar: numeric
    - name: completedIndexes
      type:
        scalar: string
    - name: completionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.batch.v1.JobCondition
          elementRelationship: atomic
    - name: failed
      type:
        scalar: numeric
    - name: ready
      type:
        scalar: numeric
    - name: startTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: succeeded
      type:
        scalar: numeric
    - name: uncountedTerminatedPods
      type:
        namedType: io.k8s.api.batch.v1.UncountedTerminatedPods
- name: io.k8s.api.batch.v1.JobTemplateSpec
  map:
    fields:
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.batch.v1.JobSpec
      default: {}
- name: io.k8s.api.batch.v1.UncountedTerminatedPods
  map:
    fields:
    - name: failed
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: succeeded
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: io.k8s.api.batch.v1beta1.CronJob
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.batch.v1beta1.CronJobSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.batch.v1beta1.CronJobStatus
      default: {}
- name: io.k8s.api.batch.v1beta1.CronJobSpec
  map:
    fields:
    - name: concurrencyPolicy
      type:
        scalar: string
    - name: failedJobsHistoryLimit
      type:
        scalar: numeric
    - name: jobTemplate
      type:
        namedType: io.k8s.api.batch.v1beta1.JobTemplateSpec
      default: {}
    - name: schedule
      type:
        scalar: string
      default: ""
    - name: startingDeadlineSeconds
      type:
        scalar: numeric
    - name: successfulJobsHistoryLimit
      type:
        scalar: numeric
    - name: suspend
      type:
        scalar: boolean
    - name: timeZone
      type:
        scalar: string
- name: io.k8s.api.batch.v1beta1.CronJobStatus
  map:
    fields:
    - name: active
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ObjectReference
          elementRelationship: atomic
    - name: lastScheduleTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: lastSuccessfulTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.api.batch.v1beta1.JobTemplateSpec
  map:
    fields:
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.batch.v1.JobSpec
      default: {}
- name: io.k8s.api.certificates.v1.CertificateSigningRequest
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.certificates.v1.CertificateSigningRequestSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.certificates.v1.CertificateSigningRequestStatus
      default: {}
- name: io.k8s.api.certificates.v1.CertificateSigningRequestCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastUpdateTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.certificates.v1.CertificateSigningRequestSpec
  map:
    fields:
    - name: expirationSeconds
      type:
        scalar: numeric
    - name: extra
      type:
        map:
          elementType:
            list:
              elementType:
                scalar: string
              elementRelationship: atomic
    - name: groups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: request
      type:
        scalar: string
    - name: signerName
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
    - name: usages
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: username
      type:
        scalar: string
- name: io.k8s.api.certificates.v1.CertificateSigningRequestStatus
  map:
    fields:
    - name: certificate
      type:
        scalar: string
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.certificates.v1.CertificateSigningRequestCondition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.certificates.v1beta1.CertificateSigningRequest
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.certificates.v1beta1.CertificateSigningRequestSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.certificates.v1beta1.CertificateSigningRequestStatus
      default: {}
- name: io.k8s.api.certificates.v1beta1.CertificateSigningRequestCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastUpdateTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.certificates.v1beta1.CertificateSigningRequestSpec
  map:
    fields:
    - name: expirationSeconds
      type:
        scalar: numeric
    - name: extra
      type:
        map:
          elementType:
            list:
              elementType:
                scalar: string
              elementRelationship: atomic
    - name: groups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: request
      type:
        scalar: string
    - name: signerName
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
    - name: usages
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: username
      type:
        scalar: string
- name: io.k8s.api.certificates.v1beta1.CertificateSigningRequestStatus
  map:
    fields:
    - name: certificate
      type:
        scalar: string
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.certificates.v1beta1.CertificateSigningRequestCondition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.coordination.v1.Lease
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.coordination.v1.LeaseSpec
      default: {}
- name: io.k8s.api.coordination.v1.LeaseSpec
  map:
    fields:
    - name: acquireTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
    - name: holderIdentity
      type:
        scalar: string
    - name: leaseDurationSeconds
      type:
        scalar: numeric
    - name: leaseTransitions
      type:
        scalar: numeric
    - name: renewTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
- name: io.k8s.api.coordination.v1beta1.Lease
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.coordination.v1beta1.LeaseSpec
      default: {}
- name: io.k8s.api.coordination.v1beta1.LeaseSpec
  map:
    fields:
    - name: acquireTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
    - name: holderIdentity
      type:
        scalar: string
    - name: leaseDurationSeconds
      type:
        scalar: numeric
    - name: leaseTransitions
      type:
        scalar: numeric
    - name: renewTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
- name: io.k8s.api.core.v1.AWSElasticBlockStoreVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: partition
      type:
        scalar: numeric
    - name: readOnly
      type:
        scalar: boolean
    - name: volumeID
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.Affinity
  map:
    fields:
    - name: nodeAffinity
      type:
        namedType: io.k8s.api.core.v1.NodeAffinity
    - name: podAffinity
      type:
        namedType: io.k8s.api.core.v1.PodAffinity
    - name: podAntiAffinity
      type:
        namedType: io.k8s.api.core.v1.PodAntiAffinity
- name: io.k8s.api.core.v1.AttachedVolume
  map:
    fields:
    - name: devicePath
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.AzureDiskVolumeSource
  map:
    fields:
    - name: cachingMode
      type:
        scalar: string
    - name: diskName
      type:
        scalar: string
      default: ""
    - name: diskURI
      type:
        scalar: string
      default: ""
    - name: fsType
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.AzureFilePersistentVolumeSource
  map:
    fields:
    - name: readOnly
      type:
        scalar: boolean
    - name: secretName
      type:
        scalar: string
      default: ""
    - name: secretNamespace
      type:
        scalar: string
    - name: shareName
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.AzureFileVolumeSource
  map:
    fields:
    - name: readOnly
      type:
        scalar: boolean
    - name: secretName
      type:
        scalar: string
      default: ""
    - name: shareName
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.CSIPersistentVolumeSource
  map:
    fields:
    - name: controllerExpandSecretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
    - name: controllerPublishSecretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
    - name: driver
      type:
        scalar: string
      default: ""
    - name: fsType
      type:
        scalar: string
    - name: nodePublishSecretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
    - name: nodeStageSecretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
    - name: readOnly
      type:
        scalar: boolean
    - name: volumeAttributes
      type:
        map:
          elementType:
            scalar: string
    - name: volumeHandle
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.CSIVolumeSource
  map:
    fields:
    - name: driver
      type:
        scalar: string
      default: ""
    - name: fsType
      type:
        scalar: string
    - name: nodePublishSecretRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
    - name: readOnly
      type:
        scalar: boolean
    - name: volumeAttributes
      type:
        map:
          elementType:
            scalar: string
- name: io.k8s.api.core.v1.Capabilities
  map:
    fields:
    - name: add
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: drop
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.CephFSPersistentVolumeSource
  map:
    fields:
    - name: monitors
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: path
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretFile
      type:
        scalar: string
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
    - name: user
      type:
        scalar: string
- name: io.k8s.api.core.v1.CephFSVolumeSource
  map:
    fields:
    - name: monitors
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: path
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretFile
      type:
        scalar: string
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
    - name: user
      type:
        scalar: string
- name: io.k8s.api.core.v1.CinderPersistentVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
    - name: volumeID
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.CinderVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
    - name: volumeID
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.ClientIPConfig
  map:
    fields:
    - name: timeoutSeconds
      type:
        scalar: numeric
- name: io.k8s.api.core.v1.ComponentCondition
  map:
    fields:
    - name: error
      type:
        scalar: string
    - name: message
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.ComponentStatus
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ComponentCondition
          elementRelationship: associative
          keys:
          - type
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
- name: io.k8s.api.core.v1.ConfigMap
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: binaryData
      type:
        map:
          elementType:
            scalar: string
    - name: data
      type:
        map:
          elementType:
            scalar: string
    - name: immutable
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
- name: io.k8s.api.core.v1.ConfigMapEnvSource
  map:
    fields:
    - name: name
      type:
        scalar: string
    - name: optional
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.ConfigMapKeySelector
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
    - name: optional
      type:
        scalar: boolean
    elementRelationship: atomic
- name: io.k8s.api.core.v1.ConfigMapNodeConfigSource
  map:
    fields:
    - name: kubeletConfigKey
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
    - name: resourceVersion
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.api.core.v1.ConfigMapProjection
  map:
    fields:
    - name: items
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.KeyToPath
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: optional
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.ConfigMapVolumeSource
  map:
    fields:
    - name: defaultMode
      type:
        scalar: numeric
    - name: items
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.KeyToPath
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: optional
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.Container
  map:
    fields:
    - name: args
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: command
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: env
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.EnvVar
          elementRelationship: associative
          keys:
          - name
    - name: envFrom
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.EnvFromSource
          elementRelationship: atomic
    - name: image
      type:
        scalar: string
    - name: imagePullPolicy
      type:
        scalar: string
    - name: lifecycle
      type:
        namedType: io.k8s.api.core.v1.Lifecycle
    - name: livenessProbe
      type:
        namedType: io.k8s.api.core.v1.Probe
    - name: name
      type:
        scalar: string
      default: ""
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ContainerPort
          elementRelationship: associative
          keys:
          - containerPort
          - protocol
    - name: readinessProbe
      type:
        namedType: io.k8s.api.core.v1.Probe
    - name: resources
      type:
        namedType: io.k8s.api.core.v1.ResourceRequirements
      default: {}
    - name: securityContext
      type:
        namedType: io.k8s.api.core.v1.SecurityContext
    - name: startupProbe
      type:
        namedType: io.k8s.api.core.v1.Probe
    - name: stdin
      type:
        scalar: boolean
    - name: stdinOnce
      type:
        scalar: boolean
    - name: terminationMessagePath
      type:
        scalar: string
    - name: terminationMessagePolicy
      type:
        scalar: string
    - name: tty
      type:
        scalar: boolean
    - name: volumeDevices
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.VolumeDevice
          elementRelationship: associative
          keys:
          - devicePath
    - name: volumeMounts
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.VolumeMount
          elementRelationship: associative
          keys:
          - mountPath
    - name: workingDir
      type:
        scalar: string
- name: io.k8s.api.core.v1.ContainerImage
  map:
    fields:
    - name: names
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: sizeBytes
      type:
        scalar: numeric
- name: io.k8s.api.core.v1.ContainerPort
  map:
    fields:
    - name: containerPort
      type:
        scalar: numeric
      default: 0
    - name: hostIP
      type:
        scalar: string
    - name: hostPort
      type:
        scalar: numeric
    - name: name
      type:
        scalar: string
    - name: protocol
      type:
        scalar: string
      default: TCP
- name: io.k8s.api.core.v1.ContainerState
  map:
    fields:
    - name: running
      type:
        namedType: io.k8s.api.core.v1.ContainerStateRunning
    - name: terminated
      type:
        namedType: io.k8s.api.core.v1.ContainerStateTerminated
    - name: waiting
      type:
        namedType: io.k8s.api.core.v1.ContainerStateWaiting
- name: io.k8s.api.core.v1.ContainerStateRunning
  map:
    fields:
    - name: startedAt
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
- name: io.k8s.api.core.v1.ContainerStateTerminated
  map:
    fields:
    - name: containerID
      type:
        scalar: string
    - name: exitCode
      type:
        scalar: numeric
      default: 0
    - name: finishedAt
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: signal
      type:
        scalar: numeric
    - name: startedAt
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
- name: io.k8s.api.core.v1.ContainerStateWaiting
  map:
    fields:
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
- name: io.k8s.api.core.v1.ContainerStatus
  map:
    fields:
    - name: containerID
      type:
        scalar: string
    - name: image
      type:
        scalar: string
      default: ""
    - name: imageID
      type:
        scalar: string
      default: ""
    - name: lastState
      type:
        namedType: io.k8s.api.core.v1.ContainerState
      default: {}
    - name: name
      type:
        scalar: string
      default: ""
    - name: ready
      type:
        scalar: boolean
      default: false
    - name: restartCount
      type:
        scalar: numeric
      default: 0
    - name: started
      type:
        scalar: boolean
    - name: state
      type:
        namedType: io.k8s.api.core.v1.ContainerState
      default: {}
- name: io.k8s.api.core.v1.DaemonEndpoint
  map:
    fields:
    - name: Port
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.core.v1.DownwardAPIProjection
  map:
    fields:
    - name: items
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.DownwardAPIVolumeFile
          elementRelationship: atomic
- name: io.k8s.api.core.v1.DownwardAPIVolumeFile
  map:
    fields:
    - name: fieldRef
      type:
        namedType: io.k8s.api.core.v1.ObjectFieldSelector
    - name: mode
      type:
        scalar: numeric
    - name: path
      type:
        scalar: string
      default: ""
    - name: resourceFieldRef
      type:
        namedType: io.k8s.api.core.v1.ResourceFieldSelector
- name: io.k8s.api.core.v1.DownwardAPIVolumeSource
  map:
    fields:
    - name: defaultMode
      type:
        scalar: numeric
    - name: items
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.DownwardAPIVolumeFile
          elementRelationship: atomic
- name: io.k8s.api.core.v1.EmptyDirVolumeSource
  map:
    fields:
    - name: medium
      type:
        scalar: string
    - name: sizeLimit
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.core.v1.EndpointAddress
  map:
    fields:
    - name: hostname
      type:
        scalar: string
    - name: ip
      type:
        scalar: string
      default: ""
    - name: nodeName
      type:
        scalar: string
    - name: targetRef
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    elementRelationship: atomic
- name: io.k8s.api.core.v1.EndpointPort
  map:
    fields:
    - name: appProtocol
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
      default: 0
    - name: protocol
      type:
        scalar: string
    elementRelationship: atomic
- name: io.k8s.api.core.v1.EndpointSubset
  map:
    fields:
    - name: addresses
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.EndpointAddress
          elementRelationship: atomic
    - name: notReadyAddresses
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.EndpointAddress
          elementRelationship: atomic
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.EndpointPort
          elementRelationship: atomic
- name: io.k8s.api.core.v1.Endpoints
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: subsets
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.EndpointSubset
          elementRelationship: atomic
- name: io.k8s.api.core.v1.EnvFromSource
  map:
    fields:
    - name: configMapRef
      type:
        namedType: io.k8s.api.core.v1.ConfigMapEnvSource
    - name: prefix
      type:
        scalar: string
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.SecretEnvSource
- name: io.k8s.api.core.v1.EnvVar
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: value
      type:
        scalar: string
    - name: valueFrom
      type:
        namedType: io.k8s.api.core.v1.EnvVarSource
- name: io.k8s.api.core.v1.EnvVarSource
  map:
    fields:
    - name: configMapKeyRef
      type:
        namedType: io.k8s.api.core.v1.ConfigMapKeySelector
    - name: fieldRef
      type:
        namedType: io.k8s.api.core.v1.ObjectFieldSelector
    - name: resourceFieldRef
      type:
        namedType: io.k8s.api.core.v1.ResourceFieldSelector
    - name: secretKeyRef
      type:
        namedType: io.k8s.api.core.v1.SecretKeySelector
- name: io.k8s.api.core.v1.EphemeralContainer
  map:
    fields:
    - name: args
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: command
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: env
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.EnvVar
          elementRelationship: associative
          keys:
          - name
    - name: envFrom
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.EnvFromSource
          elementRelationship: atomic
    - name: image
      type:
        scalar: string
    - name: imagePullPolicy
      type:
        scalar: string
    - name: lifecycle
      type:
        namedType: io.k8s.api.core.v1.Lifecycle
    - name: livenessProbe
      type:
        namedType: io.k8s.api.core.v1.Probe
    - name: name
      type:
        scalar: string
      default: ""
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ContainerPort
          elementRelationship: associative
          keys:
          - containerPort
          - protocol
    - name: readinessProbe
      type:
        namedType: io.k8s.api.core.v1.Probe
    - name: resources
      type:
        namedType: io.k8s.api.core.v1.ResourceRequirements
      default: {}
    - name: securityContext
      type:
        namedType: io.k8s.api.core.v1.SecurityContext
    - name: startupProbe
      type:
        namedType: io.k8s.api.core.v1.Probe
    - name: stdin
      type:
        scalar: boolean
    - name: stdinOnce
      type:
        scalar: boolean
    - name: targetContainerName
      type:
        scalar: string
    - name: terminationMessagePath
      type:
        scalar: string
    - name: terminationMessagePolicy
      type:
        scalar: string
    - name: tty
      type:
        scalar: boolean
    - name: volumeDevices
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.VolumeDevice
          elementRelationship: associative
          keys:
          - devicePath
    - name: volumeMounts
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.VolumeMount
          elementRelationship: associative
          keys:
          - mountPath
    - name: workingDir
      type:
        scalar: string
- name: io.k8s.api.core.v1.EphemeralVolumeSource
  map:
    fields:
    - name: volumeClaimTemplate
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeClaimTemplate
- name: io.k8s.api.core.v1.Event
  map:
    fields:
    - name: action
      type:
        scalar: string
    - name: apiVersion
      type:
        scalar: string
    - name: count
      type:
        scalar: numeric
    - name: eventTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
      default: {}
    - name: firstTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: involvedObject
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
      default: {}
    - name: kind
      type:
        scalar: string
    - name: lastTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: reason
      type:
        scalar: string
    - name: related
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    - name: reportingComponent
      type:
        scalar: string
      default: ""
    - name: reportingInstance
      type:
        scalar: string
      default: ""
    - name: series
      type:
        namedType: io.k8s.api.core.v1.EventSeries
    - name: source
      type:
        namedType: io.k8s.api.core.v1.EventSource
      default: {}
    - name: type
      type:
        scalar: string
- name: io.k8s.api.core.v1.EventSeries
  map:
    fields:
    - name: count
      type:
        scalar: numeric
    - name: lastObservedTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
      default: {}
- name: io.k8s.api.core.v1.EventSource
  map:
    fields:
    - name: component
      type:
        scalar: string
    - name: host
      type:
        scalar: string
- name: io.k8s.api.core.v1.ExecAction
  map:
    fields:
    - name: command
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.FCVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: lun
      type:
        scalar: numeric
    - name: readOnly
      type:
        scalar: boolean
    - name: targetWWNs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: wwids
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.FlexPersistentVolumeSource
  map:
    fields:
    - name: driver
      type:
        scalar: string
      default: ""
    - name: fsType
      type:
        scalar: string
    - name: options
      type:
        map:
          elementType:
            scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
- name: io.k8s.api.core.v1.FlexVolumeSource
  map:
    fields:
    - name: driver
      type:
        scalar: string
      default: ""
    - name: fsType
      type:
        scalar: string
    - name: options
      type:
        map:
          elementType:
            scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
- name: io.k8s.api.core.v1.FlockerVolumeSource
  map:
    fields:
    - name: datasetName
      type:
        scalar: string
    - name: datasetUUID
      type:
        scalar: string
- name: io.k8s.api.core.v1.GCEPersistentDiskVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: partition
      type:
        scalar: numeric
    - name: pdName
      type:
        scalar: string
      default: ""
    - name: readOnly
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.GRPCAction
  map:
    fields:
    - name: port
      type:
        scalar: numeric
      default: 0
    - name: service
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.GitRepoVolumeSource
  map:
    fields:
    - name: directory
      type:
        scalar: string
    - name: repository
      type:
        scalar: string
      default: ""
    - name: revision
      type:
        scalar: string
- name: io.k8s.api.core.v1.GlusterfsPersistentVolumeSource
  map:
    fields:
    - name: endpoints
      type:
        scalar: string
      default: ""
    - name: endpointsNamespace
      type:
        scalar: string
    - name: path
      type:
        scalar: string
      default: ""
    - name: readOnly
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.GlusterfsVolumeSource
  map:
    fields:
    - name: endpoints
      type:
        scalar: string
      default: ""
    - name: path
      type:
        scalar: string
      default: ""
    - name: readOnly
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.HTTPGetAction
  map:
    fields:
    - name: host
      type:
        scalar: string
    - name: httpHeaders
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.HTTPHeader
          elementRelationship: atomic
    - name: path
      type:
        scalar: string
    - name: port
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
      default: {}
    - name: scheme
      type:
        scalar: string
- name: io.k8s.api.core.v1.HTTPHeader
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: value
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.HostAlias
  map:
    fields:
    - name: hostnames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: ip
      type:
        scalar: string
- name: io.k8s.api.core.v1.HostPathVolumeSource
  map:
    fields:
    - name: path
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
- name: io.k8s.api.core.v1.ISCSIPersistentVolumeSource
  map:
    fields:
    - name: chapAuthDiscovery
      type:
        scalar: boolean
    - name: chapAuthSession
      type:
        scalar: boolean
    - name: fsType
      type:
        scalar: string
    - name: initiatorName
      type:
        scalar: string
    - name: iqn
      type:
        scalar: string
      default: ""
    - name: iscsiInterface
      type:
        scalar: string
    - name: lun
      type:
        scalar: numeric
      default: 0
    - name: portals
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
    - name: targetPortal
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.ISCSIVolumeSource
  map:
    fields:
    - name: chapAuthDiscovery
      type:
        scalar: boolean
    - name: chapAuthSession
      type:
        scalar: boolean
    - name: fsType
      type:
        scalar: string
    - name: initiatorName
      type:
        scalar: string
    - name: iqn
      type:
        scalar: string
      default: ""
    - name: iscsiInterface
      type:
        scalar: string
    - name: lun
      type:
        scalar: numeric
      default: 0
    - name: portals
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
    - name: targetPortal
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.KeyToPath
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: mode
      type:
        scalar: numeric
    - name: path
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.Lifecycle
  map:
    fields:
    - name: postStart
      type:
        namedType: io.k8s.api.core.v1.LifecycleHandler
    - name: preStop
      type:
        namedType: io.k8s.api.core.v1.LifecycleHandler
- name: io.k8s.api.core.v1.LifecycleHandler
  map:
    fields:
    - name: exec
      type:
        namedType: io.k8s.api.core.v1.ExecAction
    - name: httpGet
      type:
        namedType: io.k8s.api.core.v1.HTTPGetAction
    - name: tcpSocket
      type:
        namedType: io.k8s.api.core.v1.TCPSocketAction
- name: io.k8s.api.core.v1.LimitRange
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.LimitRangeSpec
      default: {}
- name: io.k8s.api.core.v1.LimitRangeItem
  map:
    fields:
    - name: default
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: defaultRequest
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: max
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: maxLimitRequestRatio
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: min
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.LimitRangeSpec
  map:
    fields:
    - name: limits
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.LimitRangeItem
          elementRelationship: atomic
- name: io.k8s.api.core.v1.LoadBalancerIngress
  map:
    fields:
    - name: hostname
      type:
        scalar: string
    - name: ip
      type:
        scalar: string
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PortStatus
          elementRelationship: atomic
- name: io.k8s.api.core.v1.LoadBalancerStatus
  map:
    fields:
    - name: ingress
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.LoadBalancerIngress
          elementRelationship: atomic
- name: io.k8s.api.core.v1.LocalObjectReference
  map:
    fields:
    - name: name
      type:
        scalar: string
    elementRelationship: atomic
- name: io.k8s.api.core.v1.LocalVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: path
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.NFSVolumeSource
  map:
    fields:
    - name: path
      type:
        scalar: string
      default: ""
    - name: readOnly
      type:
        scalar: boolean
    - name: server
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.Namespace
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.NamespaceSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.core.v1.NamespaceStatus
      default: {}
- name: io.k8s.api.core.v1.NamespaceCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.NamespaceSpec
  map:
    fields:
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.NamespaceStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.NamespaceCondition
          elementRelationship: associative
          keys:
          - type
    - name: phase
      type:
        scalar: string
- name: io.k8s.api.core.v1.Node
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.NodeSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.core.v1.NodeStatus
      default: {}
- name: io.k8s.api.core.v1.NodeAddress
  map:
    fields:
    - name: address
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.NodeAffinity
  map:
    fields:
    - name: preferredDuringSchedulingIgnoredDuringExecution
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PreferredSchedulingTerm
          elementRelationship: atomic
    - name: requiredDuringSchedulingIgnoredDuringExecution
      type:
        namedType: io.k8s.api.core.v1.NodeSelector
- name: io.k8s.api.core.v1.NodeCondition
  map:
    fields:
    - name: lastHeartbeatTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.NodeConfigSource
  map:
    fields:
    - name: configMap
      type:
        namedType: io.k8s.api.core.v1.ConfigMapNodeConfigSource
- name: io.k8s.api.core.v1.NodeConfigStatus
  map:
    fields:
    - name: active
      type:
        namedType: io.k8s.api.core.v1.NodeConfigSource
    - name: assigned
      type:
        namedType: io.k8s.api.core.v1.NodeConfigSource
    - name: error
      type:
        scalar: string
    - name: lastKnownGood
      type:
        namedType: io.k8s.api.core.v1.NodeConfigSource
- name: io.k8s.api.core.v1.NodeDaemonEndpoints
  map:
    fields:
    - name: kubeletEndpoint
      type:
        namedType: io.k8s.api.core.v1.DaemonEndpoint
      default: {}
- name: io.k8s.api.core.v1.NodeSelector
  map:
    fields:
    - name: nodeSelectorTerms
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.NodeSelectorTerm
          elementRelationship: atomic
    elementRelationship: atomic
- name: io.k8s.api.core.v1.NodeSelectorRequirement
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: operator
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.NodeSelectorTerm
  map:
    fields:
    - name: matchExpressions
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.NodeSelectorRequirement
          elementRelationship: atomic
    - name: matchFields
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.NodeSelectorRequirement
          elementRelationship: atomic
    elementRelationship: atomic
- name: io.k8s.api.core.v1.NodeSpec
  map:
    fields:
    - name: configSource
      type:
        namedType: io.k8s.api.core.v1.NodeConfigSource
    - name: externalID
      type:
        scalar: string
    - name: podCIDR
      type:
        scalar: string
    - name: podCIDRs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: providerID
      type:
        scalar: string
    - name: taints
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Taint
          elementRelationship: atomic
    - name: unschedulable
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.NodeStatus
  map:
    fields:
    - name: addresses
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.NodeAddress
          elementRelationship: associative
          keys:
          - type
    - name: allocatable
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: capacity
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.NodeCondition
          elementRelationship: associative
          keys:
          - type
    - name: config
      type:
        namedType: io.k8s.api.core.v1.NodeConfigStatus
    - name: daemonEndpoints
      type:
        namedType: io.k8s.api.core.v1.NodeDaemonEndpoints
      default: {}
    - name: images
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ContainerImage
          elementRelationship: atomic
    - name: nodeInfo
      type:
        namedType: io.k8s.api.core.v1.NodeSystemInfo
      default: {}
    - name: phase
      type:
        scalar: string
    - name: volumesAttached
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.AttachedVolume
          elementRelationship: atomic
    - name: volumesInUse
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.NodeSystemInfo
  map:
    fields:
    - name: architecture
      type:
        scalar: string
      default: ""
    - name: bootID
      type:
        scalar: string
      default: ""
    - name: containerRuntimeVersion
      type:
        scalar: string
      default: ""
    - name: kernelVersion
      type:
        scalar: string
      default: ""
    - name: kubeProxyVersion
      type:
        scalar: string
      default: ""
    - name: kubeletVersion
      type:
        scalar: string
      default: ""
    - name: machineID
      type:
        scalar: string
      default: ""
    - name: operatingSystem
      type:
        scalar: string
      default: ""
    - name: osImage
      type:
        scalar: string
      default: ""
    - name: systemUUID
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.ObjectFieldSelector
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldPath
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.api.core.v1.ObjectReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldPath
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: resourceVersion
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
    elementRelationship: atomic
- name: io.k8s.api.core.v1.PersistentVolume
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeStatus
      default: {}
- name: io.k8s.api.core.v1.PersistentVolumeClaim
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeClaimSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeClaimStatus
      default: {}
- name: io.k8s.api.core.v1.PersistentVolumeClaimCondition
  map:
    fields:
    - name: lastProbeTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.PersistentVolumeClaimSpec
  map:
    fields:
    - name: accessModes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: dataSource
      type:
        namedType: io.k8s.api.core.v1.TypedLocalObjectReference
    - name: dataSourceRef
      type:
        namedType: io.k8s.api.core.v1.TypedLocalObjectReference
    - name: resources
      type:
        namedType: io.k8s.api.core.v1.ResourceRequirements
      default: {}
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: storageClassName
      type:
        scalar: string
    - name: volumeMode
      type:
        scalar: string
    - name: volumeName
      type:
        scalar: string
- name: io.k8s.api.core.v1.PersistentVolumeClaimStatus
  map:
    fields:
    - name: accessModes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: allocatedResources
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: capacity
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PersistentVolumeClaimCondition
          elementRelationship: associative
          keys:
          - type
    - name: phase
      type:
        scalar: string
    - name: resizeStatus
      type:
        scalar: string
- name: io.k8s.api.core.v1.PersistentVolumeClaimTemplate
  map:
    fields:
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeClaimSpec
      default: {}
- name: io.k8s.api.core.v1.PersistentVolumeClaimVolumeSource
  map:
    fields:
    - name: claimName
      type:
        scalar: string
      default: ""
    - name: readOnly
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.PersistentVolumeSpec
  map:
    fields:
    - name: accessModes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: awsElasticBlockStore
      type:
        namedType: io.k8s.api.core.v1.AWSElasticBlockStoreVolumeSource
    - name: azureDisk
      type:
        namedType: io.k8s.api.core.v1.AzureDiskVolumeSource
    - name: azureFile
      type:
        namedType: io.k8s.api.core.v1.AzureFilePersistentVolumeSource
    - name: capacity
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: cephfs
      type:
        namedType: io.k8s.api.core.v1.CephFSPersistentVolumeSource
    - name: cinder
      type:
        namedType: io.k8s.api.core.v1.CinderPersistentVolumeSource
    - name: claimRef
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    - name: csi
      type:
        namedType: io.k8s.api.core.v1.CSIPersistentVolumeSource
    - name: fc
      type:
        namedType: io.k8s.api.core.v1.FCVolumeSource
    - name: flexVolume
      type:
        namedType: io.k8s.api.core.v1.FlexPersistentVolumeSource
    - name: flocker
      type:
        namedType: io.k8s.api.core.v1.FlockerVolumeSource
    - name: gcePersistentDisk
      type:
        namedType: io.k8s.api.core.v1.GCEPersistentDiskVolumeSource
    - name: glusterfs
      type:
        namedType: io.k8s.api.core.v1.GlusterfsPersistentVolumeSource
    - name: hostPath
      type:
        namedType: io.k8s.api.core.v1.HostPathVolumeSource
    - name: iscsi
      type:
        namedType: io.k8s.api.core.v1.ISCSIPersistentVolumeSource
    - name: local
      type:
        namedType: io.k8s.api.core.v1.LocalVolumeSource
    - name: mountOptions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: nfs
      type:
        namedType: io.k8s.api.core.v1.NFSVolumeSource
    - name: nodeAffinity
      type:
        namedType: io.k8s.api.core.v1.VolumeNodeAffinity
    - name: persistentVolumeReclaimPolicy
      type:
        scalar: string
    - name: photonPersistentDisk
      type:
        namedType: io.k8s.api.core.v1.PhotonPersistentDiskVolumeSource
    - name: portworxVolume
      type:
        namedType: io.k8s.api.core.v1.PortworxVolumeSource
    - name: quobyte
      type:
        namedType: io.k8s.api.core.v1.QuobyteVolumeSource
    - name: rbd
      type:
        namedType: io.k8s.api.core.v1.RBDPersistentVolumeSource
    - name: scaleIO
      type:
        namedType: io.k8s.api.core.v1.ScaleIOPersistentVolumeSource
    - name: storageClassName
      type:
        scalar: string
    - name: storageos
      type:
        namedType: io.k8s.api.core.v1.StorageOSPersistentVolumeSource
    - name: volumeMode
      type:
        scalar: string
    - name: vsphereVolume
      type:
        namedType: io.k8s.api.core.v1.VsphereVirtualDiskVolumeSource
- name: io.k8s.api.core.v1.PersistentVolumeStatus
  map:
    fields:
    - name: message
      type:
        scalar: string
    - name: phase
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
- name: io.k8s.api.core.v1.PhotonPersistentDiskVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: pdID
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.Pod
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.PodSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.core.v1.PodStatus
      default: {}
- name: io.k8s.api.core.v1.PodAffinity
  map:
    fields:
    - name: preferredDuringSchedulingIgnoredDuringExecution
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.WeightedPodAffinityTerm
          elementRelationship: atomic
    - name: requiredDuringSchedulingIgnoredDuringExecution
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PodAffinityTerm
          elementRelationship: atomic
- name: io.k8s.api.core.v1.PodAffinityTerm
  map:
    fields:
    - name: labelSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: namespaceSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: namespaces
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: topologyKey
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.PodAntiAffinity
  map:
    fields:
    - name: preferredDuringSchedulingIgnoredDuringExecution
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.WeightedPodAffinityTerm
          elementRelationship: atomic
    - name: requiredDuringSchedulingIgnoredDuringExecution
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PodAffinityTerm
          elementRelationship: atomic
- name: io.k8s.api.core.v1.PodCondition
  map:
    fields:
    - name: lastProbeTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.PodDNSConfig
  map:
    fields:
    - name: nameservers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: options
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PodDNSConfigOption
          elementRelationship: atomic
    - name: searches
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.PodDNSConfigOption
  map:
    fields:
    - name: name
      type:
        scalar: string
    - name: value
      type:
        scalar: string
- name: io.k8s.api.core.v1.PodIP
  map:
    fields:
    - name: ip
      type:
        scalar: string
- name: io.k8s.api.core.v1.PodOS
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.PodReadinessGate
  map:
    fields:
    - name: conditionType
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.PodSecurityContext
  map:
    fields:
    - name: fsGroup
      type:
        scalar: numeric
    - name: fsGroupChangePolicy
      type:
        scalar: string
    - name: runAsGroup
      type:
        scalar: numeric
    - name: runAsNonRoot
      type:
        scalar: boolean
    - name: runAsUser
      type:
        scalar: numeric
    - name: seLinuxOptions
      type:
        namedType: io.k8s.api.core.v1.SELinuxOptions
    - name: seccompProfile
      type:
        namedType: io.k8s.api.core.v1.SeccompProfile
    - name: supplementalGroups
      type:
        list:
          elementType:
            scalar: numeric
          elementRelationship: atomic
    - name: sysctls
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Sysctl
          elementRelationship: atomic
    - name: windowsOptions
      type:
        namedType: io.k8s.api.core.v1.WindowsSecurityContextOptions
- name: io.k8s.api.core.v1.PodSpec
  map:
    fields:
    - name: activeDeadlineSeconds
      type:
        scalar: numeric
    - name: affinity
      type:
        namedType: io.k8s.api.core.v1.Affinity
    - name: automountServiceAccountToken
      type:
        scalar: boolean
    - name: containers
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Container
          elementRelationship: associative
          keys:
          - name
    - name: dnsConfig
      type:
        namedType: io.k8s.api.core.v1.PodDNSConfig
    - name: dnsPolicy
      type:
        scalar: string
    - name: enableServiceLinks
      type:
        scalar: boolean
    - name: ephemeralContainers
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.EphemeralContainer
          elementRelationship: associative
          keys:
          - name
    - name: hostAliases
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.HostAlias
          elementRelationship: associative
          keys:
          - ip
    - name: hostIPC
      type:
        scalar: boolean
    - name: hostNetwork
      type:
        scalar: boolean
    - name: hostPID
      type:
        scalar: boolean
    - name: hostname
      type:
        scalar: string
    - name: imagePullSecrets
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.LocalObjectReference
          elementRelationship: associative
          keys:
          - name
    - name: initContainers
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Container
          elementRelationship: associative
          keys:
          - name
    - name: nodeName
      type:
        scalar: string
    - name: nodeSelector
      type:
        map:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: os
      type:
        namedType: io.k8s.api.core.v1.PodOS
    - name: overhead
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: preemptionPolicy
      type:
        scalar: string
    - name: priority
      type:
        scalar: numeric
    - name: priorityClassName
      type:
        scalar: string
    - name: readinessGates
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PodReadinessGate
          elementRelationship: atomic
    - name: restartPolicy
      type:
        scalar: string
    - name: runtimeClassName
      type:
        scalar: string
    - name: schedulerName
      type:
        scalar: string
    - name: securityContext
      type:
        namedType: io.k8s.api.core.v1.PodSecurityContext
    - name: serviceAccount
      type:
        scalar: string
    - name: serviceAccountName
      type:
        scalar: string
    - name: setHostnameAsFQDN
      type:
        scalar: boolean
    - name: shareProcessNamespace
      type:
        scalar: boolean
    - name: subdomain
      type:
        scalar: string
    - name: terminationGracePeriodSeconds
      type:
        scalar: numeric
    - name: tolerations
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Toleration
          elementRelationship: atomic
    - name: topologySpreadConstraints
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.TopologySpreadConstraint
          elementRelationship: associative
          keys:
          - topologyKey
          - whenUnsatisfiable
    - name: volumes
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Volume
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.api.core.v1.PodStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PodCondition
          elementRelationship: associative
          keys:
          - type
    - name: containerStatuses
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ContainerStatus
          elementRelationship: atomic
    - name: ephemeralContainerStatuses
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ContainerStatus
          elementRelationship: atomic
    - name: hostIP
      type:
        scalar: string
    - name: initContainerStatuses
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ContainerStatus
          elementRelationship: atomic
    - name: message
      type:
        scalar: string
    - name: nominatedNodeName
      type:
        scalar: string
    - name: phase
      type:
        scalar: string
    - name: podIP
      type:
        scalar: string
    - name: podIPs
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.PodIP
          elementRelationship: associative
          keys:
          - ip
    - name: qosClass
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: startTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.api.core.v1.PodTemplate
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
- name: io.k8s.api.core.v1.PodTemplateSpec
  map:
    fields:
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.PodSpec
      default: {}
- name: io.k8s.api.core.v1.PortStatus
  map:
    fields:
    - name: error
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
      default: 0
    - name: protocol
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.PortworxVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: volumeID
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.PreferredSchedulingTerm
  map:
    fields:
    - name: preference
      type:
        namedType: io.k8s.api.core.v1.NodeSelectorTerm
      default: {}
    - name: weight
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.core.v1.Probe
  map:
    fields:
    - name: exec
      type:
        namedType: io.k8s.api.core.v1.ExecAction
    - name: failureThreshold
      type:
        scalar: numeric
    - name: grpc
      type:
        namedType: io.k8s.api.core.v1.GRPCAction
    - name: httpGet
      type:
        namedType: io.k8s.api.core.v1.HTTPGetAction
    - name: initialDelaySeconds
      type:
        scalar: numeric
    - name: periodSeconds
      type:
        scalar: numeric
    - name: successThreshold
      type:
        scalar: numeric
    - name: tcpSocket
      type:
        namedType: io.k8s.api.core.v1.TCPSocketAction
    - name: terminationGracePeriodSeconds
      type:
        scalar: numeric
    - name: timeoutSeconds
      type:
        scalar: numeric
- name: io.k8s.api.core.v1.ProjectedVolumeSource
  map:
    fields:
    - name: defaultMode
      type:
        scalar: numeric
    - name: sources
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.VolumeProjection
          elementRelationship: atomic
- name: io.k8s.api.core.v1.QuobyteVolumeSource
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: registry
      type:
        scalar: string
      default: ""
    - name: tenant
      type:
        scalar: string
    - name: user
      type:
        scalar: string
    - name: volume
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.RBDPersistentVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: image
      type:
        scalar: string
      default: ""
    - name: keyring
      type:
        scalar: string
    - name: monitors
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: pool
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
    - name: user
      type:
        scalar: string
- name: io.k8s.api.core.v1.RBDVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: image
      type:
        scalar: string
      default: ""
    - name: keyring
      type:
        scalar: string
    - name: monitors
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: pool
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
    - name: user
      type:
        scalar: string
- name: io.k8s.api.core.v1.ReplicationController
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.ReplicationControllerSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.core.v1.ReplicationControllerStatus
      default: {}
- name: io.k8s.api.core.v1.ReplicationControllerCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.ReplicationControllerSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: selector
      type:
        map:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
- name: io.k8s.api.core.v1.ReplicationControllerStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ReplicationControllerCondition
          elementRelationship: associative
          keys:
          - type
    - name: fullyLabeledReplicas
      type:
        scalar: numeric
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.core.v1.ResourceFieldSelector
  map:
    fields:
    - name: containerName
      type:
        scalar: string
    - name: divisor
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
      default: {}
    - name: resource
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.api.core.v1.ResourceQuota
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.ResourceQuotaSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.core.v1.ResourceQuotaStatus
      default: {}
- name: io.k8s.api.core.v1.ResourceQuotaSpec
  map:
    fields:
    - name: hard
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: scopeSelector
      type:
        namedType: io.k8s.api.core.v1.ScopeSelector
    - name: scopes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.ResourceQuotaStatus
  map:
    fields:
    - name: hard
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: used
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.core.v1.ResourceRequirements
  map:
    fields:
    - name: limits
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: requests
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.core.v1.SELinuxOptions
  map:
    fields:
    - name: level
      type:
        scalar: string
    - name: role
      type:
        scalar: string
    - name: type
      type:
        scalar: string
    - name: user
      type:
        scalar: string
- name: io.k8s.api.core.v1.ScaleIOPersistentVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: gateway
      type:
        scalar: string
      default: ""
    - name: protectionDomain
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.SecretReference
    - name: sslEnabled
      type:
        scalar: boolean
    - name: storageMode
      type:
        scalar: string
    - name: storagePool
      type:
        scalar: string
    - name: system
      type:
        scalar: string
      default: ""
    - name: volumeName
      type:
        scalar: string
- name: io.k8s.api.core.v1.ScaleIOVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: gateway
      type:
        scalar: string
      default: ""
    - name: protectionDomain
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
    - name: sslEnabled
      type:
        scalar: boolean
    - name: storageMode
      type:
        scalar: string
    - name: storagePool
      type:
        scalar: string
    - name: system
      type:
        scalar: string
      default: ""
    - name: volumeName
      type:
        scalar: string
- name: io.k8s.api.core.v1.ScopeSelector
  map:
    fields:
    - name: matchExpressions
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ScopedResourceSelectorRequirement
          elementRelationship: atomic
    elementRelationship: atomic
- name: io.k8s.api.core.v1.ScopedResourceSelectorRequirement
  map:
    fields:
    - name: operator
      type:
        scalar: string
      default: ""
    - name: scopeName
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.SeccompProfile
  map:
    fields:
    - name: localhostProfile
      type:
        scalar: string
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: localhostProfile
        discriminatorValue: LocalhostProfile
- name: io.k8s.api.core.v1.Secret
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: data
      type:
        map:
          elementType:
            scalar: string
    - name: immutable
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: stringData
      type:
        map:
          elementType:
            scalar: string
    - name: type
      type:
        scalar: string
- name: io.k8s.api.core.v1.SecretEnvSource
  map:
    fields:
    - name: name
      type:
        scalar: string
    - name: optional
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.SecretKeySelector
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
    - name: optional
      type:
        scalar: boolean
    elementRelationship: atomic
- name: io.k8s.api.core.v1.SecretProjection
  map:
    fields:
    - name: items
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.KeyToPath
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: optional
      type:
        scalar: boolean
- name: io.k8s.api.core.v1.SecretReference
  map:
    fields:
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    elementRelationship: atomic
- name: io.k8s.api.core.v1.SecretVolumeSource
  map:
    fields:
    - name: defaultMode
      type:
        scalar: numeric
    - name: items
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.KeyToPath
          elementRelationship: atomic
    - name: optional
      type:
        scalar: boolean
    - name: secretName
      type:
        scalar: string
- name: io.k8s.api.core.v1.SecurityContext
  map:
    fields:
    - name: allowPrivilegeEscalation
      type:
        scalar: boolean
    - name: capabilities
      type:
        namedType: io.k8s.api.core.v1.Capabilities
    - name: privileged
      type:
        scalar: boolean
    - name: procMount
      type:
        scalar: string
    - name: readOnlyRootFilesystem
      type:
        scalar: boolean
    - name: runAsGroup
      type:
        scalar: numeric
    - name: runAsNonRoot
      type:
        scalar: boolean
    - name: runAsUser
      type:
        scalar: numeric
    - name: seLinuxOptions
      type:
        namedType: io.k8s.api.core.v1.SELinuxOptions
    - name: seccompProfile
      type:
        namedType: io.k8s.api.core.v1.SeccompProfile
    - name: windowsOptions
      type:
        namedType: io.k8s.api.core.v1.WindowsSecurityContextOptions
- name: io.k8s.api.core.v1.Service
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.core.v1.ServiceSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.core.v1.ServiceStatus
      default: {}
- name: io.k8s.api.core.v1.ServiceAccount
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: automountServiceAccountToken
      type:
        scalar: boolean
    - name: imagePullSecrets
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.LocalObjectReference
          elementRelationship: atomic
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: secrets
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ObjectReference
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.api.core.v1.ServiceAccountTokenProjection
  map:
    fields:
    - name: audience
      type:
        scalar: string
    - name: expirationSeconds
      type:
        scalar: numeric
    - name: path
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.ServicePort
  map:
    fields:
    - name: appProtocol
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: nodePort
      type:
        scalar: numeric
    - name: port
      type:
        scalar: numeric
      default: 0
    - name: protocol
      type:
        scalar: string
      default: TCP
    - name: targetPort
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
      default: {}
- name: io.k8s.api.core.v1.ServiceSpec
  map:
    fields:
    - name: allocateLoadBalancerNodePorts
      type:
        scalar: boolean
    - name: clusterIP
      type:
        scalar: string
    - name: clusterIPs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: externalIPs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: externalName
      type:
        scalar: string
    - name: externalTrafficPolicy
      type:
        scalar: string
    - name: healthCheckNodePort
      type:
        scalar: numeric
    - name: internalTrafficPolicy
      type:
        scalar: string
    - name: ipFamilies
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: ipFamilyPolicy
      type:
        scalar: string
    - name: loadBalancerClass
      type:
        scalar: string
    - name: loadBalancerIP
      type:
        scalar: string
    - name: loadBalancerSourceRanges
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ServicePort
          elementRelationship: associative
          keys:
          - port
          - protocol
    - name: publishNotReadyAddresses
      type:
        scalar: boolean
    - name: selector
      type:
        map:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: sessionAffinity
      type:
        scalar: string
    - name: sessionAffinityConfig
      type:
        namedType: io.k8s.api.core.v1.SessionAffinityConfig
    - name: type
      type:
        scalar: string
- name: io.k8s.api.core.v1.ServiceStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: loadBalancer
      type:
        namedType: io.k8s.api.core.v1.LoadBalancerStatus
      default: {}
- name: io.k8s.api.core.v1.SessionAffinityConfig
  map:
    fields:
    - name: clientIP
      type:
        namedType: io.k8s.api.core.v1.ClientIPConfig
- name: io.k8s.api.core.v1.StorageOSPersistentVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    - name: volumeName
      type:
        scalar: string
    - name: volumeNamespace
      type:
        scalar: string
- name: io.k8s.api.core.v1.StorageOSVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
    - name: secretRef
      type:
        namedType: io.k8s.api.core.v1.LocalObjectReference
    - name: volumeName
      type:
        scalar: string
    - name: volumeNamespace
      type:
        scalar: string
- name: io.k8s.api.core.v1.Sysctl
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: value
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.TCPSocketAction
  map:
    fields:
    - name: host
      type:
        scalar: string
    - name: port
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
      default: {}
- name: io.k8s.api.core.v1.Taint
  map:
    fields:
    - name: effect
      type:
        scalar: string
      default: ""
    - name: key
      type:
        scalar: string
      default: ""
    - name: timeAdded
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: value
      type:
        scalar: string
- name: io.k8s.api.core.v1.Toleration
  map:
    fields:
    - name: effect
      type:
        scalar: string
    - name: key
      type:
        scalar: string
    - name: operator
      type:
        scalar: string
    - name: tolerationSeconds
      type:
        scalar: numeric
    - name: value
      type:
        scalar: string
- name: io.k8s.api.core.v1.TopologySelectorLabelRequirement
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.TopologySelectorTerm
  map:
    fields:
    - name: matchLabelExpressions
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.TopologySelectorLabelRequirement
          elementRelationship: atomic
    elementRelationship: atomic
- name: io.k8s.api.core.v1.TopologySpreadConstraint
  map:
    fields:
    - name: labelSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: maxSkew
      type:
        scalar: numeric
      default: 0
    - name: minDomains
      type:
        scalar: numeric
    - name: topologyKey
      type:
        scalar: string
      default: ""
    - name: whenUnsatisfiable
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.TypedLocalObjectReference
  map:
    fields:
    - name: apiGroup
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.api.core.v1.Volume
  map:
    fields:
    - name: awsElasticBlockStore
      type:
        namedType: io.k8s.api.core.v1.AWSElasticBlockStoreVolumeSource
    - name: azureDisk
      type:
        namedType: io.k8s.api.core.v1.AzureDiskVolumeSource
    - name: azureFile
      type:
        namedType: io.k8s.api.core.v1.AzureFileVolumeSource
    - name: cephfs
      type:
        namedType: io.k8s.api.core.v1.CephFSVolumeSource
    - name: cinder
      type:
        namedType: io.k8s.api.core.v1.CinderVolumeSource
    - name: configMap
      type:
        namedType: io.k8s.api.core.v1.ConfigMapVolumeSource
    - name: csi
      type:
        namedType: io.k8s.api.core.v1.CSIVolumeSource
    - name: downwardAPI
      type:
        namedType: io.k8s.api.core.v1.DownwardAPIVolumeSource
    - name: emptyDir
      type:
        namedType: io.k8s.api.core.v1.EmptyDirVolumeSource
    - name: ephemeral
      type:
        namedType: io.k8s.api.core.v1.EphemeralVolumeSource
    - name: fc
      type:
        namedType: io.k8s.api.core.v1.FCVolumeSource
    - name: flexVolume
      type:
        namedType: io.k8s.api.core.v1.FlexVolumeSource
    - name: flocker
      type:
        namedType: io.k8s.api.core.v1.FlockerVolumeSource
    - name: gcePersistentDisk
      type:
        namedType: io.k8s.api.core.v1.GCEPersistentDiskVolumeSource
    - name: gitRepo
      type:
        namedType: io.k8s.api.core.v1.GitRepoVolumeSource
    - name: glusterfs
      type:
        namedType: io.k8s.api.core.v1.GlusterfsVolumeSource
    - name: hostPath
      type:
        namedType: io.k8s.api.core.v1.HostPathVolumeSource
    - name: iscsi
      type:
        namedType: io.k8s.api.core.v1.ISCSIVolumeSource
    - name: name
      type:
        scalar: string
      default: ""
    - name: nfs
      type:
        namedType: io.k8s.api.core.v1.NFSVolumeSource
    - name: persistentVolumeClaim
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeClaimVolumeSource
    - name: photonPersistentDisk
      type:
        namedType: io.k8s.api.core.v1.PhotonPersistentDiskVolumeSource
    - name: portworxVolume
      type:
        namedType: io.k8s.api.core.v1.PortworxVolumeSource
    - name: projected
      type:
        namedType: io.k8s.api.core.v1.ProjectedVolumeSource
    - name: quobyte
      type:
        namedType: io.k8s.api.core.v1.QuobyteVolumeSource
    - name: rbd
      type:
        namedType: io.k8s.api.core.v1.RBDVolumeSource
    - name: scaleIO
      type:
        namedType: io.k8s.api.core.v1.ScaleIOVolumeSource
    - name: secret
      type:
        namedType: io.k8s.api.core.v1.SecretVolumeSource
    - name: storageos
      type:
        namedType: io.k8s.api.core.v1.StorageOSVolumeSource
    - name: vsphereVolume
      type:
        namedType: io.k8s.api.core.v1.VsphereVirtualDiskVolumeSource
- name: io.k8s.api.core.v1.VolumeDevice
  map:
    fields:
    - name: devicePath
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.VolumeMount
  map:
    fields:
    - name: mountPath
      type:
        scalar: string
      default: ""
    - name: mountPropagation
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: readOnly
      type:
        scalar: boolean
    - name: subPath
      type:
        scalar: string
    - name: subPathExpr
      type:
        scalar: string
- name: io.k8s.api.core.v1.VolumeNodeAffinity
  map:
    fields:
    - name: required
      type:
        namedType: io.k8s.api.core.v1.NodeSelector
- name: io.k8s.api.core.v1.VolumeProjection
  map:
    fields:
    - name: configMap
      type:
        namedType: io.k8s.api.core.v1.ConfigMapProjection
    - name: downwardAPI
      type:
        namedType: io.k8s.api.core.v1.DownwardAPIProjection
    - name: secret
      type:
        namedType: io.k8s.api.core.v1.SecretProjection
    - name: serviceAccountToken
      type:
        namedType: io.k8s.api.core.v1.ServiceAccountTokenProjection
- name: io.k8s.api.core.v1.VsphereVirtualDiskVolumeSource
  map:
    fields:
    - name: fsType
      type:
        scalar: string
    - name: storagePolicyID
      type:
        scalar: string
    - name: storagePolicyName
      type:
        scalar: string
    - name: volumePath
      type:
        scalar: string
      default: ""
- name: io.k8s.api.core.v1.WeightedPodAffinityTerm
  map:
    fields:
    - name: podAffinityTerm
      type:
        namedType: io.k8s.api.core.v1.PodAffinityTerm
      default: {}
    - name: weight
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.core.v1.WindowsSecurityContextOptions
  map:
    fields:
    - name: gmsaCredentialSpec
      type:
        scalar: string
    - name: gmsaCredentialSpecName
      type:
        scalar: string
    - name: hostProcess
      type:
        scalar: boolean
    - name: runAsUserName
      type:
        scalar: string
- name: io.k8s.api.discovery.v1.Endpoint
  map:
    fields:
    - name: addresses
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: conditions
      type:
        namedType: io.k8s.api.discovery.v1.EndpointConditions
      default: {}
    - name: deprecatedTopology
      type:
        map:
          elementType:
            scalar: string
    - name: hints
      type:
        namedType: io.k8s.api.discovery.v1.EndpointHints
    - name: hostname
      type:
        scalar: string
    - name: nodeName
      type:
        scalar: string
    - name: targetRef
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    - name: zone
      type:
        scalar: string
- name: io.k8s.api.discovery.v1.EndpointConditions
  map:
    fields:
    - name: ready
      type:
        scalar: boolean
    - name: serving
      type:
        scalar: boolean
    - name: terminating
      type:
        scalar: boolean
- name: io.k8s.api.discovery.v1.EndpointHints
  map:
    fields:
    - name: forZones
      type:
        list:
          elementType:
            namedType: io.k8s.api.discovery.v1.ForZone
          elementRelationship: atomic
- name: io.k8s.api.discovery.v1.EndpointPort
  map:
    fields:
    - name: appProtocol
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
    - name: protocol
      type:
        scalar: string
    elementRelationship: atomic
- name: io.k8s.api.discovery.v1.EndpointSlice
  map:
    fields:
    - name: addressType
      type:
        scalar: string
      default: ""
    - name: apiVersion
      type:
        scalar: string
    - name: endpoints
      type:
        list:
          elementType:
            namedType: io.k8s.api.discovery.v1.Endpoint
          elementRelationship: atomic
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.discovery.v1.EndpointPort
          elementRelationship: atomic
- name: io.k8s.api.discovery.v1.ForZone
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.discovery.v1beta1.Endpoint
  map:
    fields:
    - name: addresses
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: conditions
      type:
        namedType: io.k8s.api.discovery.v1beta1.EndpointConditions
      default: {}
    - name: hints
      type:
        namedType: io.k8s.api.discovery.v1beta1.EndpointHints
    - name: hostname
      type:
        scalar: string
    - name: nodeName
      type:
        scalar: string
    - name: targetRef
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    - name: topology
      type:
        map:
          elementType:
            scalar: string
- name: io.k8s.api.discovery.v1beta1.EndpointConditions
  map:
    fields:
    - name: ready
      type:
        scalar: boolean
    - name: serving
      type:
        scalar: boolean
    - name: terminating
      type:
        scalar: boolean
- name: io.k8s.api.discovery.v1beta1.EndpointHints
  map:
    fields:
    - name: forZones
      type:
        list:
          elementType:
            namedType: io.k8s.api.discovery.v1beta1.ForZone
          elementRelationship: atomic
- name: io.k8s.api.discovery.v1beta1.EndpointPort
  map:
    fields:
    - name: appProtocol
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
    - name: protocol
      type:
        scalar: string
- name: io.k8s.api.discovery.v1beta1.EndpointSlice
  map:
    fields:
    - name: addressType
      type:
        scalar: string
      default: ""
    - name: apiVersion
      type:
        scalar: string
    - name: endpoints
      type:
        list:
          elementType:
            namedType: io.k8s.api.discovery.v1beta1.Endpoint
          elementRelationship: atomic
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.discovery.v1beta1.EndpointPort
          elementRelationship: atomic
- name: io.k8s.api.discovery.v1beta1.ForZone
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.events.v1.Event
  map:
    fields:
    - name: action
      type:
        scalar: string
    - name: apiVersion
      type:
        scalar: string
    - name: deprecatedCount
      type:
        scalar: numeric
    - name: deprecatedFirstTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: deprecatedLastTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: deprecatedSource
      type:
        namedType: io.k8s.api.core.v1.EventSource
      default: {}
    - name: eventTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
      default: {}
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: note
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: regarding
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
      default: {}
    - name: related
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    - name: reportingController
      type:
        scalar: string
    - name: reportingInstance
      type:
        scalar: string
    - name: series
      type:
        namedType: io.k8s.api.events.v1.EventSeries
    - name: type
      type:
        scalar: string
- name: io.k8s.api.events.v1.EventSeries
  map:
    fields:
    - name: count
      type:
        scalar: numeric
      default: 0
    - name: lastObservedTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
      default: {}
- name: io.k8s.api.events.v1beta1.Event
  map:
    fields:
    - name: action
      type:
        scalar: string
    - name: apiVersion
      type:
        scalar: string
    - name: deprecatedCount
      type:
        scalar: numeric
    - name: deprecatedFirstTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: deprecatedLastTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: deprecatedSource
      type:
        namedType: io.k8s.api.core.v1.EventSource
      default: {}
    - name: eventTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
      default: {}
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: note
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: regarding
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
      default: {}
    - name: related
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
    - name: reportingController
      type:
        scalar: string
    - name: reportingInstance
      type:
        scalar: string
    - name: series
      type:
        namedType: io.k8s.api.events.v1beta1.EventSeries
    - name: type
      type:
        scalar: string
- name: io.k8s.api.events.v1beta1.EventSeries
  map:
    fields:
    - name: count
      type:
        scalar: numeric
      default: 0
    - name: lastObservedTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
      default: {}
- name: io.k8s.api.extensions.v1beta1.AllowedCSIDriver
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.extensions.v1beta1.AllowedFlexVolume
  map:
    fields:
    - name: driver
      type:
        scalar: string
      default: ""
- name: io.k8s.api.extensions.v1beta1.AllowedHostPath
  map:
    fields:
    - name: pathPrefix
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
- name: io.k8s.api.extensions.v1beta1.DaemonSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.extensions.v1beta1.DaemonSetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.extensions.v1beta1.DaemonSetStatus
      default: {}
- name: io.k8s.api.extensions.v1beta1.DaemonSetCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.extensions.v1beta1.DaemonSetSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
    - name: templateGeneration
      type:
        scalar: numeric
    - name: updateStrategy
      type:
        namedType: io.k8s.api.extensions.v1beta1.DaemonSetUpdateStrategy
      default: {}
- name: io.k8s.api.extensions.v1beta1.DaemonSetStatus
  map:
    fields:
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.DaemonSetCondition
          elementRelationship: associative
          keys:
          - type
    - name: currentNumberScheduled
      type:
        scalar: numeric
      default: 0
    - name: desiredNumberScheduled
      type:
        scalar: numeric
      default: 0
    - name: numberAvailable
      type:
        scalar: numeric
    - name: numberMisscheduled
      type:
        scalar: numeric
      default: 0
    - name: numberReady
      type:
        scalar: numeric
      default: 0
    - name: numberUnavailable
      type:
        scalar: numeric
    - name: observedGeneration
      type:
        scalar: numeric
    - name: updatedNumberScheduled
      type:
        scalar: numeric
- name: io.k8s.api.extensions.v1beta1.DaemonSetUpdateStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.extensions.v1beta1.RollingUpdateDaemonSet
    - name: type
      type:
        scalar: string
- name: io.k8s.api.extensions.v1beta1.Deployment
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.extensions.v1beta1.DeploymentSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.extensions.v1beta1.DeploymentStatus
      default: {}
- name: io.k8s.api.extensions.v1beta1.DeploymentCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: lastUpdateTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.extensions.v1beta1.DeploymentSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: paused
      type:
        scalar: boolean
    - name: progressDeadlineSeconds
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: revisionHistoryLimit
      type:
        scalar: numeric
    - name: rollbackTo
      type:
        namedType: io.k8s.api.extensions.v1beta1.RollbackConfig
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: strategy
      type:
        namedType: io.k8s.api.extensions.v1beta1.DeploymentStrategy
      default: {}
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
- name: io.k8s.api.extensions.v1beta1.DeploymentStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
    - name: collisionCount
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.DeploymentCondition
          elementRelationship: associative
          keys:
          - type
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: unavailableReplicas
      type:
        scalar: numeric
    - name: updatedReplicas
      type:
        scalar: numeric
- name: io.k8s.api.extensions.v1beta1.DeploymentStrategy
  map:
    fields:
    - name: rollingUpdate
      type:
        namedType: io.k8s.api.extensions.v1beta1.RollingUpdateDeployment
    - name: type
      type:
        scalar: string
- name: io.k8s.api.extensions.v1beta1.FSGroupStrategyOptions
  map:
    fields:
    - name: ranges
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.IDRange
          elementRelationship: atomic
    - name: rule
      type:
        scalar: string
- name: io.k8s.api.extensions.v1beta1.HTTPIngressPath
  map:
    fields:
    - name: backend
      type:
        namedType: io.k8s.api.extensions.v1beta1.IngressBackend
      default: {}
    - name: path
      type:
        scalar: string
    - name: pathType
      type:
        scalar: string
- name: io.k8s.api.extensions.v1beta1.HTTPIngressRuleValue
  map:
    fields:
    - name: paths
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.HTTPIngressPath
          elementRelationship: atomic
- name: io.k8s.api.extensions.v1beta1.HostPortRange
  map:
    fields:
    - name: max
      type:
        scalar: numeric
      default: 0
    - name: min
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.extensions.v1beta1.IDRange
  map:
    fields:
    - name: max
      type:
        scalar: numeric
      default: 0
    - name: min
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.extensions.v1beta1.IPBlock
  map:
    fields:
    - name: cidr
      type:
        scalar: string
      default: ""
    - name: except
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.extensions.v1beta1.Ingress
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.extensions.v1beta1.IngressSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.extensions.v1beta1.IngressStatus
      default: {}
- name: io.k8s.api.extensions.v1beta1.IngressBackend
  map:
    fields:
    - name: resource
      type:
        namedType: io.k8s.api.core.v1.TypedLocalObjectReference
    - name: serviceName
      type:
        scalar: string
    - name: servicePort
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
      default: {}
- name: io.k8s.api.extensions.v1beta1.IngressRule
  map:
    fields:
    - name: host
      type:
        scalar: string
    - name: http
      type:
        namedType: io.k8s.api.extensions.v1beta1.HTTPIngressRuleValue
- name: io.k8s.api.extensions.v1beta1.IngressSpec
  map:
    fields:
    - name: backend
      type:
        namedType: io.k8s.api.extensions.v1beta1.IngressBackend
    - name: ingressClassName
      type:
        scalar: string
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.IngressRule
          elementRelationship: atomic
    - name: tls
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.IngressTLS
          elementRelationship: atomic
- name: io.k8s.api.extensions.v1beta1.IngressStatus
  map:
    fields:
    - name: loadBalancer
      type:
        namedType: io.k8s.api.core.v1.LoadBalancerStatus
      default: {}
- name: io.k8s.api.extensions.v1beta1.IngressTLS
  map:
    fields:
    - name: hosts
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: secretName
      type:
        scalar: string
- name: io.k8s.api.extensions.v1beta1.NetworkPolicy
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.extensions.v1beta1.NetworkPolicySpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.extensions.v1beta1.NetworkPolicyStatus
      default: {}
- name: io.k8s.api.extensions.v1beta1.NetworkPolicyEgressRule
  map:
    fields:
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.NetworkPolicyPort
          elementRelationship: atomic
    - name: to
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.NetworkPolicyPeer
          elementRelationship: atomic
- name: io.k8s.api.extensions.v1beta1.NetworkPolicyIngressRule
  map:
    fields:
    - name: from
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.NetworkPolicyPeer
          elementRelationship: atomic
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.NetworkPolicyPort
          elementRelationship: atomic
- name: io.k8s.api.extensions.v1beta1.NetworkPolicyPeer
  map:
    fields:
    - name: ipBlock
      type:
        namedType: io.k8s.api.extensions.v1beta1.IPBlock
    - name: namespaceSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: podSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
- name: io.k8s.api.extensions.v1beta1.NetworkPolicyPort
  map:
    fields:
    - name: endPort
      type:
        scalar: numeric
    - name: port
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: protocol
      type:
        scalar: string
- name: io.k8s.api.extensions.v1beta1.NetworkPolicySpec
  map:
    fields:
    - name: egress
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.NetworkPolicyEgressRule
          elementRelationship: atomic
    - name: ingress
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.NetworkPolicyIngressRule
          elementRelationship: atomic
    - name: podSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
      default: {}
    - name: policyTypes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.extensions.v1beta1.NetworkPolicyStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.extensions.v1beta1.PodSecurityPolicy
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.extensions.v1beta1.PodSecurityPolicySpec
      default: {}
- name: io.k8s.api.extensions.v1beta1.PodSecurityPolicySpec
  map:
    fields:
    - name: allowPrivilegeEscalation
      type:
        scalar: boolean
    - name: allowedCSIDrivers
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.AllowedCSIDriver
          elementRelationship: atomic
    - name: allowedCapabilities
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: allowedFlexVolumes
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.AllowedFlexVolume
          elementRelationship: atomic
    - name: allowedHostPaths
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.AllowedHostPath
          elementRelationship: atomic
    - name: allowedProcMountTypes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: allowedUnsafeSysctls
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: defaultAddCapabilities
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: defaultAllowPrivilegeEscalation
      type:
        scalar: boolean
    - name: forbiddenSysctls
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: fsGroup
      type:
        namedType: io.k8s.api.extensions.v1beta1.FSGroupStrategyOptions
      default: {}
    - name: hostIPC
      type:
        scalar: boolean
    - name: hostNetwork
      type:
        scalar: boolean
    - name: hostPID
      type:
        scalar: boolean
    - name: hostPorts
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.HostPortRange
          elementRelationship: atomic
    - name: privileged
      type:
        scalar: boolean
    - name: readOnlyRootFilesystem
      type:
        scalar: boolean
    - name: requiredDropCapabilities
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: runAsGroup
      type:
        namedType: io.k8s.api.extensions.v1beta1.RunAsGroupStrategyOptions
    - name: runAsUser
      type:
        namedType: io.k8s.api.extensions.v1beta1.RunAsUserStrategyOptions
      default: {}
    - name: runtimeClass
      type:
        namedType: io.k8s.api.extensions.v1beta1.RuntimeClassStrategyOptions
    - name: seLinux
      type:
        namedType: io.k8s.api.extensions.v1beta1.SELinuxStrategyOptions
      default: {}
    - name: supplementalGroups
      type:
        namedType: io.k8s.api.extensions.v1beta1.SupplementalGroupsStrategyOptions
      default: {}
    - name: volumes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.extensions.v1beta1.ReplicaSet
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.extensions.v1beta1.ReplicaSetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.extensions.v1beta1.ReplicaSetStatus
      default: {}
- name: io.k8s.api.extensions.v1beta1.ReplicaSetCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.extensions.v1beta1.ReplicaSetSpec
  map:
    fields:
    - name: minReadySeconds
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: template
      type:
        namedType: io.k8s.api.core.v1.PodTemplateSpec
      default: {}
- name: io.k8s.api.extensions.v1beta1.ReplicaSetStatus
  map:
    fields:
    - name: availableReplicas
      type:
        scalar: numeric
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.ReplicaSetCondition
          elementRelationship: associative
          keys:
          - type
    - name: fullyLabeledReplicas
      type:
        scalar: numeric
    - name: observedGeneration
      type:
        scalar: numeric
    - name: readyReplicas
      type:
        scalar: numeric
    - name: replicas
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.extensions.v1beta1.RollbackConfig
  map:
    fields:
    - name: revision
      type:
        scalar: numeric
- name: io.k8s.api.extensions.v1beta1.RollingUpdateDaemonSet
  map:
    fields:
    - name: maxSurge
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
- name: io.k8s.api.extensions.v1beta1.RollingUpdateDeployment
  map:
    fields:
    - name: maxSurge
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
- name: io.k8s.api.extensions.v1beta1.RunAsGroupStrategyOptions
  map:
    fields:
    - name: ranges
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.IDRange
          elementRelationship: atomic
    - name: rule
      type:
        scalar: string
      default: ""
- name: io.k8s.api.extensions.v1beta1.RunAsUserStrategyOptions
  map:
    fields:
    - name: ranges
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.IDRange
          elementRelationship: atomic
    - name: rule
      type:
        scalar: string
      default: ""
- name: io.k8s.api.extensions.v1beta1.RuntimeClassStrategyOptions
  map:
    fields:
    - name: allowedRuntimeClassNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: defaultRuntimeClassName
      type:
        scalar: string
- name: io.k8s.api.extensions.v1beta1.SELinuxStrategyOptions
  map:
    fields:
    - name: rule
      type:
        scalar: string
      default: ""
    - name: seLinuxOptions
      type:
        namedType: io.k8s.api.core.v1.SELinuxOptions
- name: io.k8s.api.extensions.v1beta1.SupplementalGroupsStrategyOptions
  map:
    fields:
    - name: ranges
      type:
        list:
          elementType:
            namedType: io.k8s.api.extensions.v1beta1.IDRange
          elementRelationship: atomic
    - name: rule
      type:
        scalar: string
- name: io.k8s.api.flowcontrol.v1alpha1.FlowDistinguisherMethod
  map:
    fields:
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1alpha1.FlowSchema
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.FlowSchemaSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.FlowSchemaStatus
      default: {}
- name: io.k8s.api.flowcontrol.v1alpha1.FlowSchemaCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
    - name: type
      type:
        scalar: string
- name: io.k8s.api.flowcontrol.v1alpha1.FlowSchemaSpec
  map:
    fields:
    - name: distinguisherMethod
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.FlowDistinguisherMethod
    - name: matchingPrecedence
      type:
        scalar: numeric
      default: 0
    - name: priorityLevelConfiguration
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.PriorityLevelConfigurationReference
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1alpha1.PolicyRulesWithSubjects
          elementRelationship: atomic
- name: io.k8s.api.flowcontrol.v1alpha1.FlowSchemaStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1alpha1.FlowSchemaCondition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.flowcontrol.v1alpha1.GroupSubject
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1alpha1.LimitResponse
  map:
    fields:
    - name: queuing
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.QueuingConfiguration
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: queuing
        discriminatorValue: Queuing
- name: io.k8s.api.flowcontrol.v1alpha1.LimitedPriorityLevelConfiguration
  map:
    fields:
    - name: assuredConcurrencyShares
      type:
        scalar: numeric
      default: 0
    - name: limitResponse
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.LimitResponse
      default: {}
- name: io.k8s.api.flowcontrol.v1alpha1.NonResourcePolicyRule
  map:
    fields:
    - name: nonResourceURLs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: io.k8s.api.flowcontrol.v1alpha1.PolicyRulesWithSubjects
  map:
    fields:
    - name: nonResourceRules
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1alpha1.NonResourcePolicyRule
          elementRelationship: atomic
    - name: resourceRules
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1alpha1.ResourcePolicyRule
          elementRelationship: atomic
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1alpha1.Subject
          elementRelationship: atomic
- name: io.k8s.api.flowcontrol.v1alpha1.PriorityLevelConfiguration
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.PriorityLevelConfigurationSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.PriorityLevelConfigurationStatus
      default: {}
- name: io.k8s.api.flowcontrol.v1alpha1.PriorityLevelConfigurationCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
    - name: type
      type:
        scalar: string
- name: io.k8s.api.flowcontrol.v1alpha1.PriorityLevelConfigurationReference
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1alpha1.PriorityLevelConfigurationSpec
  map:
    fields:
    - name: limited
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.LimitedPriorityLevelConfiguration
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: limited
        discriminatorValue: Limited
- name: io.k8s.api.flowcontrol.v1alpha1.PriorityLevelConfigurationStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1alpha1.PriorityLevelConfigurationCondition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.flowcontrol.v1alpha1.QueuingConfiguration
  map:
    fields:
    - name: handSize
      type:
        scalar: numeric
      default: 0
    - name: queueLengthLimit
      type:
        scalar: numeric
      default: 0
    - name: queues
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.flowcontrol.v1alpha1.ResourcePolicyRule
  map:
    fields:
    - name: apiGroups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: clusterScope
      type:
        scalar: boolean
    - name: namespaces
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: resources
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: io.k8s.api.flowcontrol.v1alpha1.ServiceAccountSubject
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1alpha1.Subject
  map:
    fields:
    - name: group
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.GroupSubject
    - name: kind
      type:
        scalar: string
      default: ""
    - name: serviceAccount
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.ServiceAccountSubject
    - name: user
      type:
        namedType: io.k8s.api.flowcontrol.v1alpha1.UserSubject
    unions:
    - discriminator: kind
      fields:
      - fieldName: group
        discriminatorValue: Group
      - fieldName: serviceAccount
        discriminatorValue: ServiceAccount
      - fieldName: user
        discriminatorValue: User
- name: io.k8s.api.flowcontrol.v1alpha1.UserSubject
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta1.FlowDistinguisherMethod
  map:
    fields:
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta1.FlowSchema
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.FlowSchemaSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.FlowSchemaStatus
      default: {}
- name: io.k8s.api.flowcontrol.v1beta1.FlowSchemaCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
    - name: type
      type:
        scalar: string
- name: io.k8s.api.flowcontrol.v1beta1.FlowSchemaSpec
  map:
    fields:
    - name: distinguisherMethod
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.FlowDistinguisherMethod
    - name: matchingPrecedence
      type:
        scalar: numeric
      default: 0
    - name: priorityLevelConfiguration
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.PriorityLevelConfigurationReference
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta1.PolicyRulesWithSubjects
          elementRelationship: atomic
- name: io.k8s.api.flowcontrol.v1beta1.FlowSchemaStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta1.FlowSchemaCondition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.flowcontrol.v1beta1.GroupSubject
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta1.LimitResponse
  map:
    fields:
    - name: queuing
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.QueuingConfiguration
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: queuing
        discriminatorValue: Queuing
- name: io.k8s.api.flowcontrol.v1beta1.LimitedPriorityLevelConfiguration
  map:
    fields:
    - name: assuredConcurrencyShares
      type:
        scalar: numeric
      default: 0
    - name: limitResponse
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.LimitResponse
      default: {}
- name: io.k8s.api.flowcontrol.v1beta1.NonResourcePolicyRule
  map:
    fields:
    - name: nonResourceURLs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: io.k8s.api.flowcontrol.v1beta1.PolicyRulesWithSubjects
  map:
    fields:
    - name: nonResourceRules
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta1.NonResourcePolicyRule
          elementRelationship: atomic
    - name: resourceRules
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta1.ResourcePolicyRule
          elementRelationship: atomic
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta1.Subject
          elementRelationship: atomic
- name: io.k8s.api.flowcontrol.v1beta1.PriorityLevelConfiguration
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.PriorityLevelConfigurationSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.PriorityLevelConfigurationStatus
      default: {}
- name: io.k8s.api.flowcontrol.v1beta1.PriorityLevelConfigurationCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
    - name: type
      type:
        scalar: string
- name: io.k8s.api.flowcontrol.v1beta1.PriorityLevelConfigurationReference
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta1.PriorityLevelConfigurationSpec
  map:
    fields:
    - name: limited
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.LimitedPriorityLevelConfiguration
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: limited
        discriminatorValue: Limited
- name: io.k8s.api.flowcontrol.v1beta1.PriorityLevelConfigurationStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta1.PriorityLevelConfigurationCondition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.flowcontrol.v1beta1.QueuingConfiguration
  map:
    fields:
    - name: handSize
      type:
        scalar: numeric
      default: 0
    - name: queueLengthLimit
      type:
        scalar: numeric
      default: 0
    - name: queues
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.flowcontrol.v1beta1.ResourcePolicyRule
  map:
    fields:
    - name: apiGroups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: clusterScope
      type:
        scalar: boolean
    - name: namespaces
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: resources
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: io.k8s.api.flowcontrol.v1beta1.ServiceAccountSubject
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta1.Subject
  map:
    fields:
    - name: group
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.GroupSubject
    - name: kind
      type:
        scalar: string
      default: ""
    - name: serviceAccount
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.ServiceAccountSubject
    - name: user
      type:
        namedType: io.k8s.api.flowcontrol.v1beta1.UserSubject
    unions:
    - discriminator: kind
      fields:
      - fieldName: group
        discriminatorValue: Group
      - fieldName: serviceAccount
        discriminatorValue: ServiceAccount
      - fieldName: user
        discriminatorValue: User
- name: io.k8s.api.flowcontrol.v1beta1.UserSubject
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta2.FlowDistinguisherMethod
  map:
    fields:
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta2.FlowSchema
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.FlowSchemaSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.FlowSchemaStatus
      default: {}
- name: io.k8s.api.flowcontrol.v1beta2.FlowSchemaCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
    - name: type
      type:
        scalar: string
- name: io.k8s.api.flowcontrol.v1beta2.FlowSchemaSpec
  map:
    fields:
    - name: distinguisherMethod
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.FlowDistinguisherMethod
    - name: matchingPrecedence
      type:
        scalar: numeric
      default: 0
    - name: priorityLevelConfiguration
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationReference
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta2.PolicyRulesWithSubjects
          elementRelationship: atomic
- name: io.k8s.api.flowcontrol.v1beta2.FlowSchemaStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta2.FlowSchemaCondition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.flowcontrol.v1beta2.GroupSubject
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta2.LimitResponse
  map:
    fields:
    - name: queuing
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.QueuingConfiguration
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: queuing
        discriminatorValue: Queuing
- name: io.k8s.api.flowcontrol.v1beta2.LimitedPriorityLevelConfiguration
  map:
    fields:
    - name: assuredConcurrencyShares
      type:
        scalar: numeric
      default: 0
    - name: limitResponse
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.LimitResponse
      default: {}
- name: io.k8s.api.flowcontrol.v1beta2.NonResourcePolicyRule
  map:
    fields:
    - name: nonResourceURLs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: io.k8s.api.flowcontrol.v1beta2.PolicyRulesWithSubjects
  map:
    fields:
    - name: nonResourceRules
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta2.NonResourcePolicyRule
          elementRelationship: atomic
    - name: resourceRules
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta2.ResourcePolicyRule
          elementRelationship: atomic
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta2.Subject
          elementRelationship: atomic
- name: io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfiguration
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationStatus
      default: {}
- name: io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationCondition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
    - name: reason
      type:
        scalar: string
    - name: status
      type:
        scalar: string
    - name: type
      type:
        scalar: string
- name: io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationReference
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationSpec
  map:
    fields:
    - name: limited
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.LimitedPriorityLevelConfiguration
    - name: type
      type:
        scalar: string
      default: ""
    unions:
    - discriminator: type
      fields:
      - fieldName: limited
        discriminatorValue: Limited
- name: io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationCondition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.flowcontrol.v1beta2.QueuingConfiguration
  map:
    fields:
    - name: handSize
      type:
        scalar: numeric
      default: 0
    - name: queueLengthLimit
      type:
        scalar: numeric
      default: 0
    - name: queues
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.flowcontrol.v1beta2.ResourcePolicyRule
  map:
    fields:
    - name: apiGroups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: clusterScope
      type:
        scalar: boolean
    - name: namespaces
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: resources
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: io.k8s.api.flowcontrol.v1beta2.ServiceAccountSubject
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
- name: io.k8s.api.flowcontrol.v1beta2.Subject
  map:
    fields:
    - name: group
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.GroupSubject
    - name: kind
      type:
        scalar: string
      default: ""
    - name: serviceAccount
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.ServiceAccountSubject
    - name: user
      type:
        namedType: io.k8s.api.flowcontrol.v1beta2.UserSubject
    unions:
    - discriminator: kind
      fields:
      - fieldName: group
        discriminatorValue: Group
      - fieldName: serviceAccount
        discriminatorValue: ServiceAccount
      - fieldName: user
        discriminatorValue: User
- name: io.k8s.api.flowcontrol.v1beta2.UserSubject
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.imagepolicy.v1alpha1.ImageReview
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.imagepolicy.v1alpha1.ImageReviewSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.imagepolicy.v1alpha1.ImageReviewStatus
      default: {}
- name: io.k8s.api.imagepolicy.v1alpha1.ImageReviewContainerSpec
  map:
    fields:
    - name: image
      type:
        scalar: string
- name: io.k8s.api.imagepolicy.v1alpha1.ImageReviewSpec
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: containers
      type:
        list:
          elementType:
            namedType: io.k8s.api.imagepolicy.v1alpha1.ImageReviewContainerSpec
          elementRelationship: atomic
    - name: namespace
      type:
        scalar: string
- name: io.k8s.api.imagepolicy.v1alpha1.ImageReviewStatus
  map:
    fields:
    - name: allowed
      type:
        scalar: boolean
      default: false
    - name: auditAnnotations
      type:
        map:
          elementType:
            scalar: string
    - name: reason
      type:
        scalar: string
- name: io.k8s.api.networking.v1.HTTPIngressPath
  map:
    fields:
    - name: backend
      type:
        namedType: io.k8s.api.networking.v1.IngressBackend
      default: {}
    - name: path
      type:
        scalar: string
    - name: pathType
      type:
        scalar: string
- name: io.k8s.api.networking.v1.HTTPIngressRuleValue
  map:
    fields:
    - name: paths
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1.HTTPIngressPath
          elementRelationship: atomic
- name: io.k8s.api.networking.v1.IPBlock
  map:
    fields:
    - name: cidr
      type:
        scalar: string
      default: ""
    - name: except
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.networking.v1.Ingress
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.networking.v1.IngressSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.networking.v1.IngressStatus
      default: {}
- name: io.k8s.api.networking.v1.IngressBackend
  map:
    fields:
    - name: resource
      type:
        namedType: io.k8s.api.core.v1.TypedLocalObjectReference
    - name: service
      type:
        namedType: io.k8s.api.networking.v1.IngressServiceBackend
- name: io.k8s.api.networking.v1.IngressClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.networking.v1.IngressClassSpec
      default: {}
- name: io.k8s.api.networking.v1.IngressClassParametersReference
  map:
    fields:
    - name: apiGroup
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    - name: scope
      type:
        scalar: string
- name: io.k8s.api.networking.v1.IngressClassSpec
  map:
    fields:
    - name: controller
      type:
        scalar: string
    - name: parameters
      type:
        namedType: io.k8s.api.networking.v1.IngressClassParametersReference
- name: io.k8s.api.networking.v1.IngressRule
  map:
    fields:
    - name: host
      type:
        scalar: string
    - name: http
      type:
        namedType: io.k8s.api.networking.v1.HTTPIngressRuleValue
- name: io.k8s.api.networking.v1.IngressServiceBackend
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: port
      type:
        namedType: io.k8s.api.networking.v1.ServiceBackendPort
      default: {}
- name: io.k8s.api.networking.v1.IngressSpec
  map:
    fields:
    - name: defaultBackend
      type:
        namedType: io.k8s.api.networking.v1.IngressBackend
    - name: ingressClassName
      type:
        scalar: string
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1.IngressRule
          elementRelationship: atomic
    - name: tls
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1.IngressTLS
          elementRelationship: atomic
- name: io.k8s.api.networking.v1.IngressStatus
  map:
    fields:
    - name: loadBalancer
      type:
        namedType: io.k8s.api.core.v1.LoadBalancerStatus
      default: {}
- name: io.k8s.api.networking.v1.IngressTLS
  map:
    fields:
    - name: hosts
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: secretName
      type:
        scalar: string
- name: io.k8s.api.networking.v1.NetworkPolicy
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.networking.v1.NetworkPolicySpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.networking.v1.NetworkPolicyStatus
      default: {}
- name: io.k8s.api.networking.v1.NetworkPolicyEgressRule
  map:
    fields:
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1.NetworkPolicyPort
          elementRelationship: atomic
    - name: to
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1.NetworkPolicyPeer
          elementRelationship: atomic
- name: io.k8s.api.networking.v1.NetworkPolicyIngressRule
  map:
    fields:
    - name: from
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1.NetworkPolicyPeer
          elementRelationship: atomic
    - name: ports
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1.NetworkPolicyPort
          elementRelationship: atomic
- name: io.k8s.api.networking.v1.NetworkPolicyPeer
  map:
    fields:
    - name: ipBlock
      type:
        namedType: io.k8s.api.networking.v1.IPBlock
    - name: namespaceSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: podSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
- name: io.k8s.api.networking.v1.NetworkPolicyPort
  map:
    fields:
    - name: endPort
      type:
        scalar: numeric
    - name: port
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: protocol
      type:
        scalar: string
- name: io.k8s.api.networking.v1.NetworkPolicySpec
  map:
    fields:
    - name: egress
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1.NetworkPolicyEgressRule
          elementRelationship: atomic
    - name: ingress
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1.NetworkPolicyIngressRule
          elementRelationship: atomic
    - name: podSelector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
      default: {}
    - name: policyTypes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.networking.v1.NetworkPolicyStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
- name: io.k8s.api.networking.v1.ServiceBackendPort
  map:
    fields:
    - name: name
      type:
        scalar: string
    - name: number
      type:
        scalar: numeric
- name: io.k8s.api.networking.v1alpha1.ClusterCIDRConfig
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.networking.v1alpha1.ClusterCIDRConfigSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.networking.v1alpha1.ClusterCIDRConfigStatus
      default: {}
- name: io.k8s.api.networking.v1alpha1.ClusterCIDRConfigSpec
  map:
    fields:
    - name: ipv4CIDR
      type:
        scalar: string
      default: ""
    - name: ipv6CIDR
      type:
        scalar: string
      default: ""
    - name: nodeSelector
      type:
        namedType: io.k8s.api.core.v1.NodeSelector
    - name: perNodeHostBits
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.networking.v1alpha1.ClusterCIDRConfigStatus
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.api.networking.v1beta1.HTTPIngressPath
  map:
    fields:
    - name: backend
      type:
        namedType: io.k8s.api.networking.v1beta1.IngressBackend
      default: {}
    - name: path
      type:
        scalar: string
    - name: pathType
      type:
        scalar: string
- name: io.k8s.api.networking.v1beta1.HTTPIngressRuleValue
  map:
    fields:
    - name: paths
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1beta1.HTTPIngressPath
          elementRelationship: atomic
- name: io.k8s.api.networking.v1beta1.Ingress
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.networking.v1beta1.IngressSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.networking.v1beta1.IngressStatus
      default: {}
- name: io.k8s.api.networking.v1beta1.IngressBackend
  map:
    fields:
    - name: resource
      type:
        namedType: io.k8s.api.core.v1.TypedLocalObjectReference
    - name: serviceName
      type:
        scalar: string
    - name: servicePort
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
      default: {}
- name: io.k8s.api.networking.v1beta1.IngressClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.networking.v1beta1.IngressClassSpec
      default: {}
- name: io.k8s.api.networking.v1beta1.IngressClassParametersReference
  map:
    fields:
    - name: apiGroup
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    - name: scope
      type:
        scalar: string
- name: io.k8s.api.networking.v1beta1.IngressClassSpec
  map:
    fields:
    - name: controller
      type:
        scalar: string
    - name: parameters
      type:
        namedType: io.k8s.api.networking.v1beta1.IngressClassParametersReference
- name: io.k8s.api.networking.v1beta1.IngressRule
  map:
    fields:
    - name: host
      type:
        scalar: string
    - name: http
      type:
        namedType: io.k8s.api.networking.v1beta1.HTTPIngressRuleValue
- name: io.k8s.api.networking.v1beta1.IngressSpec
  map:
    fields:
    - name: backend
      type:
        namedType: io.k8s.api.networking.v1beta1.IngressBackend
    - name: ingressClassName
      type:
        scalar: string
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1beta1.IngressRule
          elementRelationship: atomic
    - name: tls
      type:
        list:
          elementType:
            namedType: io.k8s.api.networking.v1beta1.IngressTLS
          elementRelationship: atomic
- name: io.k8s.api.networking.v1beta1.IngressStatus
  map:
    fields:
    - name: loadBalancer
      type:
        namedType: io.k8s.api.core.v1.LoadBalancerStatus
      default: {}
- name: io.k8s.api.networking.v1beta1.IngressTLS
  map:
    fields:
    - name: hosts
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: secretName
      type:
        scalar: string
- name: io.k8s.api.node.v1.Overhead
  map:
    fields:
    - name: podFixed
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.node.v1.RuntimeClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: handler
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: overhead
      type:
        namedType: io.k8s.api.node.v1.Overhead
    - name: scheduling
      type:
        namedType: io.k8s.api.node.v1.Scheduling
- name: io.k8s.api.node.v1.Scheduling
  map:
    fields:
    - name: nodeSelector
      type:
        map:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: tolerations
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Toleration
          elementRelationship: atomic
- name: io.k8s.api.node.v1alpha1.Overhead
  map:
    fields:
    - name: podFixed
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.node.v1alpha1.RuntimeClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.node.v1alpha1.RuntimeClassSpec
      default: {}
- name: io.k8s.api.node.v1alpha1.RuntimeClassSpec
  map:
    fields:
    - name: overhead
      type:
        namedType: io.k8s.api.node.v1alpha1.Overhead
    - name: runtimeHandler
      type:
        scalar: string
      default: ""
    - name: scheduling
      type:
        namedType: io.k8s.api.node.v1alpha1.Scheduling
- name: io.k8s.api.node.v1alpha1.Scheduling
  map:
    fields:
    - name: nodeSelector
      type:
        map:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: tolerations
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Toleration
          elementRelationship: atomic
- name: io.k8s.api.node.v1beta1.Overhead
  map:
    fields:
    - name: podFixed
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
- name: io.k8s.api.node.v1beta1.RuntimeClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: handler
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: overhead
      type:
        namedType: io.k8s.api.node.v1beta1.Overhead
    - name: scheduling
      type:
        namedType: io.k8s.api.node.v1beta1.Scheduling
- name: io.k8s.api.node.v1beta1.Scheduling
  map:
    fields:
    - name: nodeSelector
      type:
        map:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: tolerations
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.Toleration
          elementRelationship: atomic
- name: io.k8s.api.policy.v1.Eviction
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: deleteOptions
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
- name: io.k8s.api.policy.v1.PodDisruptionBudget
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.policy.v1.PodDisruptionBudgetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.policy.v1.PodDisruptionBudgetStatus
      default: {}
- name: io.k8s.api.policy.v1.PodDisruptionBudgetSpec
  map:
    fields:
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: minAvailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
- name: io.k8s.api.policy.v1.PodDisruptionBudgetStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: currentHealthy
      type:
        scalar: numeric
      default: 0
    - name: desiredHealthy
      type:
        scalar: numeric
      default: 0
    - name: disruptedPods
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: disruptionsAllowed
      type:
        scalar: numeric
      default: 0
    - name: expectedPods
      type:
        scalar: numeric
      default: 0
    - name: observedGeneration
      type:
        scalar: numeric
- name: io.k8s.api.policy.v1beta1.AllowedCSIDriver
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.policy.v1beta1.AllowedFlexVolume
  map:
    fields:
    - name: driver
      type:
        scalar: string
      default: ""
- name: io.k8s.api.policy.v1beta1.AllowedHostPath
  map:
    fields:
    - name: pathPrefix
      type:
        scalar: string
    - name: readOnly
      type:
        scalar: boolean
- name: io.k8s.api.policy.v1beta1.Eviction
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: deleteOptions
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
- name: io.k8s.api.policy.v1beta1.FSGroupStrategyOptions
  map:
    fields:
    - name: ranges
      type:
        list:
          elementType:
            namedType: io.k8s.api.policy.v1beta1.IDRange
          elementRelationship: atomic
    - name: rule
      type:
        scalar: string
- name: io.k8s.api.policy.v1beta1.HostPortRange
  map:
    fields:
    - name: max
      type:
        scalar: numeric
      default: 0
    - name: min
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.policy.v1beta1.IDRange
  map:
    fields:
    - name: max
      type:
        scalar: numeric
      default: 0
    - name: min
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.policy.v1beta1.PodDisruptionBudget
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.policy.v1beta1.PodDisruptionBudgetSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.policy.v1beta1.PodDisruptionBudgetStatus
      default: {}
- name: io.k8s.api.policy.v1beta1.PodDisruptionBudgetSpec
  map:
    fields:
    - name: maxUnavailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: minAvailable
      type:
        namedType: io.k8s.apimachinery.pkg.util.intstr.IntOrString
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
- name: io.k8s.api.policy.v1beta1.PodDisruptionBudgetStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: currentHealthy
      type:
        scalar: numeric
      default: 0
    - name: desiredHealthy
      type:
        scalar: numeric
      default: 0
    - name: disruptedPods
      type:
        map:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: disruptionsAllowed
      type:
        scalar: numeric
      default: 0
    - name: expectedPods
      type:
        scalar: numeric
      default: 0
    - name: observedGeneration
      type:
        scalar: numeric
- name: io.k8s.api.policy.v1beta1.PodSecurityPolicy
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.policy.v1beta1.PodSecurityPolicySpec
      default: {}
- name: io.k8s.api.policy.v1beta1.PodSecurityPolicySpec
  map:
    fields:
    - name: allowPrivilegeEscalation
      type:
        scalar: boolean
    - name: allowedCSIDrivers
      type:
        list:
          elementType:
            namedType: io.k8s.api.policy.v1beta1.AllowedCSIDriver
          elementRelationship: atomic
    - name: allowedCapabilities
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: allowedFlexVolumes
      type:
        list:
          elementType:
            namedType: io.k8s.api.policy.v1beta1.AllowedFlexVolume
          elementRelationship: atomic
    - name: allowedHostPaths
      type:
        list:
          elementType:
            namedType: io.k8s.api.policy.v1beta1.AllowedHostPath
          elementRelationship: atomic
    - name: allowedProcMountTypes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: allowedUnsafeSysctls
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: defaultAddCapabilities
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: defaultAllowPrivilegeEscalation
      type:
        scalar: boolean
    - name: forbiddenSysctls
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: fsGroup
      type:
        namedType: io.k8s.api.policy.v1beta1.FSGroupStrategyOptions
      default: {}
    - name: hostIPC
      type:
        scalar: boolean
    - name: hostNetwork
      type:
        scalar: boolean
    - name: hostPID
      type:
        scalar: boolean
    - name: hostPorts
      type:
        list:
          elementType:
            namedType: io.k8s.api.policy.v1beta1.HostPortRange
          elementRelationship: atomic
    - name: privileged
      type:
        scalar: boolean
    - name: readOnlyRootFilesystem
      type:
        scalar: boolean
    - name: requiredDropCapabilities
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: runAsGroup
      type:
        namedType: io.k8s.api.policy.v1beta1.RunAsGroupStrategyOptions
    - name: runAsUser
      type:
        namedType: io.k8s.api.policy.v1beta1.RunAsUserStrategyOptions
      default: {}
    - name: runtimeClass
      type:
        namedType: io.k8s.api.policy.v1beta1.RuntimeClassStrategyOptions
    - name: seLinux
      type:
        namedType: io.k8s.api.policy.v1beta1.SELinuxStrategyOptions
      default: {}
    - name: supplementalGroups
      type:
        namedType: io.k8s.api.policy.v1beta1.SupplementalGroupsStrategyOptions
      default: {}
    - name: volumes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.policy.v1beta1.RunAsGroupStrategyOptions
  map:
    fields:
    - name: ranges
      type:
        list:
          elementType:
            namedType: io.k8s.api.policy.v1beta1.IDRange
          elementRelationship: atomic
    - name: rule
      type:
        scalar: string
      default: ""
- name: io.k8s.api.policy.v1beta1.RunAsUserStrategyOptions
  map:
    fields:
    - name: ranges
      type:
        list:
          elementType:
            namedType: io.k8s.api.policy.v1beta1.IDRange
          elementRelationship: atomic
    - name: rule
      type:
        scalar: string
      default: ""
- name: io.k8s.api.policy.v1beta1.RuntimeClassStrategyOptions
  map:
    fields:
    - name: allowedRuntimeClassNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: defaultRuntimeClassName
      type:
        scalar: string
- name: io.k8s.api.policy.v1beta1.SELinuxStrategyOptions
  map:
    fields:
    - name: rule
      type:
        scalar: string
      default: ""
    - name: seLinuxOptions
      type:
        namedType: io.k8s.api.core.v1.SELinuxOptions
- name: io.k8s.api.policy.v1beta1.SupplementalGroupsStrategyOptions
  map:
    fields:
    - name: ranges
      type:
        list:
          elementType:
            namedType: io.k8s.api.policy.v1beta1.IDRange
          elementRelationship: atomic
    - name: rule
      type:
        scalar: string
- name: io.k8s.api.rbac.v1.AggregationRule
  map:
    fields:
    - name: clusterRoleSelectors
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1.ClusterRole
  map:
    fields:
    - name: aggregationRule
      type:
        namedType: io.k8s.api.rbac.v1.AggregationRule
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1.PolicyRule
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1.ClusterRoleBinding
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: roleRef
      type:
        namedType: io.k8s.api.rbac.v1.RoleRef
      default: {}
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1.Subject
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1.PolicyRule
  map:
    fields:
    - name: apiGroups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: nonResourceURLs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resourceNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resources
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1.Role
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1.PolicyRule
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1.RoleBinding
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: roleRef
      type:
        namedType: io.k8s.api.rbac.v1.RoleRef
      default: {}
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1.Subject
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1.RoleRef
  map:
    fields:
    - name: apiGroup
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.api.rbac.v1.Subject
  map:
    fields:
    - name: apiGroup
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    elementRelationship: atomic
- name: io.k8s.api.rbac.v1alpha1.AggregationRule
  map:
    fields:
    - name: clusterRoleSelectors
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1alpha1.ClusterRole
  map:
    fields:
    - name: aggregationRule
      type:
        namedType: io.k8s.api.rbac.v1alpha1.AggregationRule
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1alpha1.PolicyRule
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1alpha1.ClusterRoleBinding
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: roleRef
      type:
        namedType: io.k8s.api.rbac.v1alpha1.RoleRef
      default: {}
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1alpha1.Subject
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1alpha1.PolicyRule
  map:
    fields:
    - name: apiGroups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: nonResourceURLs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resourceNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resources
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1alpha1.Role
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1alpha1.PolicyRule
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1alpha1.RoleBinding
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: roleRef
      type:
        namedType: io.k8s.api.rbac.v1alpha1.RoleRef
      default: {}
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1alpha1.Subject
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1alpha1.RoleRef
  map:
    fields:
    - name: apiGroup
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.rbac.v1alpha1.Subject
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
- name: io.k8s.api.rbac.v1beta1.AggregationRule
  map:
    fields:
    - name: clusterRoleSelectors
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1beta1.ClusterRole
  map:
    fields:
    - name: aggregationRule
      type:
        namedType: io.k8s.api.rbac.v1beta1.AggregationRule
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1beta1.PolicyRule
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1beta1.ClusterRoleBinding
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: roleRef
      type:
        namedType: io.k8s.api.rbac.v1beta1.RoleRef
      default: {}
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1beta1.Subject
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1beta1.PolicyRule
  map:
    fields:
    - name: apiGroups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: nonResourceURLs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resourceNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resources
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1beta1.Role
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1beta1.PolicyRule
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1beta1.RoleBinding
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: roleRef
      type:
        namedType: io.k8s.api.rbac.v1beta1.RoleRef
      default: {}
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.rbac.v1beta1.Subject
          elementRelationship: atomic
- name: io.k8s.api.rbac.v1beta1.RoleRef
  map:
    fields:
    - name: apiGroup
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.api.rbac.v1beta1.Subject
  map:
    fields:
    - name: apiGroup
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
- name: io.k8s.api.scheduling.v1.PriorityClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: description
      type:
        scalar: string
    - name: globalDefault
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: preemptionPolicy
      type:
        scalar: string
    - name: value
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.scheduling.v1alpha1.PriorityClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: description
      type:
        scalar: string
    - name: globalDefault
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: preemptionPolicy
      type:
        scalar: string
    - name: value
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.scheduling.v1beta1.PriorityClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: description
      type:
        scalar: string
    - name: globalDefault
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: preemptionPolicy
      type:
        scalar: string
    - name: value
      type:
        scalar: numeric
      default: 0
- name: io.k8s.api.storage.v1.CSIDriver
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.storage.v1.CSIDriverSpec
      default: {}
- name: io.k8s.api.storage.v1.CSIDriverSpec
  map:
    fields:
    - name: attachRequired
      type:
        scalar: boolean
    - name: fsGroupPolicy
      type:
        scalar: string
    - name: podInfoOnMount
      type:
        scalar: boolean
    - name: requiresRepublish
      type:
        scalar: boolean
    - name: storageCapacity
      type:
        scalar: boolean
    - name: tokenRequests
      type:
        list:
          elementType:
            namedType: io.k8s.api.storage.v1.TokenRequest
          elementRelationship: atomic
    - name: volumeLifecycleModes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: io.k8s.api.storage.v1.CSINode
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.storage.v1.CSINodeSpec
      default: {}
- name: io.k8s.api.storage.v1.CSINodeDriver
  map:
    fields:
    - name: allocatable
      type:
        namedType: io.k8s.api.storage.v1.VolumeNodeResources
    - name: name
      type:
        scalar: string
      default: ""
    - name: nodeID
      type:
        scalar: string
      default: ""
    - name: topologyKeys
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.storage.v1.CSINodeSpec
  map:
    fields:
    - name: drivers
      type:
        list:
          elementType:
            namedType: io.k8s.api.storage.v1.CSINodeDriver
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.api.storage.v1.CSIStorageCapacity
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: capacity
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: kind
      type:
        scalar: string
    - name: maximumVolumeSize
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: nodeTopology
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: storageClassName
      type:
        scalar: string
      default: ""
- name: io.k8s.api.storage.v1.StorageClass
  map:
    fields:
    - name: allowVolumeExpansion
      type:
        scalar: boolean
    - name: allowedTopologies
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.TopologySelectorTerm
          elementRelationship: atomic
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: mountOptions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: parameters
      type:
        map:
          elementType:
            scalar: string
    - name: provisioner
      type:
        scalar: string
      default: ""
    - name: reclaimPolicy
      type:
        scalar: string
    - name: volumeBindingMode
      type:
        scalar: string
- name: io.k8s.api.storage.v1.TokenRequest
  map:
    fields:
    - name: audience
      type:
        scalar: string
      default: ""
    - name: expirationSeconds
      type:
        scalar: numeric
- name: io.k8s.api.storage.v1.VolumeAttachment
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.storage.v1.VolumeAttachmentSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.storage.v1.VolumeAttachmentStatus
      default: {}
- name: io.k8s.api.storage.v1.VolumeAttachmentSource
  map:
    fields:
    - name: inlineVolumeSpec
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeSpec
    - name: persistentVolumeName
      type:
        scalar: string
- name: io.k8s.api.storage.v1.VolumeAttachmentSpec
  map:
    fields:
    - name: attacher
      type:
        scalar: string
      default: ""
    - name: nodeName
      type:
        scalar: string
      default: ""
    - name: source
      type:
        namedType: io.k8s.api.storage.v1.VolumeAttachmentSource
      default: {}
- name: io.k8s.api.storage.v1.VolumeAttachmentStatus
  map:
    fields:
    - name: attachError
      type:
        namedType: io.k8s.api.storage.v1.VolumeError
    - name: attached
      type:
        scalar: boolean
      default: false
    - name: attachmentMetadata
      type:
        map:
          elementType:
            scalar: string
    - name: detachError
      type:
        namedType: io.k8s.api.storage.v1.VolumeError
- name: io.k8s.api.storage.v1.VolumeError
  map:
    fields:
    - name: message
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
- name: io.k8s.api.storage.v1.VolumeNodeResources
  map:
    fields:
    - name: count
      type:
        scalar: numeric
- name: io.k8s.api.storage.v1alpha1.CSIStorageCapacity
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: capacity
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: kind
      type:
        scalar: string
    - name: maximumVolumeSize
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: nodeTopology
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: storageClassName
      type:
        scalar: string
      default: ""
- name: io.k8s.api.storage.v1alpha1.VolumeAttachment
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.storage.v1alpha1.VolumeAttachmentSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.storage.v1alpha1.VolumeAttachmentStatus
      default: {}
- name: io.k8s.api.storage.v1alpha1.VolumeAttachmentSource
  map:
    fields:
    - name: inlineVolumeSpec
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeSpec
    - name: persistentVolumeName
      type:
        scalar: string
- name: io.k8s.api.storage.v1alpha1.VolumeAttachmentSpec
  map:
    fields:
    - name: attacher
      type:
        scalar: string
      default: ""
    - name: nodeName
      type:
        scalar: string
      default: ""
    - name: source
      type:
        namedType: io.k8s.api.storage.v1alpha1.VolumeAttachmentSource
      default: {}
- name: io.k8s.api.storage.v1alpha1.VolumeAttachmentStatus
  map:
    fields:
    - name: attachError
      type:
        namedType: io.k8s.api.storage.v1alpha1.VolumeError
    - name: attached
      type:
        scalar: boolean
      default: false
    - name: attachmentMetadata
      type:
        map:
          elementType:
            scalar: string
    - name: detachError
      type:
        namedType: io.k8s.api.storage.v1alpha1.VolumeError
- name: io.k8s.api.storage.v1alpha1.VolumeError
  map:
    fields:
    - name: message
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
- name: io.k8s.api.storage.v1beta1.CSIDriver
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.storage.v1beta1.CSIDriverSpec
      default: {}
- name: io.k8s.api.storage.v1beta1.CSIDriverSpec
  map:
    fields:
    - name: attachRequired
      type:
        scalar: boolean
    - name: fsGroupPolicy
      type:
        scalar: string
    - name: podInfoOnMount
      type:
        scalar: boolean
    - name: requiresRepublish
      type:
        scalar: boolean
    - name: storageCapacity
      type:
        scalar: boolean
    - name: tokenRequests
      type:
        list:
          elementType:
            namedType: io.k8s.api.storage.v1beta1.TokenRequest
          elementRelationship: atomic
    - name: volumeLifecycleModes
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.storage.v1beta1.CSINode
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.storage.v1beta1.CSINodeSpec
      default: {}
- name: io.k8s.api.storage.v1beta1.CSINodeDriver
  map:
    fields:
    - name: allocatable
      type:
        namedType: io.k8s.api.storage.v1beta1.VolumeNodeResources
    - name: name
      type:
        scalar: string
      default: ""
    - name: nodeID
      type:
        scalar: string
      default: ""
    - name: topologyKeys
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.storage.v1beta1.CSINodeSpec
  map:
    fields:
    - name: drivers
      type:
        list:
          elementType:
            namedType: io.k8s.api.storage.v1beta1.CSINodeDriver
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.api.storage.v1beta1.CSIStorageCapacity
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: capacity
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: kind
      type:
        scalar: string
    - name: maximumVolumeSize
      type:
        namedType: io.k8s.apimachinery.pkg.api.resource.Quantity
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: nodeTopology
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
    - name: storageClassName
      type:
        scalar: string
      default: ""
- name: io.k8s.api.storage.v1beta1.StorageClass
  map:
    fields:
    - name: allowVolumeExpansion
      type:
        scalar: boolean
    - name: allowedTopologies
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.TopologySelectorTerm
          elementRelationship: atomic
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: mountOptions
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: parameters
      type:
        map:
          elementType:
            scalar: string
    - name: provisioner
      type:
        scalar: string
      default: ""
    - name: reclaimPolicy
      type:
        scalar: string
    - name: volumeBindingMode
      type:
        scalar: string
- name: io.k8s.api.storage.v1beta1.TokenRequest
  map:
    fields:
    - name: audience
      type:
        scalar: string
      default: ""
    - name: expirationSeconds
      type:
        scalar: numeric
- name: io.k8s.api.storage.v1beta1.VolumeAttachment
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.api.storage.v1beta1.VolumeAttachmentSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.api.storage.v1beta1.VolumeAttachmentStatus
      default: {}
- name: io.k8s.api.storage.v1beta1.VolumeAttachmentSource
  map:
    fields:
    - name: inlineVolumeSpec
      type:
        namedType: io.k8s.api.core.v1.PersistentVolumeSpec
    - name: persistentVolumeName
      type:
        scalar: string
- name: io.k8s.api.storage.v1beta1.VolumeAttachmentSpec
  map:
    fields:
    - name: attacher
      type:
        scalar: string
      default: ""
    - name: nodeName
      type:
        scalar: string
      default: ""
    - name: source
      type:
        namedType: io.k8s.api.storage.v1beta1.VolumeAttachmentSource
      default: {}
- name: io.k8s.api.storage.v1beta1.VolumeAttachmentStatus
  map:
    fields:
    - name: attachError
      type:
        namedType: io.k8s.api.storage.v1beta1.VolumeError
    - name: attached
      type:
        scalar: boolean
      default: false
    - name: attachmentMetadata
      type:
        map:
          elementType:
            scalar: string
    - name: detachError
      type:
        namedType: io.k8s.api.storage.v1beta1.VolumeError
- name: io.k8s.api.storage.v1beta1.VolumeError
  map:
    fields:
    - name: message
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
- name: io.k8s.api.storage.v1beta1.VolumeNodeResources
  map:
    fields:
    - name: count
      type:
        scalar: numeric
- name: io.k8s.apimachinery.pkg.api.resource.Quantity
  scalar: untyped
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: message
      type:
        scalar: string
      default: ""
    - name: observedGeneration
      type:
        scalar: numeric
    - name: reason
      type:
        scalar: string
      default: ""
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.apimachinery.pkg.apis.meta.v1.DeleteOptions
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: dryRun
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: gracePeriodSeconds
      type:
        scalar: numeric
    - name: kind
      type:
        scalar: string
    - name: orphanDependents
      type:
        scalar: boolean
    - name: preconditions
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Preconditions
    - name: propagationPolicy
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
  map:
    fields:
    - name: matchExpressions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
          elementRelationship: atomic
    - name: matchLabels
      type:
        map:
          elementType:
            scalar: string
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: operator
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldsType
      type:
        scalar: string
    - name: fieldsV1
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
    - name: manager
      type:
        scalar: string
    - name: operation
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
  scalar: untyped
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: clusterName
      type:
        scalar: string
    - name: creationTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
      default: {}
    - name: deletionGracePeriodSeconds
      type:
        scalar: numeric
    - name: deletionTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: generateName
      type:
        scalar: string
    - name: generation
      type:
        scalar: numeric
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: managedFields
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
    - name: resourceVersion
      type:
        scalar: string
    - name: selfLink
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
      default: ""
    - name: blockOwnerDeletion
      type:
        scalar: boolean
    - name: controller
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Preconditions
  map:
    fields:
    - name: resourceVersion
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Time
  scalar: untyped
- name: io.k8s.apimachinery.pkg.runtime.RawExtension
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.util.intstr.IntOrString
  scalar: untyped
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
