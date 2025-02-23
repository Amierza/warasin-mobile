import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:google_fonts/google_fonts.dart';

class LogoWarasin extends StatelessWidget {
  const LogoWarasin({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Container(
          width: 174,
          height: 144,
          margin: const EdgeInsets.only(top: 25),
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
      ],
    );
  }
}