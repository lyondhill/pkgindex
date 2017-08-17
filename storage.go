package main

import (
	"errors"
	"sync"
)

var (
	// locking device ensure success on write
	locker = sync.RWMutex{}

	// all of the packages that are in the index
	indexedPkgs = map[string]*pkg{}

	// fail to index because its dependencies arent indexed
	missingPackage      = errors.New("missing package")
	missingDependencies = errors.New("missing dependencies")
	isDependant         = errors.New("other package depend on me")
)

// add a pkg to the index
func index(pkg *pkg) error {

	// ensure dependencies are indexed
	for _, dependency := range pkg.Dependencies {
		if _, ok := indexRead(dependency); !ok {
			return missingDependencies
		}
	}

	// if this index already exists we need to just update dependencies
	existingPkg, ok := indexRead(pkg.Name)
	if ok {
		return updateDependents(existingPkg, pkg)
	}

	// update any dependants of this package
	updateDependents(nil, pkg)

	// add the new index (possibly replacing the old)
	indexWrite(pkg.Name, pkg)

	return nil
}

// remove a pkg from the index if it doesnt have any dependencies
func remove(pkg *pkg) error {
	// ensure no package are depending on this pkg
	existingPkg, ok := indexRead(pkg.Name)
	if ok {
		if len(existingPkg.DependantPkgs) != 0 {
			return isDependant
		}

		updateDependents(existingPkg, nil)
	}

	locker.Lock()
	defer locker.Unlock()

	// remove the index
	delete(indexedPkgs, pkg.Name)

	return nil
}

// check to see if a pkg is in the index
func query(pkg *pkg) error {
	// check the indexed packages to see if this new pkg exists
	pkg, ok := indexRead(pkg.Name)

	if !ok {
		return missingPackage
	}
	return nil
}

// just update the pkgs that this one depends on
func updateDependents(existingPkg, newPkg *pkg) error {

	// if there is no existing pkg we will just add newpkg to the dependant list
	if existingPkg == nil {
		for _, pkg := range newPkg.Dependencies {
			newDependant, ok := indexRead(pkg)
			if ok {
				newDependant.addDependant(newPkg)
			}
		}
		return nil
	}

	// if there is no new package we can just remove ourself from the existing dependencies
	if newPkg == nil {
		for _, pkg := range existingPkg.Dependencies {
			oldDependant, ok := indexRead(pkg)
			if ok {
				oldDependant.removeDependant(existingPkg)
			}
		}
		return nil
	}

	// we dont want to do this concurrently
	existingPkg.Lock()
	defer existingPkg.Unlock()

	// remove myself as a dependant from any package new doesnt rely on that existing did
	for _, pkg := range minus(existingPkg.Dependencies, newPkg.Dependencies) {
		oldDependant, ok := indexRead(pkg)
		if ok {
			oldDependant.removeDependant(existingPkg)
		}
	}

	// add the existing package to any dependant that is new
	for _, pkg := range minus(newPkg.Dependencies, existingPkg.Dependencies) {
		newDependant, ok := indexRead(pkg)
		if ok {
			newDependant.addDependant(existingPkg)
		}
	}

	// set existing packages dependencies to the new one
	existingPkg.Dependencies = newPkg.Dependencies

	return nil
}

func indexRead(key string) (*pkg, bool) {
	locker.RLock()
	defer locker.RUnlock()

	p, ok := indexedPkgs[key]
	return p, ok
}

func indexWrite(key string, p *pkg) {
	locker.Lock()
	defer locker.Unlock()

	indexedPkgs[key] = p
}

func minus(slice1, slice2 []string) []string {
	rtn := []string{}
	for _, elem1 := range slice1 {
		found := false
		for _, elem2 := range slice2 {
			if elem1 == elem2 {
				found = true
				break
			}
		}
		if !found {
			rtn = append(rtn, elem1)
		}
	}
	return rtn
}
