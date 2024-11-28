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

func CreateFabricPeer(c *gin.Context) {
	var params models.CreatePeerOptions

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "peer", "create",
		"--statedb="+params.StateDB,
		"--image="+params.PeerImage,
		"--version="+params.PeerVersion,
		"--storage-class="+params.SCName,
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
		"message": "Fabric Peer creado exitosamente",
	})
}

func CheckFabricPeerStatus(c *gin.Context) {
	var params models.CheckPeerOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	peerName := params.Name
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

	pods, err := clientset.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{LabelSelector: fmt.Sprintf("release=%s", peerName)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list pods: " + err.Error()})
		return
	}

	var podStatus *v1.PodPhase
	for _, pod := range pods.Items {
		podStatus = &pod.Status.Phase
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Peer %s is %v in namespace %s", peerName, *podStatus, ns),
			"pod":     pod.Name,
			"status":  *podStatus,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf("Peer %s not found in namespace %s", peerName, ns),
		"status":  podStatus,
	})
}
