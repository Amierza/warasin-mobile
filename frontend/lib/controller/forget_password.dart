import 'package:flutter/material.dart';
import 'package:get/get.dart';

class ForgetPassword extends GetxController{
  var isEmail = true.obs;
  var isLoading = false.obs;

  final emailController = TextEditingController();

  void forgetPasswordEmail (context) async {
    if(emailController.text.isEmpty || !emailController.text.isEmail){
      isEmail.value = false;
    } else {
      isEmail.value = true;
    }
  }
} 