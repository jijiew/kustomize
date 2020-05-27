package settersutil

import (
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/openapi"
	"sigs.k8s.io/kustomize/kyaml/setters2"
)

// DeleterCreator delete a setter in the OpenAPI definitions, and removes references
// to the setter from matching resource fields.
type DeleterCreator struct {
	// Name is the name of the setter to create or update.
	Name string

	DeleteBy string

	SchemaPath string
}

func (d DeleterCreator) Delete(openAPIPath, resourcesPath string) error {
	schema, err := schemaFromFile(d.SchemaPath)
	if err != nil {
		return err
	}
	// Update the OpenAPI definitions to hace the setter
	dd := setters2.DeleterDefinition{
		Name: d.Name,
		DeleteBy: d.DeleteBy,
		Schema: schema,
	}
	if err := dd.DeleteFromFile(openAPIPath); err != nil {
		return err
	}

	// Load the updated definitions
	if err := openapi.AddSchemaFromFile(openAPIPath); err != nil {
		return err
	}

	// Update the resources with the deleter reference
	inout := &kio.LocalPackageReadWriter{PackagePath: resourcesPath}
	return kio.Pipeline{
		Inputs: []kio.Reader{inout},
		Filters: []kio.Filter{kio.FilterAll(
			&setters2.Delete{
				FieldName:  d.Name,
			})},
		Outputs: []kio.Writer{inout},
	}.Execute()
}