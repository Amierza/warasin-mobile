import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

List<Map<String, String>> problems = [
  {
    "title": "Tidak ada perubahan signifikan",
    "desc":
        "Kamu dalam keadaan yang relatif stabil dan tidak mengalami perubahan besar dalam pekerjaan, hubungan, atau kehidupan secara keseluruhan.",
  },
  {
    "title": "Ada beberapa perubahan pekerjaan atau hubungan",
    "desc":
        "Kamu dalam keadaan mengalami perubahan baik positif maupun menantang.",
  },
  {
    "title": "Menghadapi tantangan besar",
    "desc":
        "Kamu dalam keadaan menghadapi tantangan besar seperti kehilangan pekerjaan yang signifikan atau kehilangan orang yang dicintai.",
  },
];

class DataCompletenessProblemsPage extends StatefulWidget {
  const DataCompletenessProblemsPage({super.key});

  @override
  State<DataCompletenessProblemsPage> createState() =>
      _DataCompletenessProblemsPageState();
}

class _DataCompletenessProblemsPageState
    extends State<DataCompletenessProblemsPage> {
  int? selectedIndex;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: Center(
        child: Padding(
          padding: const EdgeInsets.symmetric(vertical: 60, horizontal: 40),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              Text(
                "Apa yang menjadi perubahan atau tantangan dalam hidup kamu belakangan ini?",
                style: GoogleFonts.poppins(fontSize: 20, fontWeight: semiBold),
              ),
              Expanded(
                child: ListView.separated(
                  itemCount: problems.length,
                  separatorBuilder:
                      (context, index) => const SizedBox(height: 14),
                  itemBuilder: (context, index) {
                    return GestureDetector(
                      onTap: () {
                        setState(() {
                          selectedIndex = index;
                        });
                      },
                      child: ProblemCards(
                        title: problems[index]['title']!,
                        desc: problems[index]['desc']!,
                        isSelected:
                            selectedIndex == index,
                      ),
                    );
                  },
                ),
              ),
              SizedBox(
                width: double.infinity,
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Expanded(
                      child: Opacity(
                        opacity: 0.7,
                        child: ElevatedButton(
                          onPressed: () {
                            Get.back();
                          },
                          style: ElevatedButton.styleFrom(
                            backgroundColor: primaryColor,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(20),
                            ),
                            padding: const EdgeInsets.symmetric(vertical: 16),
                          ),
                          child: Text(
                            "Kembali",
                            style: GoogleFonts.poppins(
                              fontSize: 18,
                              fontWeight: semiBold,
                              color: secondaryTextColor,
                            ),
                          ),
                        ),
                      ),
                    ),
                    const SizedBox(width: 10),
                    Expanded(
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
            ],
          ),
        ),
      ),
    );
  }
}

class ProblemCards extends StatelessWidget {
  final String title;
  final String desc;
  final bool isSelected;

  const ProblemCards({
    Key? key,
    required this.title,
    required this.desc,
    this.isSelected = false,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AnimatedOpacity(
      opacity: isSelected ? 1.0 : 0.6,
      duration: const Duration(milliseconds: 200),
      child: Container(
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(16),
          color: primaryColor,
          border: isSelected
              ? Border.all(color: Colors.white, width: 2)
              : null,
        ),
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              title,
              style: GoogleFonts.poppins(
                fontSize: 16,
                fontWeight: regular,
                color: Colors.white,
              ),
            ),
            const SizedBox(height: 6),
            Text(
              desc,
              style: GoogleFonts.poppins(
                fontSize: 14,
                color: Colors.white,
              ),
            ),
          ],
        ),
      ),
    );
  }
}