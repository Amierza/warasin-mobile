import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend/service/api_service.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/dialog_auth.dart';
import 'package:get/get.dart';

class ForgetPassword extends GetxController {
  var isEmail = true.obs;
  var emailExist = false.obs;
  var isLoading = false.obs;

  final emailController = TextEditingController();

  void forgetPasswordEmail(context) async {
    if (emailController.text.isEmpty || !emailController.text.isEmail) {
      isEmail.value = false;
    } else {
      isEmail.value = true;
    }

    if (!isEmail.value) {
      return;
    }

    isLoading.value = true;

    try {
      final response = await ApiService.checkEmail(emailController.text);

      if (response.statusCode == 200) {
        final Map<String, dynamic> responseData = jsonDecode(response.body);
        emailExist.value = responseData['status'];

        if (!emailExist.value) {
          showCustomDialog(
            context: context,
            icon: Icons.cancel,
            title: 'Reset Password Gagal',
            message: 'Email belum terdaftar',
            onPressed: () {
              Navigator.of(context).pop();
            },
            backgroundColor: dangerColor,
          );
        }

        showCustomDialog(
          context: context,
          icon: Icons.check_circle,
          title: 'Reset Password Berhasil',
          message: 'Silahkan cek email anda',
          onPressed: () {
            Navigator.pushNamedAndRemoveUntil(
              context,
              '/login',
              (route) => false,
            );
          },
          backgroundColor: successColor,
        );
      }
    } catch (err) {
      showCustomDialog(
        context: context,
        icon: Icons.cancel,
        title: 'Reset Password Gagal',
        message: 'Terjadi kesalahan',
        onPressed: () {
          Navigator.of(context).pop();
        },
        backgroundColor: dangerColor,
      );
    } finally {
      isLoading.value = false;
    }
  }
}
