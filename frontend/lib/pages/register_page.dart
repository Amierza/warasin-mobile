import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:google_fonts/google_fonts.dart';

class RegisterPage extends StatefulWidget {
  const RegisterPage({super.key});

  @override
  State<RegisterPage> createState() => _RegisterPage();
}

class _RegisterPage extends State<RegisterPage> {
  bool _isObsecure = true;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: ListView(
        padding: const EdgeInsets.symmetric(horizontal: 50),
        children: [
          Column(
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
                  fontWeight: extraBold,
                  color: primaryColor,
                ),
              ),
            ],
          ),

          const SizedBox(height: 30),

          // Bagian Form Login
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Column(
                children: [
                  Text(
                    "Selamat Datang di Warasin!",
                    style: GoogleFonts.poppins(
                      fontSize: 20,
                      fontWeight: FontWeight.w700,
                      color: primaryTextColor,
                    ),
                    textAlign: TextAlign.center,
                  ),
                  const SizedBox(height: 8),
                  Text(
                    "Jaga kesehatan mentalmu bersama Warasin",
                    style: GoogleFonts.poppins(
                      fontSize: 16,
                      fontWeight: FontWeight.w600,
                      color: tertiaryTextColor,
                    ),
                    textAlign: TextAlign.center,
                  ),
                ],
              ),

              const SizedBox(height: 30),

              Column(
                children: [
                  TextField(
                    decoration: InputDecoration(
                      hintText: "Nama",
                      filled: true,
                      fillColor: secondaryTextColor,
                      contentPadding: EdgeInsets.symmetric(
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

                  TextField(
                    decoration: InputDecoration(
                      hintText: "Email",
                      filled: true,
                      fillColor: secondaryTextColor,
                      contentPadding: EdgeInsets.symmetric(
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

                  TextField(
                    obscureText: _isObsecure,
                    decoration: InputDecoration(
                      hintText: "Password",
                      filled: true,
                      fillColor: secondaryTextColor,
                      contentPadding: EdgeInsets.symmetric(
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
                            _isObsecure = !_isObsecure;
                          });
                        }, 
                        icon: Icon(
                          _isObsecure ? Icons.visibility_off : Icons.visibility
                        ),
                      ),
                    ),
                  ),

                  const SizedBox(height: 24),

                  SizedBox(
                    width: double.infinity,
                    child: ElevatedButton(
                      onPressed: () {},
                      style: ElevatedButton.styleFrom(
                        backgroundColor: primaryColor,
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(20),
                        ),
                        padding: const EdgeInsets.symmetric(vertical: 16),
                      ),
                      child: Text(
                        "Register",
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
                      Navigator.pushNamedAndRemoveUntil(
                        context,
                        '/login',
                        (route) => false,
                      );
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
        ],
      ),
    );
  }
}
