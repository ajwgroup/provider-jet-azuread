package applicationcertificate

import "github.com/crossplane/terrajet/pkg/config"

// Configure : Provider Configuration
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_application_certificate", func(r *config.Resource) {

		r.ShortGroup = "applicationcertificate"
		r.ExternalName = config.IdentifierFromProvider

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["applicationcertificate"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-azuread/apis/applicationcertificate/v1alpha1.Certificate",
		}
	})
}
