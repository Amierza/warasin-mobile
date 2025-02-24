import 'package:flutter/material.dart';
import 'package:frontend/controller/forget_password.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/logo_widget.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class ForgetPasswordEmailPage extends StatelessWidget {
  const ForgetPasswordEmailPage({super.key});

  @override
  Widget build(BuildContext context) {
    final ForgetPassword controller = Get.put(ForgetPassword());
    final TextEditingController emailController = TextEditingController();

    return Scaffold(
      backgroundColor: backgroundColor,
      body: SingleChildScrollView(
        child: SizedBox(
          height: MediaQuery.of(context).size.height,
          child: Padding(
            padding: const EdgeInsets.symmetric(
              horizontal: 40,
              vertical: 45
            ),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Column(
                  children: [
                    const LogoWarasin(),
                    const SizedBox(height: 30),
                    Text(
                      "Jika kamu lupa kata sandi, tidak perlu khawatir. Sederhananya, berikan kami alamat email kamu, dan kami akan mengirimkan email berisi tautan untuk mereset kata sandi sehingga kamu dapat membuat yang baru.",
                      style: GoogleFonts.poppins(
                        fontSize: 16,
                        fontWeight: regular,
                        color: tertiaryTextColor,
                      ),
                      textAlign: TextAlign.center,
                    ),
                    const SizedBox(height: 20),
                    TextFormField(
                      controller: emailController,
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
                  ],
                ),

                Row(
                  children: [
                    ElevatedButton(
                      onPressed: () {
                        Navigator.of(context).pop();
                      },
                      style: ElevatedButton.styleFrom(
                        backgroundColor: Colors.transparent,
                        elevation: 0,
                      ),
                      child: Icon(
                        Icons.arrow_back_ios,
                        size: 26,
                        color: Colors.orange,
                      ),
                    ),
                    Expanded(
                      child: ElevatedButton(
                        onPressed: () {
                          controller.forgetPasswordEmail(context);
                          Navigator.pushNamed(context, '/forget_password_password');
                        },
                        style: ElevatedButton.styleFrom(
                          backgroundColor: primaryColor,
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(20),
                          ),
                          padding: const EdgeInsets.symmetric(
                            vertical: 16,
                            horizontal: 12,
                          ),
                        ),
                        child: Text(
                          "Reset Password Link",
                          style: GoogleFonts.poppins(
                            fontSize: 18,
                            fontWeight: semiBold,
                            color: secondaryTextColor,
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ),
        )
      ),
    );
  }
}