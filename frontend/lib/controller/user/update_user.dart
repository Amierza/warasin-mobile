import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/model/user_detail_response.dart';
import 'package:frontend/service/api_service.dart';
import 'package:get/get.dart';
import 'package:image_picker/image_picker.dart';

class UpdateUser extends GetxController {
  final isLoading = false.obs;
  final nameController = TextEditingController();
  final emailController = TextEditingController();
  final phoneController = TextEditingController();
  final birthDateController = TextEditingController();
  final provinceController = TextEditingController();
  final cityController = TextEditingController();
  final name = ''.obs;
  final email = ''.obs;
  final phoneNumber = ''.obs;
  final birthDate = ''.obs;
  final cityId = ''.obs;
  final gender = Rxn<bool>();
  final province = ''.obs;

  final imageFile = Rx<File?>(null);
  final imageBase64 = ''.obs;

  @override
  void onInit() {
    fetchUserDetail();
    super.onInit();
  }

  Future<void> pickImage() async {
    try {
      final picked = await ImagePicker().pickImage(
        source: ImageSource.gallery,
        maxWidth: 800,
        maxHeight: 800,
        imageQuality: 85,
      );

      if (picked != null) {
        imageFile.value = File(picked.path);
        final bytes = await picked.readAsBytes();
        imageBase64.value = "data:image/png;base64,${base64Encode(bytes)}";
      }
    } catch (e) {
      Get.snackbar(
        "Error",
        "Gagal memilih gambar: ${e.toString()}",
        snackPosition: SnackPosition.BOTTOM,
      );
    }
  }

  Future<void> updateProfile() async {
    try {
      if (!_validateFields()) return;

      name.value = nameController.text;
      email.value = emailController.text;
      phoneNumber.value = phoneController.text;
      birthDate.value = birthDateController.text;
      isLoading.value = true;

      final response = await ApiService.updateUserService(
        name.value,
        email.value,
        imageBase64.value,
        gender.value,
        birthDate.value,
        phoneNumber.value,
        cityId.value,
      );

      if (response is UserDetailResponse) {
        print(response.data.userBirthDate);
        Get.snackbar("Sukses", "Profil berhasil diperbarui");
        await fetchUserDetail();
      } else if (response is ErrorResponse) {
        Get.snackbar("Error", response.message);
      }
    } catch (e) {
      _handleError(e);
    } finally {
      isLoading.value = false;
    }
  }

  bool _validateFields() {
    if (name.value.isEmpty ||
        email.value.isEmpty ||
        gender.value == null ||
        birthDate.value.isEmpty ||
        phoneNumber.value.isEmpty ||
        cityId.value.isEmpty) {
      Get.snackbar(
        "Error",
        "Semua field wajib diisi",
        snackPosition: SnackPosition.BOTTOM,
      );
      return false;
    }
    return true;
  }

  void _handleError(dynamic e) {
    print("Error updateProfile: ${e.toString()}");
    print("Stack trace: ${StackTrace.current}");

    Get.snackbar(
      "Error",
      "Gagal update profil: ${e.toString()}",
      snackPosition: SnackPosition.BOTTOM,
      duration: const Duration(seconds: 5),
    );
  }

  Future<void> fetchUserDetail() async {
    try {
      isLoading.value = true;
      final result = await ApiService.getUserDetailService();

      if (result is UserDetailResponse) {
        _updateUserData(result);
      }
    } finally {
      isLoading.value = false;
    }
  }

  void _updateUserData(UserDetailResponse result) {
    // Update nilai controller langsung
    nameController.text = result.data.userName;
    emailController.text = result.data.userEmail;
    phoneController.text = result.data.userPhoneNumber;
    birthDateController.text = result.data.userBirthDate;

    // Update reactive values
    name.value = result.data.userName;
    email.value = result.data.userEmail;
    phoneNumber.value = result.data.userPhoneNumber;
    birthDate.value = result.data.userBirthDate;
    cityId.value = result.data.city.cityId.toString();
    province.value = result.data.city.province?.provinceId ?? '';
    gender.value = result.data.userGender;
  }
}
