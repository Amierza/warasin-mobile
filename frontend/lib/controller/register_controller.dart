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
      print("Validasi gagal");
      return;
    }

    isLoading.value = true;
    print("Mengirim request ke server...");

    try {
      final response = await ApiService.register(
        nameController.text,
        emailController.text,
        passwordController.text,
      );

      print("Response Status Code: ${response.statusCode}");
      print("Response Body: ${response.body}");

      if (response.statusCode == 200) {
        showCustomDialog(
          context: context, 
          icon: Icons.check_circle, 
          title: 'Registrasi Berhasil', 
          message: "Semoga harimu menyenangkan", 
          backgroundColor: successColor
        );
      } else {
        showCustomDialog(
          context: context, 
          icon: Icons.cancel, 
          title: 'Registrasi Gagal', 
          message: "Oops..! Coba cek lagi", 
          backgroundColor: dangerColor
        );
      }
    } catch (err) {
        showCustomDialog(
          context: context, 
          icon: Icons.cancel, 
          title: 'Registrasi Gagal', 
          message: "Oops..! Coba cek lagi", 
          backgroundColor: dangerColor
        );
    } finally {
      isLoading.value = false;
    }
  }
}
