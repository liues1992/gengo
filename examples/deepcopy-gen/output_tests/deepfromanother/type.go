package deepfromanother

// +k8s:deepcopy-gen=otherpackage.A
type A struct {
	F1 string
	F2 *B
}

type B struct {
	F1 string
	F2 string
}
