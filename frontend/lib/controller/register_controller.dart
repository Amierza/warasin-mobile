import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend/service/api_service.dart';
import 'package:get/get.dart';

class RegisterController extends GetxController {
  var isName = true.obs;
  var isEmail = true.obs;
  var isPassword = true.obs;
  var isLoading = false.obs;

  final nameController = TextEditingController();
  final emailController = TextEditingController();
  final passwordController = TextEditingController();

  void register() async {
    if (emailController.text.isEmpty || !emailController.text.isEmail) {
      isEmail.value = false;
    } else {
      isEmail.value = true;
    }

    if (nameController.text.isEmpty) {
      isName.value = false;
    } else {
      isName.value = true;
    }

    if (passwordController.text.isEmpty) {
      isPassword.value = false;
    } else {
      isPassword.value = true;
    }

    if (!isEmail.value || !isName.value || !isPassword.value) {
      return ;
    }

    isLoading.value = true;

    try {
      final response = await ApiService.register(
        nameController.text,
        emailController.text,
        passwordController.text,
      );

      if (response.statusCode == 200) {
        Get.offNamed('/login');
        Get.snackbar('Success', 'Registrasi berhasil!');
      } else {
        final responseData = jsonDecode(response.body);
        Get.snackbar('Error', responseData['message'] ?? 'Registrasi gagal');
      }
    } catch (err) {
      if (Get.context != null) {
        Get.snackbar('Error', 'Terjadi kesalahan: $err');
      }
    } finally {
      isLoading.value = false;
    }
  }
}
