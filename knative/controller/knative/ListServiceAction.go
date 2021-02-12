package knative

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"test/knative/client"
	"test/knative/controller"
)

type ListServiceAction struct {
	controller.Action
	Namespace string
	Name      string
}

func (c *ListServiceAction) Process(ctx context.Context) interface{} {
	option := metav1.ListOptions{}
	if c.Name != "" {
		option.LabelSelector = fmt.Sprintf("name=%s", c.Name)
	}
	var service, err = client.GetClient().ServingClient.ServingV1().Services(c.Namespace).List(ctx, option)
	if err != nil || service == nil {
		log.Printf("create service error info: %v", err)
	}
	result := Result{}
	for i, e := range service.Items {
		result.List[i] = Item{Name: e.Name}
	}
	return result
}

type Result struct {
	List []Item
}
type Item struct {
	Name string
}
