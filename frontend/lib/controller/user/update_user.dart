import 'dart:io';

import 'package:flutter/material.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/model/user_detail_response.dart';
import 'package:frontend/service/api_service.dart';
import 'package:get/get.dart';
import 'package:image_picker/image_picker.dart';

class UpdateUser extends GetxController {
  final isLoading = false.obs;
  final user = Rxn<User>();
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
  var isInitialized = false.obs;

  File? selectedPhoto;

  final imageFile = Rx<File?>(null);
  final imageBase64 = ''.obs;

  @override
  void onInit() {
    fetchUserDetail();
    super.onInit();
  }

  Future<void> pickImage() async {
    try {
      final pickedFile = await ImagePicker().pickImage(
        source: ImageSource.gallery,
      );

      if (pickedFile != null) {
        selectedPhoto = File(pickedFile.path);
        update();
      }
    } catch (e) {
      Get.snackbar(
        "Error",
        "Gagal memilih gambar: ${e.toString()}",
        snackPosition: SnackPosition.BOTTOM,
      );
    }
  }

  void updateProfile() async {
    isLoading.value = true;

    try {
      if (!_validateFields()) return;

      final Map<String, dynamic> updatedFields = {};

      if (nameController.text != user.value?.userName) {
        updatedFields['name'] = nameController.text;
      }

      if (emailController.text != user.value?.userEmail) {
        updatedFields['email'] = emailController.text;
      }

      if (phoneController.text != user.value?.userPhoneNumber) {
        updatedFields['phone_number'] = phoneController.text;
      }

      if (birthDateController.text != user.value?.userBirthDate) {
        updatedFields['birth_date'] = birthDateController.text;
      }

      if (cityId.value != user.value?.city.cityId) {
        updatedFields['city_id'] = cityId.value;
      }

      // if (selectedPhoto == user.value?.userImage) {
      //   selectedPhoto = user.value?.userImage;
      // }

      isLoading.value = true;

      final response = await ApiService.updateUserService(
        updatedFields: updatedFields,
        imageFile: selectedPhoto,
      );

      if (response is UserDetailResponse) {
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
        user.value = result.data;
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
     cityController.text = result.data.city.cityId.toString(); 

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
