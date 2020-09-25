package android
type Product_variables struct {
	Target_init_vendor_lib struct {
		Whole_static_libs []string
	}
	Target_process_sdk_version_override struct {
		Cppflags []string
	}
	Target_shim_libs struct {
		Cppflags []string
	}
		Target_surfaceflinger_fod_lib struct {
		Cflags []string
		Whole_static_libs []string
	}
	Target_uses_color_metadata struct {
		Cppflags []string
	}
	Additional_gralloc_10_usage_bits struct {
		Cppflags []string
	}
		Bootloader_message_offset struct {
		Cflags []string
	}
		Qcom_um_soong_namespace struct {
		Cflags []string
		Header_libs []string
	}
		Uses_qcom_um_family struct {
		Cflags []string
		Shared_libs []string
	}
	Should_wait_for_qsee struct {
		Cflags []string
	}
	Uses_qcom_um_3_18_family struct {
		Header_libs []string
		Shared_libs []string
	}
	Uses_qcom_um_4_4_family struct {
		Header_libs []string
		Shared_libs []string
	}
	Uses_qcom_um_4_9_family struct {
		Header_libs []string
		Shared_libs []string
	}
	Uses_qcom_um_4_14_family struct {
		Header_libs []string
		Shared_libs []string
	}
	Supports_extended_compress_format struct {
		Cflags []string
	}
}

type ProductVariables struct {
	Target_init_vendor_lib				*string	`json:",omitempty"`
	Target_process_sdk_version_override	*string	`json:",omitempty"`
	Target_shim_libs					*string	`json:",omitempty"`
	Target_surfaceflinger_fod_lib		*string	`json:",omitempty"`
	Target_uses_color_metadata			*bool	`json:",omitempty"`
	Additional_gralloc_10_usage_bits	*string	`json:",omitempty"`
	Bootloader_message_offset			*int	`json:",omitempty"`
	Should_wait_for_qsee				*bool	`json:",omitempty"`
	Qcom_um_soong_namespace				*string	`json:",omitempty"`
	Uses_qcom_um_family					*bool	`json:",omitempty"`
	Uses_qcom_um_3_18_family			*bool	`json:",omitempty"`
	Uses_qcom_um_4_4_family				*bool	`json:",omitempty"`
	Uses_qcom_um_4_9_family				*bool	`json:",omitempty"`
	Uses_qcom_um_4_14_family 			*bool	`json:",omitempty"`
	Supports_extended_compress_format	*bool	`json:",omitempty"`
}