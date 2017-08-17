package main

import "testing"
import "fmt"

func TestIndex(t *testing.T) {

	pkg1 := &pkg{Name: "pkg1", Dependencies: []string{"fake"}}
	if err := index(pkg1); err != missingDependencies {
		t.Fatal("added a package even without dependencies")
	}
	pkg1.Dependencies = []string{}
	if err := index(pkg1); err != nil {
		t.Fatal("failed to index %s", err)
	}

	pkg2 := &pkg{Name: "pkg2", Dependencies: []string{"pkg1"}}
	if err := index(pkg2); err != nil {
		t.Fatal("failed to index %s", err)
	}

	if len(pkg1.DependantPkgs) != 1 {
		t.Fatal("failed to add dependent package to pkg1")
	}

	// create a third package that depends on both previous package
	pkg3 := &pkg{Name: "pkg3", Dependencies: []string{"pkg1", "pkg2"}}
	if err := index(pkg3); err != nil {
		t.Fatal("failed to index %s", err)
	}

	if len(pkg1.DependantPkgs) != 2 {
		t.Fatal("failed to add second dependent package to pkg1")
	}

	if len(pkg2.DependantPkgs) != 1 {
		t.Fatal("failed to add dependent package to pkg2")
	}

	// update pkg3 to only depend on pkg2
	pkg3 = &pkg{Name: "pkg3", Dependencies: []string{"pkg2"}}
	if err := index(pkg3); err != nil {
		t.Fatal("failed to index %s", err)
	}

	if len(pkg1.DependantPkgs) != 1 {
		t.Fatal("failed to add second dependent package to pkg1")
	}

	if len(pkg2.DependantPkgs) != 1 {
		t.Fatal("failed to add dependent package to pkg2")
	}

	// update pkg3 to only depend on pkg1
	pkg3 = &pkg{Name: "pkg3", Dependencies: []string{"pkg1"}}
	if err := index(pkg3); err != nil {
		t.Fatal("failed to index %s", err)
	}

	if len(pkg1.DependantPkgs) != 2 {
		t.Fatal("failed to add second dependent package to pkg1")
	}

	if len(pkg2.DependantPkgs) != 0 {
		t.Fatal("failed to add dependent package to pkg2")
	}

}

func TestRemove(t *testing.T) {
	pkg1 := &pkg{Name: "pkg1"}
	if err := remove(pkg1); err != isDependant {
		t.Fatal("removed a package that was dependent")
	}

	pkg2 := &pkg{Name: "pkg2"}
	if err := remove(pkg2); err != nil {
		t.Fatal("failed to remove independant pkg2")
	}

	pkg3 := &pkg{Name: "pkg3"}
	if err := remove(pkg3); err != nil {
		t.Fatal("failed to remove independant pkg3")
	}

	pkg1 = &pkg{Name: "pkg1"}
	if err := remove(pkg1); err != nil {
		t.Fatal("failed to remove independant pkg1")
	}
}

func TestQuery(t *testing.T) {
	pkg1 := &pkg{Name: "pkg1"}
	if err := query(pkg1); err != missingPackage {
		t.Fatal("found a package that shouldnt exist")
	}
	index(pkg1)
	if err := query(pkg1); err != nil {
		t.Fatal("index added but failed to query")
	}
}

func TestMinus(t *testing.T) {
	if fmt.Sprintf("%v", minus([]string{"hi", "dude", "or", "old", "friend"}, []string{"dude", "or"})) != "[hi old friend]" {
		t.Fatal("failed to subtract elements from both arrays")
	}
	if fmt.Sprintf("%v", minus([]string{"one"}, []string{"two"})) != "[one]" {

	}
}
