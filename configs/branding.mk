#
# Copyright (C) 2020 The Project Adara
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Set all versions
ADARA_BUILD_TYPE ?= Unofficial
ADARA_CODENAME := Cupcake
ADARA_ANDROID_NAME := R

# Set all versions
BUILD_DATE := $(shell date -u +%d%m%Y)
BUILD_TIME := $(shell date -u +%H%M)
BUILD_DATE_TIME := $(BUILD_TIME)$(BUILD_DATE)
ROM_FINGERPRINT := Adara/$(ADARA_BUILD_TYPE)/$(PLATFORM_VERSION)/$(BUILD_ID)/$(BUILD_DATE)/$(BUILD_TIME)

ADARA_VERSION := AdaraOS-$(ADARA_ANDROID_NAME)-$(ADARA_CODENAME)-$(ADARA_BUILD)-$(BUILD_DATE)-$(ADARA_BUILD_TYPE)

ADARA_PROPERTIES := \
    BUILD_DISPLAY_ID=$(BUILD_ID) \
    ro.build.datetime=$(BUILD_DATE_TIME) \
    ro.adara.build.date=$(BUILD_DATE) \
    ro.adara.version=$(ADARA_VERSION) \
    ro.adara.codename=$(ADARA_CODENAME) \
    ro.adara.build.type=$(ADARA_BUILD_TYPE) \
    ro.adara.build.version=$(ADARA_VERSION_BASE) \
    ro.adara.fingerprint=$(ROM_FINGERPRINT)