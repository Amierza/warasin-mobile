import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/feels_cards.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

List<Map<String, String>> feels = [
  {"image": "assets/Smiling_face.png", "label": "Senang"},
  {"image": "assets/Disappointed_face.png", "label": "Sedih"},
  {"image": "assets/Pouting_face.png", "label": "Marah"},
  {"image": "assets/Disappointed_but_relieved_face.png", "label": "Bingung"},
];

class DataCompletenessFeelsPage extends StatefulWidget {
  const DataCompletenessFeelsPage({super.key});

  @override
  State<DataCompletenessFeelsPage> createState() =>
      _DataCompletenessFeelsPageState();
}

class _DataCompletenessFeelsPageState extends State<DataCompletenessFeelsPage> {
  int? selectedIndex;

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
                child: GridView.builder(
                  gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: 2,
                    crossAxisSpacing: 16,
                    mainAxisSpacing: 16,
                  ),
                  itemCount: feels.length,
                  itemBuilder: (context, index) {
                    return GestureDetector(
                      onTap: () {
                        setState(() {
                          selectedIndex = index;
                        });
                      },
                      child: FeelCards(
                        image: feels[index]['image']!,
                        label: feels[index]['label']!,
                        isSelected: selectedIndex == index,
                      ),
                    );
                  },
                ),
              ),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed: () {
                    Get.toNamed('/data_completeness_problems');
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
