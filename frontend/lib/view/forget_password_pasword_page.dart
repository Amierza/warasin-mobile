import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/logo_widget.dart';
import 'package:frontend/widget/requiredment_reset_password.dart';
import 'package:google_fonts/google_fonts.dart';

class ForgetPasswordPaswordPage extends StatelessWidget {
  const ForgetPasswordPaswordPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: SingleChildScrollView(
        child: SizedBox(
          height: MediaQuery.of(context).size.height,
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 40, vertical: 45),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Column(
                  children: [
                    const LogoWarasin(),
                    const SizedBox(height: 30),
                    Text(
                      "Masukkan Password Baru",
                      style: GoogleFonts.poppins(
                        fontSize: 16,
                        fontWeight: regular,
                        color: tertiaryTextColor,
                      ),
                    ),
                    const SizedBox(height: 20),
                    TextFormField(
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
                      ),
                    ),
                    const SizedBox(height: 20),
                    TextFormField(
                      decoration: InputDecoration(
                        hintText: "Confirm password",
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
                    const SizedBox(height: 25),
                    RequirementResetPassword(
                      icon: Icons.check_circle,
                      color: successColor,
                      text: "Password minimal 8 karakter",
                    ),
                    const SizedBox(height: 8),
                    RequirementResetPassword(
                      icon: Icons.cancel,
                      color: dangerColor,
                      text: "Password minimal 1 huruf kapital",
                    ),
                    const SizedBox(height: 8),
                    RequirementResetPassword(
                      icon: Icons.cancel,
                      color: dangerColor,
                      text: "Konfirmasi password tidak sesuai",
                    ),
                  ],
                ),

                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton(
                    onPressed:() {
                  
                    },
                    style: ElevatedButton.styleFrom(
                      backgroundColor: primaryColor,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(20),
                      ),
                      padding: const EdgeInsets.symmetric(vertical: 16),
                    ),
                    child: Text(
                      "Selesai",
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
          ),
        ),
      ),
    );
  }
}
