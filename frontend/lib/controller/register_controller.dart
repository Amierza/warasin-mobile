import 'package:flutter/widgets.dart';
import 'package:get/get.dart';

class RegisterController extends GetxController{
  var isName = true.obs;
  var isEmail = true.obs;
  var isPassword = true.obs;

  final nameController = TextEditingController();
  final emailController = TextEditingController();
  final passwordController = TextEditingController();

  void register () async {
    if(emailController.text.isEmpty || !emailController.text.isEmail) {
      isEmail.value = false;
    } else {
      isEmail.value = true;
    }

    if(nameController.text.isEmpty) {
      isName.value = false;
    } else {
      isName.value = true;
    }

    if(passwordController.text.isEmpty) {
      isPassword.value = false;
    } else {
      isPassword.value = true;
    }

    if(isName.value && isEmail.value && isPassword.value){
      Get.offNamed('/login');
    }
  }
}
