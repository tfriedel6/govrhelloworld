
LOCAL_PATH := $(call my-dir)

#--------------------------------------------------------
# libgovrlib.so
#--------------------------------------------------------
include $(CLEAR_VARS)

LOCAL_MODULE			:= govrlib
LOCAL_MODULE_FILENAME   := libgovrlib
LOCAL_SRC_FILES			:= $(LOCAL_PATH)/../../../govrlib/$(TARGET_ARCH_ABI)/libgovrlib.so
LOCAL_EXPORT_C_INCLUDES := $(LOCAL_PATH)/../../../govrlib

include $(PREBUILT_SHARED_LIBRARY)

#--------------------------------------------------------
# govrhelloworld.so
#--------------------------------------------------------
include $(CLEAR_VARS)

LOCAL_MODULE			:= govrhelloworld
LOCAL_CFLAGS			:= -std=c99 -Werror
LOCAL_SRC_FILES			:= ../../../Src/main.c
LOCAL_LDLIBS			:= -llog -landroid -lGLESv3 -lEGL
LOCAL_LDFLAGS			:= -u ANativeActivity_onCreate
LOCAL_STATIC_LIBRARIES	:= android_native_app_glue 
LOCAL_SHARED_LIBRARIES	:= vrapi govrlib

include $(BUILD_SHARED_LIBRARY)

$(call import-add-path, $(OCULUS_SDK_PATH)/)
$(call import-module,android/native_app_glue)
$(call import-module,VrApi/Projects/AndroidPrebuilt/jni)
