import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/header.dart';
import 'package:frontend/widget/navigation_bar.dart';
import 'package:frontend/widget/news_card.dart';
import 'package:google_fonts/google_fonts.dart';

List<Map<String, String>> newList = [
  {
    "image": "assets/Alone.png",
    "title": "Halo GenZ, Seberapa Sehatkah Mentalmu? Ini...",
    "description":
        "Sekarang adalah waktu yang tepat untuk memperbaiki masalahmu!",
    "date": "25-04-2025 15:05:08",
  },
  {
    "image": "assets/berita2.png",
    "title": "Cara Mengatasi Stres dalam Kehidupan Sehari-hari",
    "description":
        "Tips dan trik agar tetap tenang menghadapi tantangan hidup.",
    "date": "26-04-2025 10:30:00",
  },
  {
    "image": "assets/berita3.png",
    "title": "Manfaat Meditasi untuk Kesehatan Mental",
    "description":
        "Mengapa meditasi bisa membantu mengurangi stres dan kecemasan.",
    "date": "27-04-2025 08:15:00",
  },
  {
    "image": "assets/Alone.png",
    "title": "Halo GenZ, Seberapa Sehatkah Mentalmu? Ini...",
    "description":
        "Sekarang adalah waktu yang tepat untuk memperbaiki masalahmu!",
    "date": "25-04-2025 15:05:08",
  },
  {
    "image": "assets/berita2.png",
    "title": "Cara Mengatasi Stres dalam Kehidupan Sehari-hari",
    "description":
        "Tips dan trik agar tetap tenang menghadapi tantangan hidup.",
    "date": "26-04-2025 10:30:00",
  },
  {
    "image": "assets/berita3.png",
    "title": "Manfaat Meditasi untuk Kesehatan Mental",
    "description":
        "Mengapa meditasi bisa membantu mengurangi stres dan kecemasan.",
    "date": "27-04-2025 08:15:00",
  },
];

class NewsPage extends StatelessWidget {
  const NewsPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      extendBody: true,
      backgroundColor: primaryColor,
      body: SingleChildScrollView(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Padding(
              padding: const EdgeInsets.only(top: 40, left: 20, right: 20),
              child: Column(
                children: [
                  Header(),
                  const SizedBox(height: 20),
                  TextFormField(
                    style: GoogleFonts.poppins(color: backgroundColor),
                    decoration: InputDecoration(
                      hintText: "Temukan Berita Menarik...",
                      hintStyle: GoogleFonts.poppins(color: backgroundColor),
                      prefixIcon: Icon(
                        Icons.search_rounded,
                        color: backgroundColor,
                        size: 25,
                      ),
                      contentPadding: const EdgeInsets.symmetric(
                        horizontal: 20,
                        vertical: 20,
                      ),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: const BorderSide(
                          width: 2,
                          color: Colors.white,
                        ),
                      ),
                      enabledBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: const BorderSide(
                          width: 2,
                          color: Colors.white,
                        ),
                      ),
                      focusedBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: const BorderSide(
                          width: 2,
                          color: Colors.white,
                        ),
                      ),
                    ),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 20),
            Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.vertical(top: Radius.circular(30)),
                color: backgroundColor,
              ),
              child: Padding(
                padding: const EdgeInsets.symmetric(
                  horizontal: 20,
                  vertical: 30,
                ),
                child: Column(
                  children: [
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Text(
                          "Daftar Berita",
                          style: GoogleFonts.poppins(
                            fontSize: 20,
                            fontWeight: semiBold,
                          ),
                        ),
                        TextButton(
                          onPressed: () {},
                          child: Text(
                            "see all",
                            style: GoogleFonts.poppins(
                              fontSize: 14,
                              fontWeight: light,
                            ),
                          ),
                        ),
                      ],
                    ),
                    Column(
                      children:
                          newList.map((news) {
                            return NewsCard(
                              image: news["image"]!,
                              title: news["title"]!,
                              desc: news["description"]!,
                              date: news["date"]!,
                            );
                          }).toList(),
                    ),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
      bottomNavigationBar: const CustomNavigationBar(currentIndex: 3),
    );
  }
}