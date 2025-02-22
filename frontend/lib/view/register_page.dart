import 'package:flutter/material.dart';
import 'package:frontend/controller/register_controller.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class RegisterPage extends StatefulWidget {
  const RegisterPage({super.key});

  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  bool _isObscure = true;
  final RegisterController controller = Get.put(RegisterController());

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 50),
          child: Column(
            children: [
              Container(
                width: 174,
                height: 144,
                margin: const EdgeInsets.only(top: 100),
                decoration: const BoxDecoration(
                  image: DecorationImage(
                    image: AssetImage("assets/logo-purple.png"),
                  ),
                ),
              ),
              Text(
                "Warasin",
                style: GoogleFonts.poppins(
                  fontSize: 30,
                  fontWeight: bold,
                  color: primaryColor,
                ),
              ),

              const SizedBox(height: 30),

              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    "Selamat Datang di Warasin!",
                    style: GoogleFonts.poppins(
                      fontSize: 20,
                      fontWeight: semiBold,
                      color: primaryTextColor,
                    ),
                    textAlign: TextAlign.center,
                  ),
                  const SizedBox(height: 8),
                  Text(
                    "Jaga kesehatan mentalmu bersama Warasin",
                    style: GoogleFonts.poppins(
                      fontSize: 16,
                      fontWeight: regular,
                      color: tertiaryTextColor,
                    ),
                    textAlign: TextAlign.center,
                  ),
                ],
              ),

              const SizedBox(height: 30),

              Column(
                children: [
                  TextFormField(
                    controller: controller.nameController,
                    decoration: InputDecoration(
                      hintText: "Nama",
                      filled: true,
                      fillColor: secondaryTextColor,
                      contentPadding: const EdgeInsets.symmetric(
                        horizontal: 25,
                        vertical: 20,
                      ),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: BorderSide.none,
                      ),
                    ),
                  ),

                  const SizedBox(height: 5),

                  Obx(
                    () =>
                        controller.isName.value
                            ? const SizedBox.shrink()
                            : Text(
                              'Nama tidak boleh kosong',
                              style: GoogleFonts.poppins(
                                color: Colors.red,
                                fontSize: 12,
                              ),
                            ),
                  ),

                  const SizedBox(height: 10),

                  // Input Email
                  TextFormField(
                    controller: controller.emailController,
                    decoration: InputDecoration(
                      hintText: "Email",
                      filled: true,
                      fillColor: secondaryTextColor,
                      contentPadding: const EdgeInsets.symmetric(
                        horizontal: 25,
                        vertical: 20,
                      ),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: BorderSide.none,
                      ),
                    ),
                  ),

                  const SizedBox(height: 5),

                  Obx(
                    () =>
                        controller.isEmail.value
                            ? const SizedBox.shrink()
                            : Text(
                              'Email tidak valid',
                              style: GoogleFonts.poppins(
                                color: Colors.red,
                                fontSize: 12,
                              ),
                            ),
                  ),

                  const SizedBox(height: 10),

                  TextFormField(
                    controller: controller.passwordController,
                    obscureText: _isObscure,
                    decoration: InputDecoration(
                      hintText: "Password",
                      filled: true,
                      fillColor: secondaryTextColor,
                      contentPadding: const EdgeInsets.symmetric(
                        horizontal: 25,
                        vertical: 20,
                      ),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: BorderSide.none,
                      ),
                      suffixIcon: IconButton(
                        onPressed: () {
                          setState(() {
                            _isObscure = !_isObscure;
                          });
                        },
                        icon: Icon(
                          _isObscure ? Icons.visibility_off : Icons.visibility,
                        ),
                      ),
                    ),
                  ),

                  const SizedBox(height: 5),

                  Obx(
                    () =>
                        controller.isPassword.value
                            ? const SizedBox.shrink()
                            : Text(
                              'Password tidak boleh kosong',
                              style: GoogleFonts.poppins(
                                color: Colors.red,
                                fontSize: 12,
                              ),
                            ),
                  ),

                  const SizedBox(height: 20),

                  Obx(
                    () => SizedBox(
                      width: double.infinity,
                      child: ElevatedButton(
                        onPressed:
                            controller.isLoading.value
                                ? null
                                : () {
                                  controller.register(context);
                                },
                        style: ElevatedButton.styleFrom(
                          backgroundColor: primaryColor,
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(20),
                          ),
                          padding: const EdgeInsets.symmetric(vertical: 16),
                        ),
                        child:
                            controller.isLoading.value
                                ? const CircularProgressIndicator(
                                  color: Colors.white,
                                )
                                : Text(
                                  "Register",
                                  style: GoogleFonts.poppins(
                                    fontSize: 18,
                                    fontWeight: semiBold,
                                    color: secondaryTextColor,
                                  ),
                                ),
                      ),
                    ),
                  ),
                ],
              ),

              const SizedBox(height: 20),

              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    "Sudah punya akun?",
                    style: GoogleFonts.poppins(
                      fontSize: 14,
                      fontWeight: medium,
                      color: primaryTextColor,
                    ),
                  ),
                  TextButton(
                    onPressed: () {
                      Navigator.pushNamed(context, '/login');
                    },
                    child: Text(
                      "Login!",
                      style: GoogleFonts.poppins(
                        fontSize: 14,
                        fontWeight: medium,
                        color: primaryColor,
                      ),
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
