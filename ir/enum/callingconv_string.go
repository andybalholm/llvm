// Code generated by "stringer -linecomment -type CallingConv"; DO NOT EDIT.

package enum

import "strconv"

const (
	_CallingConv_name_0 = "noneccc"
	_CallingConv_name_1 = "fastcccoldccghccccc 11webkit_jsccanyregccpreserve_mostccpreserve_allccswiftcccxx_fast_tlscc"
	_CallingConv_name_2 = "x86_stdcallccx86_fastcallccarm_apcsccarm_aapcsccarm_aapcs_vfpccmsp430_intrccx86_thiscallccptx_kernelptx_device"
	_CallingConv_name_3 = "spir_funcspir_kernelintel_ocl_biccx86_64_sysvccwin64ccx86_vectorcallcchhvmcchhvm_cccx86_intrccavr_intrccavr_signalcccc 86amdgpu_vsamdgpu_gsamdgpu_psamdgpu_csamdgpu_kernelx86_regcallccamdgpu_hscc 94amdgpu_lsamdgpu_es"
)

var (
	_CallingConv_index_0 = [...]uint8{0, 4, 7}
	_CallingConv_index_1 = [...]uint8{0, 6, 12, 17, 22, 33, 41, 56, 70, 77, 91}
	_CallingConv_index_2 = [...]uint8{0, 13, 27, 37, 48, 63, 76, 90, 100, 110}
	_CallingConv_index_3 = [...]uint8{0, 9, 20, 34, 47, 54, 70, 76, 84, 94, 104, 116, 121, 130, 139, 148, 157, 170, 183, 192, 197, 206, 215}
)

func (i CallingConv) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _CallingConv_name_0[_CallingConv_index_0[i]:_CallingConv_index_0[i+1]]
	case 8 <= i && i <= 17:
		i -= 8
		return _CallingConv_name_1[_CallingConv_index_1[i]:_CallingConv_index_1[i+1]]
	case 64 <= i && i <= 72:
		i -= 64
		return _CallingConv_name_2[_CallingConv_index_2[i]:_CallingConv_index_2[i+1]]
	case 75 <= i && i <= 96:
		i -= 75
		return _CallingConv_name_3[_CallingConv_index_3[i]:_CallingConv_index_3[i+1]]
	default:
		return "CallingConv(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}