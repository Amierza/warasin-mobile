import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:google_fonts/google_fonts.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: ListView(
        padding: const EdgeInsets.symmetric(
          horizontal: 50,
        ),
        children: [
          Container(
            child: Column(
              children: [
                Column(
                  children: [
                    Container(                
                      width: 174,
                      height: 144,
                      margin: const EdgeInsets.only(
                        top: 100,
                      ),
                      decoration: const BoxDecoration(
                        image: DecorationImage(
                          image: AssetImage(
                            "assets/logo-purple.png"
                          ),
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

                const SizedBox(
                  height: 30,
                ),

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
                        const SizedBox(
                          height: 8,
                        ),
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

                    const SizedBox(
                      height: 30,
                    ),
                    
                    Column(
                      children: [
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
                              borderSide: BorderSide.none
                            ),
                          ),
                        ),

                        const SizedBox(
                          height: 20,
                        ),

                        TextField(
                          obscureText: true,
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
                              borderSide: BorderSide.none
                            ),
                            suffixIcon: Icon(Icons.visibility_off),
                          ),
                        ),

                        const SizedBox(
                          height: 30,
                        ),

                        SizedBox(
                          width: double.infinity,
                          child: ElevatedButton(
                            onPressed: () {

                            },
                            style: ElevatedButton.styleFrom(
                              backgroundColor: primaryColor,
                              shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(20),
                              ),
                              padding: const EdgeInsets.symmetric(vertical: 16),
                            ),
                            child: Text(
                              "Login",
                              style: GoogleFonts.poppins(
                                fontSize: 18,
                                fontWeight: semiBold,
                                color: secondaryTextColor,       
                              ),
                            ),
                          ),
                        ),

                        TextButton(
                          onPressed: () {

                          }, 
                          child: Text(
                            "Lupa Password?",
                            style: GoogleFonts.poppins(
                              fontSize: 12,
                              fontWeight: medium,
                              color: primaryColor,
                            ),
                          ),
                        ),
                      ],
                    ),

                    const SizedBox(
                      height: 35,
                    ),

                    Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Text(
                          "Tidak punya akun?",
                          style: GoogleFonts.poppins(
                            fontSize: 14,
                            fontWeight: medium,
                            color: primaryTextColor,
                          ),
                        ),
                        TextButton(
                          onPressed: () {
                            
                          }, 
                          child: Text(
                            "Register!",
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
          )
        ],
      ),
    );
  }
}