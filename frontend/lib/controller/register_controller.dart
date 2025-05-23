import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend/service/api_service.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/dialog_auth.dart';
import 'package:get/get.dart';

class RegisterController extends GetxController {
  var isName = true.obs;
  var isEmail = true.obs;
  var isPassword = true.obs;
  var isLoading = false.obs;

  final nameController = TextEditingController();
  final emailController = TextEditingController();
  final passwordController = TextEditingController();

  void register(context) async {
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
      return;
    }

    isLoading.value = true;

    try {
      final response = await ApiService.register(
        nameController.text,
        emailController.text,
        passwordController.text,
      );

      if (response.statusCode == 200) {
        showCustomDialog(
          context: context, 
          icon: Icons.check_circle, 
          title: 'Registrasi Berhasil', 
          message: "Semoga harimu menyenangkan", 
          onPressed: () {
            Navigator.pushNamed(context, '/login');
          },
          backgroundColor: successColor
        );
      } else {
        final Map<String, dynamic> responseData = jsonDecode(response.body);
        final errorMessage = responseData['error'] ?? "Terjadi kesalahan";

        showCustomDialog(
          context: context, 
          icon: Icons.cancel, 
          title: 'Registrasi Gagal', 
          message: errorMessage,
          onPressed: () {
            Navigator.of(context).pop();
          },
          backgroundColor: dangerColor
        );
      }
    } catch (err) {
      print('TErjadi ERror $err');
      showCustomDialog(
        context: context, 
        icon: Icons.cancel, 
        title: 'Registrasi Gagal', 
        message: "Oops..! Coba cek lagi", 
        onPressed: () {
          Navigator.of(context).pop();
        },
        backgroundColor: dangerColor
      );
    } finally {
      isLoading.value = false;
    }
  }
}
