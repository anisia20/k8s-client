package routes

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Getpods get all pods
func Getpods(c *gin.Context) {
	// kubeconfig에서 현재 콘텍스트를 사용한다
	// path-to-kubeconfig -- 예를 들어, /root/.kube/config
	config, _ := clientcmd.BuildConfigFromFlags("", "/Users/mz01-josm/.kube/config")
	// clientset을 생성한다
	clientset, _ := kubernetes.NewForConfig(config)
	//test, err := ""

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	podnames := []string{} //sliece
	for i := 0; i < len(pods.Items); i++ {
		fmt.Printf("Found pod %s \n", pods.Items[i].GetObjectMeta().GetName())
		podnames = append(podnames, pods.Items[i].GetObjectMeta().GetName())
	}
	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// namespace := ""
	// pod := "kuard"
	// _, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	// if errors.IsNotFound(err) {
	// 	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	// 		pod, namespace, statusError.ErrStatus.Message)
	// } else if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	// }
	//time.Sleep(10 * time.Second)

	c.JSON(200, podnames)
}
