package setters2

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"sigs.k8s.io/kustomize/kyaml/openapi"
	"sigs.k8s.io/kustomize/kyaml/yaml"
	"strings"
	"testing"
)

func TestDelete_Filter(t *testing.T) {
	var tests = []struct {
		name        string
		description string
		setter      string
		openapi     string
		input       string
		expected_output    string
		expected_schema    string
	}{
		{
			name:   "set-replicas",
			setter: "replicas",
			openapi: `
openAPI:
  definitions:
    io.k8s.cli.setters.no-match-1':
      x-k8s-cli:
        setter:
          name: no-match-1
          value: "1"
    io.k8s.cli.setters.replicas:
      x-k8s-cli:
        setter:
          name: replicas
          value: "4"
    io.k8s.cli.setters.no-match-2':
      x-k8s-cli:
        setter:
          name: no-match-2
          value: "2"
 `,
			input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3 # {"$ref": "#/definitions/io.k8s.cli.setters.replicas"}
 `,
			expected_output: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3 # {}
 `,
            expected_schema: `
openAPI:
  definitions:
    io.k8s.cli.setters.no-match-1':
      x-k8s-cli:
        setter:
          name: no-match-1
          value: "1"
    io.k8s.cli.setters.no-match-2':
      x-k8s-cli:
        setter:
          name: no-match-2
          value: "2"
 `,
		},
		{
			name:        "set-foo-type",
			description: "if a type is specified for a setter, ensure the field is of provided type",
			setter:      "foo",
			openapi: `
openAPI:
  definitions:
    io.k8s.cli.setters.foo:
      x-k8s-cli:
        setter:
          name: foo
          value: "4"
      type: integer
 `,
			input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  annotations:
    foo: 3 # {"$ref": "#/definitions/io.k8s.cli.setters.foo"}
 `,
			expected_output: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  annotations:
    foo: 3 # {}
 `,
			expected_schema: `
openAPI:
  definitions:
 `,
		},
		{
			name:   "set-replicas-enum",
			setter: "replicas",
			openapi: `
openAPI:
  definitions:
    io.k8s.cli.setters.no-match-1':
      x-k8s-cli:
        setter:
          name: no-match-1
          value: "1"
    io.k8s.cli.setters.replicas:
      x-k8s-cli:
        setter:
          name: replicas
          value: "medium"
          enumValues:
            small: "1"
            medium: "5"
            large: "50"
    io.k8s.cli.setters.no-match-2':
      x-k8s-cli:
        setter:
          name: no-match-2
          value: "2"
 `,
			input: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 1 # {"$ref": "#/definitions/io.k8s.cli.setters.replicas"}
 `,
			expected_output: `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 1 # {}
 `,
			expected_schema: `
openAPI:
  definitions:
    io.k8s.cli.setters.no-match-1':
      x-k8s-cli:
        setter:
          name: no-match-1
          value: "1"
    io.k8s.cli.setters.no-match-2':
      x-k8s-cli:
        setter:
          name: no-match-2
          value: "2"
 `,
		},
	}
	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			// reset the openAPI afterward
			defer openapi.ResetOpenAPI()
			initSchema(t, test.openapi)

			// parse the input to be modified
			r, err := yaml.Parse(test.input)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			// invoke the delete
			instance := &Delete{FieldName: test.setter}
			result, err := instance.Filter(r)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			// compare the actual and expected output
			actual, err := result.String()
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			actual = strings.TrimSpace(actual)
			expected := strings.TrimSpace(test.expected_output)
			if !assert.Equal(t, expected, actual) {
				t.FailNow()
			}
		})
	}
}

var resourcefile2 = `apiVersion: resource.dev/v1alpha1
kind: resourcefile
metadata:
    name: hello-world-set
upstream:
    type: git
    git:
        commit: 5c1c019b59299a4f6c7edd1ff5ff54d720621bbe
        directory: /package-examples/helloworld-set
        ref: v0.1.0
packageMetadata:
    shortDescription: example package using setters
openAPI:
  definitions:
    io.k8s.cli.setters.image:
      x-k8s-cli:
        setter:
          name: image
          value: "2"
`

func TestDelete_Filter2(t *testing.T) {
	path := filepath.Join(os.TempDir(), "resourcefile2")

	//write initial resourcefile to temp path
	err := ioutil.WriteFile(path, []byte(resourcefile2), 0666)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	//add a deleter definition
	dd := DeleterDefinition{
		Name:  "image",

	}

	err = dd.DeleteFromFile(path)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		t.FailNow()
	}


	expected := `apiVersion: resource.dev/v1alpha1
kind: resourcefile
metadata:
  name: hello-world-set
upstream:
  type: git
  git:
    commit: 5c1c019b59299a4f6c7edd1ff5ff54d720621bbe
    directory: /package-examples/helloworld-set
    ref: v0.1.0
packageMetadata:
  shortDescription: example package using setters
openAPI:
  definitions: {}
`
	assert.Equal(t, expected, string(b))
}
