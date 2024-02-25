package controllers

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

type ContainerRequest struct {
	DockerImage string `json:"docker_image"`
}

type ContainerResponse struct {
	Message     string `json:"message"`
	ContainerID string `json:"container_id"`
}

func RunContainer(w http.ResponseWriter, r *http.Request) {
	var req ContainerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.DockerImage == "" {
		http.Error(w, "Parâmetro docker_image não fornecido", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("docker", "run", "-d", req.DockerImage)
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Erro ao criar e executar o container: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := ContainerResponse{
		Message:     "Container iniciado com sucesso a partir da imagem " + req.DockerImage,
		ContainerID: string(out),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
