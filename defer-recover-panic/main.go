package main

import _case "CloudNative/defer-recover-panic/case"

func main() {
	_case.DeferCase()
	_case.DeferCase1()
	_case.DeferCase2()
	_case.ExceptionCase()
	_case.FileReadCase()
}
