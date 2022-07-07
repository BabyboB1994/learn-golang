package main

var a = b + c

var b = 2
var c = a // this is to check how golang handles the cyclic dependecy between the variable initialization

// initialization cycle for a (see details)compilerInvalidInitCycle
// PackageInitialization.go(3, 5): initialization cycle for a
// PackageInitialization.go(3, 5): a refers to
// PackageInitialization.go(5, 5): c refers to (this error)
// PackageInitialization.go(3, 5): a
