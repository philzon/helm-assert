package yaml

import (
	"testing"

	"github.com/smallfish/simpleyaml"
)

const data = `
apiVersion: v1
kind: Pod
metadata:
  name: test-pod
  labels:
    app: test-app
spec:
  containers:
    - name: test-container
      image: nginx
      ports:
        - name: http
          containerPort: 80
        - name: udp
          containerPort: 25565
        - name: tcp
          containerPort: 9200
`

func TestLocateKey(t *testing.T) {
	tree, err := simpleyaml.NewYaml([]byte(data))

	key := "metadata.labels.app"

	if err != nil {
		t.Fatalf("Could not parse YAML: %s", err.Error())
	}

	tree, err = locate(key, tree)

	if err != nil {
		t.Errorf("Could not parse key: %s", err.Error())
	}

	if tree == nil {
		t.Errorf("Error: could not find key '%s'", key)
	}
}

func TestLocateKeyFromArrayFirst(t *testing.T) {
	tree, err := simpleyaml.NewYaml([]byte(data))

	key := "spec.containers[0].ports[0].name"

	if err != nil {
		t.Fatalf("could not parse YAML: %s", err.Error())
	}

	tree, err = locate(key, tree)

	if err != nil {
		t.Errorf("could not parse key: %s", err.Error())
	}

	if tree == nil {
		t.Errorf("could not find key '%s'", key)
	}
}

func TestLocateKeyFromArrayBounds(t *testing.T) {
	tree, err := simpleyaml.NewYaml([]byte(data))

	key := "spec.containers[0].ports[1].name"

	if err != nil {
		t.Fatalf("Could not parse YAML: %s", err.Error())
	}

	tree, err = locate(key, tree)

	if err != nil {
		t.Errorf("Could not parse key: %s", err.Error())
	}

	if tree == nil {
		t.Errorf("Error: could not find key '%s'", key)
	}
}

func TestLocateKeyFromArrayLast(t *testing.T) {
	tree, err := simpleyaml.NewYaml([]byte(data))

	key := "spec.containers[0].ports[2].name"

	if err != nil {
		t.Fatalf("Could not parse YAML: %s", err.Error())
	}

	tree, err = locate(key, tree)

	if err != nil {
		t.Errorf("Could not parse key: %s", err.Error())
	}

	if tree == nil {
		t.Errorf("Error: could not find key '%s'", key)
	}
}
