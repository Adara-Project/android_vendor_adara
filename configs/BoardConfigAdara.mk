# Kernel
include vendor/adara/configs/BoardConfigKernel.mk

# Qcom-specific bits
ifeq ($(BOARD_USES_QCOM_HARDWARE),true)
	include vendor/adara/configs/BoardConfigQcom.mk
endif

# Soong
include vendor/adara/configs/BoardConfigSoong.mk