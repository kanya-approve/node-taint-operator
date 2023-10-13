package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apitypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestNodeReconcile(t *testing.T) {
	mockClient := new(MockClient)
	reconciler := &NodeReconciler{
		Client: mockClient,
	}

	nodeName := "test-node"
	node := &corev1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: nodeName,
		},
		Spec: corev1.NodeSpec{
			Taints: []corev1.Taint{
				{
					Key: "kubernetes.io/arch",
				},
			},
		},
	}

	// Mocking the Get and Update methods
	mockClient.On("Get", mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*corev1.Node)
		*arg = *node
	}).Return(nil)

	mockClient.On("Update", mock.Anything, mock.Anything).Return(nil)

	_, err := reconciler.Reconcile(context.TODO(), reconcile.Request{NamespacedName: apitypes.NamespacedName{Name: nodeName}})
	assert.NoError(t, err)

	// Verify if the taint was removed in the logic
	assert.Len(t, node.Spec.Taints, 1) // assuming there were 2 taints originally and 1 was removed
}
