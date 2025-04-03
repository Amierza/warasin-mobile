import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/feels_cards.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class DataCompletenessFeelsPage extends StatelessWidget {
  const DataCompletenessFeelsPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: Center(
        child: Padding(
          padding: const EdgeInsets.symmetric(vertical: 60, horizontal: 50),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Text(
                "Bagaimana perasaanmu saat ini?",
                style: GoogleFonts.poppins(fontSize: 20, fontWeight: semiBold),
              ),
              Expanded(
                // Tambahkan Expanded agar GridView tidak error
                child: GridView.count(
                  crossAxisCount: 2,
                  crossAxisSpacing: 20,
                  mainAxisSpacing: 20,
                  children: [
                    FeelCards(
                      image: 'assets/Smiling_face.png',
                      label: 'Senang',
                    ),
                    FeelCards(
                      image: 'assets/Disappointed_face.png',
                      label: 'Sedih',
                    ),
                    FeelCards(
                      image: 'assets/Pouting_face.png',
                      label: 'Marah',
                    ),
                    FeelCards(
                      image: 'assets/Disappointed_but_relieved_face.png',
                      label: 'Bingung',
                    ),
                  ],
                ),
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
                    "Lanjut",
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
    );
  }
}