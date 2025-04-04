import 'package:carousel_slider/carousel_slider.dart';
import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

List<Map<String, String>> supports = [
  {
    "image": "assets/Doctor.png",
    "title": "Ahli Profesional",
    "description":
        "cenderung mencari dukungan dari sumber profesional, seperti konselor atau psikolog, daripada teman atau keluarga.",
  },
  {
    "image": "assets/Community.png",
    "title": "Forum atau komunitas",
    "description":
        "cenderung mencari dukungan melalui komunitas online atau forum di internet.",
  },
  {
    "image": "assets/Alone.png",
    "title": "Mandiri",
    "description":
        "cenderung merasa tidak membutuhkan dukungan sosial saat ini.",
  },
  {
    "image": "assets/SpeciallyPeople.png",
    "title": "Kelompok khusus",
    "description":
        "cenderung mencari dukungan dari kelompok atau komunitas yang memiliki fokus pada masalah atau tantangan tertentu.",
  },
  {
    "image": "assets/Family.png",
    "title": "Hubungan sulit dengan keluarga atau teman",
    "description":
        "cenderung merasa rumit atau sulit dengan keluarga atau teman.",
  },
  {
    "image": "assets/Hobby.png",
    "title": "Aktivitas atau hobbi",
    "description":
        "cenderung mencari dukungan melalui aktivitas atau hobi yang saya senangi.",
  },
  {
    "image": "assets/Pray.png",
    "title": "Agama atau spiritual",
    "description":
        "cenderung mencari dukungan dari komunitas agama atau spiritual saya.",
  },
];

class DataCompletenessSupportsPage extends StatefulWidget {
  const DataCompletenessSupportsPage({super.key});

  @override
  State<DataCompletenessSupportsPage> createState() =>
      _DataCompletenessSupportsPageState();
}

class _DataCompletenessSupportsPageState
    extends State<DataCompletenessSupportsPage> {
  int currentIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: Padding(
        padding: const EdgeInsets.symmetric(vertical: 60, horizontal: 40),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              "Apakah kamu memiliki dukungan sosial, seperti teman atau keluarga, yang dapat Anda ajak bicara?",
              style: GoogleFonts.poppins(
                fontSize: 20,
                fontWeight: semiBold,
                color: primaryTextColor,
              ),
            ),
            const SizedBox(height: 18),
            Text(
              "Silahkan Pilih...",
              style: GoogleFonts.poppins(
                fontSize: 16,
                color: tertiaryTextColor,
              ),
            ),
            const SizedBox(height: 24),
            ConstrainedBox(
              constraints: BoxConstraints(
                maxHeight: MediaQuery.of(context).size.height * 0.7,
              ),
              child: CarouselSlider.builder(
                itemCount: supports.length,
                itemBuilder: (
                  BuildContext context,
                  int itemIndex,
                  int pageViewIndex,
                ) {
                  return SizedBox(
                    child: SupportCards(
                      image: supports[itemIndex]['image']!,
                      title: supports[itemIndex]['title']!,
                      description: supports[itemIndex]['description']!,
                    ),
                  );
                },
                options: CarouselOptions(
                  height: 500,
                  viewportFraction: 0.9,
                  enableInfiniteScroll: true,
                  enlargeCenterPage: true,
                  onPageChanged: (index, reason) {
                    setState(() {
                      currentIndex = index;
                    });
                  },
                ),
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
                        Get.toNamed('/data_completeness_profile');
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
    );
  }
}

class SupportCards extends StatelessWidget {
  final String image;
  final String title;
  final String description;

  const SupportCards({
    Key? key,
    required this.image,
    required this.title,
    required this.description,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        ClipRRect(
          borderRadius: const BorderRadius.vertical(top: Radius.circular(16)),
          child: Container(
            width: 250,
            height: 280,
            decoration: BoxDecoration(color: primaryColor),
            child: Image.asset(
              image,
              width: 300,
              height: 330,
              fit: BoxFit.cover,
            ),
          ),
        ),
        Container(
          decoration: BoxDecoration(
            color: primaryColor,
            borderRadius: BorderRadius.vertical(bottom: Radius.circular(16)),
          ),
          padding: const EdgeInsets.all(18),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                title,
                style: GoogleFonts.poppins(
                  fontSize: 16,
                  fontWeight: medium,
                  color: Colors.white,
                ),
              ),
              const SizedBox(height: 8),
              Text(
                description,
                style: GoogleFonts.poppins(
                  fontSize: 14,
                  fontWeight: light,
                  color: Colors.white,
                ),
              ),
            ],
          ),
        ),
      ],
    );
  }
}
