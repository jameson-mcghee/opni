package gateway

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/rancher/opni/pkg/resources"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

func (r *Reconciler) serviceMonitor() resources.Resource {
	svcMonitor := &monitoringv1.ServiceMonitor{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "opni-gateway",
			Namespace: r.gw.Namespace,
			Labels:    resources.NewGatewayLabels(),
		},
		Spec: monitoringv1.ServiceMonitorSpec{
			Selector: metav1.LabelSelector{
				MatchLabels: resources.NewGatewayLabels(),
			},
			NamespaceSelector: monitoringv1.NamespaceSelector{
				MatchNames: []string{r.gw.Namespace},
			},
			Endpoints: []monitoringv1.Endpoint{
				{
					Port: "metrics",
				},
			},
		},
	}
	controllerutil.SetOwnerReference(r.gw, svcMonitor, r.client.Scheme())
	return resources.Present(svcMonitor)
}