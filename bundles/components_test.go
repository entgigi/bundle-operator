package bundles

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestUnmarshalling(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)

	yfile, err := ioutil.ReadFile("../config/bundle-spec/spec/descriptor.yaml")
	if err != nil {
		t.Fatal(err.Error())
	}
	data := &BundleDescriptor{}
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedName := "example"
	actualName := data.Name
	if actualName != expectedName {
		t.Fatalf("Invalid Domain for %q. Expected %q, got %q", data, expectedName, actualName)
	}

	expectedVersion := "v1.0.0"
	actualVersion := data.Version
	if actualVersion != expectedVersion {
		t.Fatalf("Invalid Domain for %q. Expected %q, got %q", data, expectedVersion, actualVersion)
	}

	expectedComponentsNumber := 2
	actualComponentsNumber := len(data.Components)
	if actualComponentsNumber != expectedComponentsNumber {
		t.Fatalf("Invalid Domain for %q. Expected %q, got %q", data, expectedComponentsNumber, actualComponentsNumber)
	}

	plugin, actualTypeIsPlugin := data.Components[0].Spec.(*Plugin)
	fmt.Println(reflect.TypeOf(data.Components[0].Spec))
	if !actualTypeIsPlugin {
		t.Fatalf("Invalid Domain for %q. Actual type is plugin %t, got %q", data, actualTypeIsPlugin, plugin)
	}

	mytype2, actualTypeIsPlugin2 := data.Components[1].Spec.(*Manifest)
	fmt.Println(reflect.TypeOf(data.Components[1].Spec))
	if !actualTypeIsPlugin2 {
		t.Fatalf("Invalid Domain for %q. Actual type is plugin %t, got %q", data, actualTypeIsPlugin2, mytype2)
	}

	expectedRepository := "nginx"
	actualRepository := plugin.Repository
	if actualRepository != expectedRepository {
		t.Fatalf("Invalid Domain for %q. Expected %q, got %q", data, expectedRepository, actualRepository)
	}

}
