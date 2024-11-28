package handlers

import (
	"net/http"
	"os/exec"

	"github.com/coffeloop/go-fabric5/models"
	"github.com/gin-gonic/gin"
)

func GetFabricConnectionChain(c *gin.Context) {
	var params models.ConnectionChain

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "inspect", "--output", params.Output, "-o", params.MSPID)

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fabric Connection Chain obtenido exitosamente",
		"output":  string(output),
	})
}

func AddUserToConnectionChain(c *gin.Context) {
	var params models.AddUserToConnectionChainOptions

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("kubectl", "hlf", "utils", "adduser",
		"--userPath="+params.UserPath,
		"--config="+params.Config,
		"--username="+params.Username,
		"--mspid="+params.MSPID)

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": string(output)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario agregado exitosamente a la cadena de conexión",
		"output":  string(output),
	})
}
