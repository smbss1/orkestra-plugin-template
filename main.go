package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/go-plugin"
	"github.com/smbss1/orkestra-plugin-template/shared"
)

// PluginExecutor est la structure qui implémente l'interface NodeExecutor.
// C'est ici que vit la logique de votre plugin.
type PluginExecutor struct{}

// GetCapabilities annonce les types de nœuds que ce plugin fournit.
func (c *PluginExecutor) GetCapabilities() ([]string, error) {
	// TODO: Mettez à jour cette liste avec les nœuds que votre plugin implémente.
	// Exemple : "discord/send-message", "aws/s3/upload", etc.
	return []string{
		"template/hello",
	}, nil
}

// Execute est la fonction principale appelée par le moteur Orkestra.
func (c *PluginExecutor) Execute(node shared.Node, ctx shared.ExecutionContext) (interface{}, error) {
	// Désérialise les paramètres 'with' du nœud pour une utilisation facile.
	var withMap map[string]interface{}
	if err := json.Unmarshal(node.With, &withMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal 'with' params for node '%s': %w", node.ID, err)
	}

	// TODO: Ajoutez ici la logique pour vos différents nœuds.
	switch node.Uses {
	case "template/hello":
		// Exemple de logique de nœud
		name, ok := withMap["name"].(string)
		if !ok {
			name = "World"
		}
		message := fmt.Sprintf("Hello, %s!", name)
		log.Printf("LOG | %s", message) // Affiche un log qui sera visible dans Orkestra
		return map[string]interface{}{"message": message}, nil

	default:
		return nil, fmt.Errorf("type de noeud inconnu dans le plugin template: '%s'", node.Uses)
	}
}

// main est le point d'entrée du programme plugin.
// Ce code est standard et ne devrait pas avoir besoin d'être modifié.
func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"executor": &shared.NodeExecutorPlugin{Impl: &PluginExecutor{}},
		},
	})
}
