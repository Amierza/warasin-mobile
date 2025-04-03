import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class DataCompletenessGreetingsPage extends StatelessWidget {
  const DataCompletenessGreetingsPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: Center(
        child: Padding(
          padding: const EdgeInsets.symmetric(
            vertical: 150,
            horizontal: 50,
            
          ),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Image.asset("assets/Welcome.png", height: 210, width: 250),
              Text(
                "Halo, Budi!",
                style: GoogleFonts.poppins(fontSize: 18, fontWeight: semiBold),
              ),
              Text(
                textAlign: TextAlign.center,
                "Warasin akan memberikan beberapa pertanyaan buat kamu nih! Jangan di skip yah, agar Warasin bisa memberikan solusi untuk kamu!",
                style: GoogleFonts.poppins(fontSize: 14, fontWeight: light),
              ),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed: () {
                    Get.toNamed('/data_completeness_feels');
                  },
                  style: ElevatedButton.styleFrom(
                    backgroundColor: primaryColor,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(20),
                    ),
                    padding: const EdgeInsets.symmetric(vertical: 16),
                  ),
                  child: Text(
                    "Mulai",
                    style: GoogleFonts.poppins(
                      fontSize: 18,
                      fontWeight: semiBold,
                      color: secondaryTextColor,
                    ),
                  ),
                ),
              ),
              Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Container(
                    width: 80,
                    height: 80,
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
                      fontSize: 20,
                      fontWeight: semiBold,
                      color: primaryColor,
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
