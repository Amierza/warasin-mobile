import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend/service/api_service.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/dialog_auth.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

class LoginController extends GetxController{
  var isEmail = true.obs;
  var isPassword = true.obs;
  var isLoading = false.obs;

  final emailController = TextEditingController();
  final passwordController = TextEditingController();

  void login(context) async {
    if (emailController.text.isEmpty || !emailController.text.isEmail) {
      isEmail.value = false;
    } else {
      isEmail.value = true;
    }

    if (passwordController.text.isEmpty) {
      isPassword.value = false;
    } else {
      isPassword.value = true;
    }

    if (!isEmail.value || !isPassword.value) {
      return;
    }

    isLoading.value = true;

    try {
      final response = await ApiService.login(
        emailController.text, 
        passwordController.text,
      );

      if (response.statusCode == 200) {
        final Map<String, dynamic> responseData = jsonDecode(response.body);

        if (responseData.containsKey('data')) {
          final data = responseData['data'];
          if (data != null && data.containsKey('access_token')) {
            final accessToken = data['access_token'];

            showCustomDialog(
              context: context,
              icon: Icons.check_circle,
              title: 'Login Berhasil',
              message: 'Semoga harimu menyenangkan',
              onPressed: () {
                Navigator.pushNamedAndRemoveUntil(
                  context, 
                  '/home',
                  (route) => false
                );
              },
              backgroundColor: successColor,
            );
            final prefs = await SharedPreferences.getInstance();
            await prefs.setString('access_token', accessToken);
          } else {
            showCustomDialog(
              context: context,
              icon: Icons.cancel,
              title: 'Login Gagal',
              message: 'Access token tidak ditemukan',
              onPressed: () {
                Navigator.of(context).pop();
              },
              backgroundColor: dangerColor,
            );
          }
        } else {
          showCustomDialog(
            context: context,
            icon: Icons.cancel,
            title: 'Login Gagal',
            message: 'Data tidak valid',
            onPressed: () {
              Navigator.of(context).pop();
            },
            backgroundColor: dangerColor,
          );
        }
      } else {
        final Map<String, dynamic> responseData = jsonDecode(response.body);
        final errorMessage = responseData['error'] ?? "Terjadi kesalahan";
        showCustomDialog(
          context: context,
          icon: Icons.cancel,
          title: 'Login Gagal',
          message: errorMessage,
          onPressed: () {
            Navigator.of(context).pop();
          },
          backgroundColor: dangerColor,
        );
      }
    } catch (err) {
      showCustomDialog(
        context: context, 
        icon: Icons.cancel, 
        title: 'Login Gagal', 
        message: "Oops..! Terjadi kesalahan: $err", 
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