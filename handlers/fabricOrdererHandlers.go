package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/coffeloop/go-fabric5/models"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateFabricOrderer(c *gin.Context) {
	var params models.OrdererOptions

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "ordnode", "create",
		"--image="+params.Image,
		"--version="+params.Version,
		"--storage-class="+params.StorageClass,
		"--enroll-id="+params.EnrollID,
		"--mspid="+params.MSPID,
		"--enroll-pw="+params.EnrollPW,
		"--capacity="+params.Capacity,
		"--name="+params.Name,
		"--ca-name="+params.CAName,
		"--hosts="+params.Hosts,
		"--istio-port="+strconv.Itoa(params.IstioPort))

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fabric Orderer creado exitosamente",
	})
}

func CheckFabricOrdererStatus(c *gin.Context) {
	var params models.CheckOrdererOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ordName := params.Name
	ns := params.Namespace

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Kubernetes client configuration: " + err.Error()})
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Kubernetes client: " + err.Error()})
		return
	}

	pods, err := clientset.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{LabelSelector: fmt.Sprintf("release=%s", ordName)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list pods: " + err.Error()})
		return
	}

	var podStatus *v1.PodPhase
	for _, pod := range pods.Items {
		podStatus = &pod.Status.Phase
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("CA %s is %v in namespace %s", ordName, *podStatus, ns),
			"pod":     pod.Name,
			"status":  *podStatus,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf("CA %s not found in namespace %s", ordName, ns),
		"status":  podStatus,
	})
}
