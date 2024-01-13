/*
Copyright 2024.

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

package controller

import (
	"context"
	"encoding/xml"
	"github.com/digitalocean/go-libvirt"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"strconv"
	virtlitev1alpha1 "virtlite/api/v1alpha1"
)

type VirtualMachineReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	VirtConn *libvirt.Libvirt
}

func (r *VirtualMachineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var vm virtlitev1alpha1.VirtualMachine
	if err := r.Get(ctx, req.NamespacedName, &vm); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var vmPod corev1.Pod
	vmPodKey := types.NamespacedName{
		Name:      vm.Name,
		Namespace: vm.Namespace,
	}
	err := r.Get(ctx, vmPodKey, &vmPod)
	if apierrors.IsNotFound(err) {
		//pod是为了使用k8s的调度功能,并不实际工作,提供网络和存储,想办法连接到虚拟机上
		//程序以daemonset的方式运行在每个节点上,pod调度到哪个节点上,虚拟机才在哪个节点上创建
		//可以想办法把虚拟机的实际cpu内存更新到pod上,这样k8s的监控也就可以工作了
		pod, err := r.buildVMPod(ctx, &vm)
		if err != nil {
			return ctrl.Result{}, err
		}
		if err := r.Create(ctx, pod); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	xmlBytes, err := xml.MarshalIndent(vm.Spec, "", "  ")
	if err != nil {
		return ctrl.Result{}, err
	}
	if vm.Status.Phase == "" {
		_, err = r.VirtConn.DomainCreateXML(string(xmlBytes), 0)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	vm.Status.Phase = "Running"
	if err := r.Status().Update(ctx, &vm); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *VirtualMachineReconciler) buildVMPod(ctx context.Context, vm *virtlitev1alpha1.VirtualMachine) (*corev1.Pod, error) {
	cpu := strconv.Itoa(int(vm.Spec.VCPU.CPUs))
	memory := strconv.FormatUint(vm.Spec.Memory.Value*1024, 10)

	vmPod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels:      vm.Labels,
			Annotations: vm.Annotations,
			Name:        vm.Name,
			Namespace:   vm.Namespace,
		},
		Spec: corev1.PodSpec{
			RestartPolicy: corev1.RestartPolicyNever,
			Containers: []corev1.Container{{
				Resources: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse(cpu),
						corev1.ResourceMemory: resource.MustParse(memory),
					},
					Limits: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse(cpu),
						corev1.ResourceMemory: resource.MustParse(memory),
					},
				},
				Name:    "holder",
				Image:   "busybox",
				Command: []string{"sh", "-c", "while true; do echo 'Still alive'; sleep 60; done"},
			}},
		},
	}
	if err := controllerutil.SetControllerReference(vm, &vmPod, r.Scheme); err != nil {
		return nil, err
	}

	return vmPod.DeepCopy(), nil
}

func (r *VirtualMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&virtlitev1alpha1.VirtualMachine{}).
		Watches(&corev1.Pod{}, handler.EnqueueRequestsFromMapFunc(func(ctx context.Context, obj client.Object) []reconcile.Request {
			pod := obj.(*corev1.Pod)
			controllerRef := metav1.GetControllerOf(obj)
			if controllerRef == nil || controllerRef.Kind != "VirtualMachine" {
				return nil
			}
			//todo 获取节点名称
			if pod.Spec.NodeName != "server" {
				return nil
			}
			return []reconcile.Request{{
				NamespacedName: types.NamespacedName{
					Namespace: obj.GetNamespace(),
					Name:      controllerRef.Name,
				},
			}}
		})).
		Complete(r)
}
