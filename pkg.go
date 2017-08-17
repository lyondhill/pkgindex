package main

import "sync"

type pkg struct {
	Command string
	Name string
	Dependencies []string
	DependantPkgs []*pkg

	// add the ability to lock and unlock each package
	sync.Mutex
}

// add a package to the list of packages that depend on me
func (self *pkg) addDependant(pkg *pkg) {
	// no pkg can depend on itself
	if self == pkg {
		return
	}

	// lock before we do anything just incase someone else is trying to remove one
	self.Lock()
	defer self.Unlock()


	// if this package is aleady in the list we shouldnt add it twice
	for i := 0; i < len(self.DependantPkgs); i++ {
		if self.DependantPkgs[i] == pkg {
			return
		}
	}

	self.DependantPkgs = append(self.DependantPkgs, pkg)
}

// remove a package from the list of package that depend on me
func (self *pkg) removeDependant(pkg *pkg) {
	self.Lock()
	defer self.Unlock()

	for i := 0; i < len(self.DependantPkgs); i++ {
		if pkg == self.DependantPkgs[i] {
			self.DependantPkgs = append(self.DependantPkgs[:i], self.DependantPkgs[i+1:]...)
			return
		}
	}
}