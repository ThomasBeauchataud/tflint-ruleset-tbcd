package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint/rules"
	"github.com/terraform-linters/tflint/tflint"
	"os"
	"strings"
)

// VersionsFileRule vérifie la présence d'un fichier versions.tf avec les déclarations required_version
type VersionsFileRule struct {
	tflint.DefaultRule
}

// NewVersionsFileRule crée une nouvelle instance de VersionsFileRule
func NewVersionsFileRule() *VersionsFileRule {
	return &VersionsFileRule{}
}

// Name retourne le nom de la règle
func (r *VersionsFileRule) Name() string {
	return "versions_file_required"
}

// Enabled indique si la règle est activée par défaut
func (r *VersionsFileRule) Enabled() bool {
	return true
}

// Severity retourne la sévérité de la règle
func (r *VersionsFileRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link fournit un lien vers la documentation de la règle
func (r *VersionsFileRule) Link() string {
	return ""
}

// Check implémente la logique de la règle
func (r *VersionsFileRule) Check(runner tflint.Runner) error {
	// Recherche du fichier versions.tf
	files := runner.GetFiles("versions.tf")
	if len(files) == 0 {
		return nil // Aucun fichier versions.tf trouvé
	}

	// Lecture du contenu du fichier versions.tf
	filePath := files[0]
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("impossible de lire le fichier %s: %v", filePath, err)
	}

	// Vérification de la présence de required_version pour Terraform
	if !strings.Contains(string(content), "required_version") {
		runner.EmitIssue(
			r,
			filePath,
			"Le fichier versions.tf doit contenir la déclaration required_version pour Terraform",
		)
	}

	// Vérification de la présence de required_version pour les providers
	if !strings.Contains(string(content), "required_providers") {
		runner.EmitIssue(
			r,
			filePath,
			"Le fichier versions.tf doit contenir la déclaration required_version pour les providers",
		)
	}

	return nil
}
