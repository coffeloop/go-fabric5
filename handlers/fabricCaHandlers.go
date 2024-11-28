package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/coffeloop/go-fabric5/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateFabricCA(c *gin.Context) {
	// Obtener parámetros del body
	var params models.CreateCAOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Construir el comando
	cmd := exec.Command("kubectl", "hlf", "ca", "create",
		"--image="+params.Image,
		"--version="+params.Version,
		"--storage-class="+params.StorageClass,
		"--capacity="+params.Capacity,
		"--name="+params.Name,
		"--enroll-id="+params.EnrollID,
		"--enroll-pw="+params.EnrollPW,
		"--hosts="+strings.Join(params.Hosts, ","),
		"--istio-port="+strconv.Itoa(params.IstioPort))

	log.Infof("Comando a ejecutar: %s", strings.Join(cmd.Args, " "))

	// Ejecutar el comando
	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fabric CA creado exitosamente",
	})
}

func RegisterFabricCA(c *gin.Context) {
	var params models.RegisterCAOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "ca", "register",
		"--name="+params.Name,
		"--user="+params.User,
		"--secret="+params.Secret,
		"--type="+params.Type,
		"--enroll-id="+params.EnrollID,
		"--enroll-secret="+params.EnrollSecret,
		"--mspid="+params.MSPID)

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fabric CA registrado exitosamente",
	})
}

func EnrollFabricCA(c *gin.Context) {
	var params models.EnrollCAOptions

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "ca", "enroll",
		"--name="+params.Name,
		"--namespace="+params.Namespace,
		"--user="+params.User,
		"--secret="+params.Secret,
		"--mspid="+params.MSPID,
		"--ca-name="+params.CAName,
		"--output="+params.Output)

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fabric CA enrolled successfully",
		"output":  string(output),
	})

}

func RegisterUserFabricCA(c *gin.Context) {
	var params models.RegisterUserCAOptions

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "ca", "register",
		"--name="+params.Name,
		"--user="+params.User,
		"--secret="+params.Secret,
		"--type="+params.Type,
		"--enroll-id="+params.EnrollID,
		"--enroll-secret="+params.EnrollSecret,
		"--mspid="+params.MSPID)

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario registrado exitosamente en Fabric CA",
	})
}

func CheckCreateFabricCA(c *gin.Context) {
	cmd := exec.Command("kubectl", "get", "fabriccas.hlf.kungfusoftware.es", "--all-namespaces")

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	status := string(output)
	if strings.Contains(status, "RUNNING") {
		c.JSON(http.StatusOK, gin.H{
			"message": "Todas las Fabric CA están en ejecución",
			"status":  status,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Alguna o todas las Fabric CA no se encontraron o no están en ejecución",
			"status":  status,
		})
	}
}

func CheckFabricCAStatus(c *gin.Context) {
	var params models.CheckCAOptions
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	caName := params.Name
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

	pods, err := clientset.CoreV1().Pods(ns).List(context.Background(), metav1.ListOptions{LabelSelector: fmt.Sprintf("release=%s", caName)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list pods: " + err.Error()})
		return
	}

	var podStatus *v1.PodPhase
	for _, pod := range pods.Items {
		podStatus = &pod.Status.Phase
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("CA %s is %v in namespace %s", caName, *podStatus, ns),
			"pod":     pod.Name,
			"status":  *podStatus,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": fmt.Sprintf("CA %s not found in namespace %s", caName, ns),
		"status":  podStatus,
	})
}
