package main

import "testing"

func TestAddDependant(t *testing.T) {
	pkg1 := pkg{
		Name: "pkg1",
	}
	pkg2 := pkg{
		Name: "pkg2",
	}
	pkg1.addDependant(&pkg2)
	if len(pkg1.DependantPkgs) != 1 {
		t.Fatal("failed to add package")
	}
	if pkg1.DependantPkgs[0] != &pkg2 {
		t.Fatal("added the wrong package")
	}

	// ensure it doesnt get added twice
	pkg1.addDependant(&pkg2)
	if len(pkg1.DependantPkgs) != 1 {
		t.Fatal("duplicate package added")
	}
	// ensure it doesnt get added twice
	pkg1.addDependant(&pkg1)
	if len(pkg1.DependantPkgs) != 1 {
		t.Fatal("added self")
	}
}

func TestRemoveDependant(t *testing.T) {
	pkg1 := pkg{
		Name: "pkg1",
	}
	pkg2 := pkg{
		Name: "pkg2",
	}
	pkg1.addDependant(&pkg2)
	pkg1.removeDependant(&pkg1)

	if len(pkg1.DependantPkgs) != 1 {
		t.Fatal("pkg count mismatch")
	}
	if pkg1.DependantPkgs[0] != &pkg2 {
		t.Fatal("removed the wrong package")
	}

	pkg1.removeDependant(&pkg2)
	if len(pkg1.DependantPkgs) != 0 {
		t.Fatal("pkg count mismatch")
	}

}