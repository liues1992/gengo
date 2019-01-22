package deepfrom

import another "github.com/liues1992/gengo/examples/deepcopy-gen/output_tests/deepfromanother"

func (out *A) DeepCopyFrom(in *another.A) {
	out.F1 = in.F1
	out.F2 = &B{}
	out.F2.DeepCopyFrom(in.F2)
}

func (out *B) DeepCopyFrom(in *another.B) {
	out.F1 = in.F1
	out.F2 = in.F2
}
