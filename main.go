package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-plugin"
	shared "github.com/orkestra-io/orkestra-shared"
)

// PluginExecutor est la structure qui implémente l'interface NodeExecutor.
// C'est ici que vit la logique de votre plugin.
type MyPluginExecutor struct{}

// GetCapabilities annonce les types de nœuds que ce plugin fournit.
func (c *MyPluginExecutor) GetCapabilities() ([]string, error) {
	// TODO: Mettez à jour cette liste avec les nœuds que votre plugin implémente.
	// Exemple : "discord/send-message", "aws/s3/upload", etc.
	return []string{
		"template/hello",
	}, nil
}

// Execute est la fonction principale appelée par le moteur Orkestra.
func (c *MyPluginExecutor) Execute(node shared.Node, ctx shared.ExecutionContext) (interface{}, error) {
	// Le 'With' contient les paramètres de votre nœud, déjà résolus.
	withMap := node.With

	// TODO: Ajoutez ici la logique pour vos différents nœuds.
	switch node.Uses {
	case "template/hello":
		return executeEcho(withMap)

	// TODO: Ajoutez des 'case' pour vos autres nœuds ici.

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
			"executor": &shared.NodeExecutorPlugin{Impl: &MyPluginExecutor{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

// --- Logique des Nœuds ---

// executeEcho est un exemple d'implémentation pour un nœud "template/echo".
// Il renvoie simplement le message qu'il reçoit en paramètre.
func executeEcho(with map[string]interface{}) (interface{}, error) {
	message, ok := with["message"].(string)
	if !ok {
		return nil, fmt.Errorf("le paramètre 'message' est requis et doit être une chaîne pour le nœud echo")
	}

	// Affiche un log qui sera visible dans Orkestra
	log.Printf("LOG | Le nœud 'echo' a reçu le message : %s", message)

	// La sortie du nœud sera accessible via `nodes.votre_id.output`
	return map[string]interface{}{
		"output": fmt.Sprintf("Le plugin a reçu: %s", message),
	}, nil
}
