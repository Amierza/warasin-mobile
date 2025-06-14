import 'package:carousel_slider/carousel_slider.dart';
import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/header.dart';
import 'package:frontend/widget/navigation_bar.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

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
              child: Header(),
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
                    MoodSection(),
                    const SizedBox(height: 20),
                    TrendingSection(),
                    const SizedBox(height: 20),
                    motivationSection(),
                    const SizedBox(height: 20),
                    ProgressSection(),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
      bottomNavigationBar: const CustomNavigationBar(currentIndex: 0),
    );
  }
}

class MoodSection extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final Color customPink = Color(0xFFEF5DA8);
    final Color customPurple = Color(0xFFAEAFF7);
    final Color customAqua = Color(0xFFA0E3E2);
    final Color customOrange = Color(0xFFF09E54);
    final Color customGreen = Color(0xFFC3F2A6);

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'How are you feeling today?',
          style: GoogleFonts.poppins(
            fontSize: 16,
            fontWeight: semiBold,
            color: primaryTextColor,
          ),
        ),
        const SizedBox(height: 10),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: [
            MoodItem(
              color: customPink,
              image: 'assets/Happy.png',
              label: 'Happy',
            ),
            MoodItem(
              color: customPurple,
              image: 'assets/Calm.png',
              label: 'Calm',
            ),
            MoodItem(
              color: customAqua,
              image: 'assets/Relax.png',
              label: 'Manic',
            ),
            MoodItem(
              color: customOrange,
              image: 'assets/Angry.png',
              label: 'Angry',
            ),
            MoodItem(color: customGreen, image: 'assets/Sad.png', label: 'Sad'),
          ],
        ),
      ],
    );
  }
}

class MoodItem extends StatelessWidget {
  final String image;
  final String label;
  final Color color;

  MoodItem({required this.color, required this.image, required this.label});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          decoration: BoxDecoration(
            color: color,
            borderRadius: BorderRadius.circular(20),
          ),
          child: Padding(
            padding: const EdgeInsets.all(12),
            child: Image.asset(
              image,
              width: 40,
              height: 40,
              color: Colors.white,
            ),
          ),
        ),
        const SizedBox(height: 5),
        Text(label, style: GoogleFonts.poppins(fontSize: 12)),
      ],
    );
  }
}

List<Map<String, String>> trendingItems = [
  {
    "image": "assets/berita1.png",
    "title": "10 Rekomendasi Musik Klasik yang Cocok untuk Relaksasi",
  },
  {
    "image": "assets/berita2.png",
    "title":
        "Halo GenZ, Seberapa Sehatkah Mentalmu? Ini Cara Mendeteksinya dari Psikolog.",
  },
  {
    "image": "assets/berita3.png",
    "title":
        "Introvert Susah Bergaul karena 'Social Battery' Cepat Habis? Begini Kata Psikolog.",
  },
];

class TrendingSection extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Yang lagi rame nih...',
          style: GoogleFonts.poppins(fontSize: 16, fontWeight: FontWeight.bold),
        ),
        const SizedBox(height: 10),
        CarouselSlider(
          items:
              trendingItems.map((item) {
                return TrendingItem(
                  title: item['title']!,
                  image: item['image']!,
                );
              }).toList(),
          options: CarouselOptions(
            height: 150,
            autoPlay: true,
            autoPlayInterval: Duration(seconds: 4),
            enlargeCenterPage: true,
            viewportFraction: 0.8,
          ),
        ),
      ],
    );
  }
}

class TrendingItem extends StatelessWidget {
  final String title;
  final String image;

  TrendingItem({required this.title, required this.image});

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.symmetric(horizontal: 8),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(15),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.1),
            blurRadius: 10,
            spreadRadius: 3,
          ),
        ],
      ),
      child: Row(
        children: [
          ClipRRect(
            borderRadius: BorderRadius.only(
              topLeft: Radius.circular(15),
              bottomLeft: Radius.circular(15),
            ),
            child: Image.asset(
              image,
              width: 130,
              height: 180,
              fit: BoxFit.cover,
            ),
          ),
          Expanded(
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 10),
              child: Text(
                title,
                style: GoogleFonts.poppins(
                  fontSize: 11,
                  fontWeight: semiBold,
                  color: primaryTextColor,
                ),
                maxLines: 3,
                overflow: TextOverflow.ellipsis,
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class ProgressSection extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          'Pantau progressmu yuk!!',
          style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
        ),
        const SizedBox(height: 10),
        GridView.count(
          shrinkWrap: true,
          physics: NeverScrollableScrollPhysics(),
          crossAxisCount: 2,
          crossAxisSpacing: 10,
          mainAxisSpacing: 10,
          children: [
            ProgressItem(
              title: 'Butuh Semangat?',
              subtitle: 'Baca Motivasi Harian',
              color: Colors.teal,
              onTap: () {
                Get.toNamed('/motivation');
              },
            ),
            ProgressItem(
              title: 'Curhat Yuk',
              subtitle: 'Konsultasi Sekarang',
              color: Colors.green,
              onTap: () {
                Get.toNamed('/concultation');
              },
            ),
            ProgressItem(
              title: 'Info Terbaru',
              subtitle: 'Baca Berita Kesehatan',
              color: Colors.orange,
              onTap: () {
                Get.toNamed('/news');
              },
            ),
            ProgressItem(
              title: 'Lengkapi Profil',
              subtitle: 'Biar Lebih Personal',
              color: Colors.deepPurple,
              onTap: () {
                Get.toNamed('/edit_profile');
              },
            ),
          ],
        ),
      ],
    );
  }
}

class motivationSection extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          "“Penyakit kamu bukanlah identitasmu. Chemistry bukanlah karaktermu.” ",
          style: GoogleFonts.poppins(
            fontSize: 14,
            fontWeight: semiBold,
            color: primaryTextColor,
          ),
        ),
        Text(
          "-Rick Warren",
          style: GoogleFonts.poppins(
            fontSize: 14,
            fontWeight: regular,
            color: primaryTextColor,
          ),
        ),
      ],
    );
  }
}

class ProgressItem extends StatelessWidget {
  final String title;
  final String subtitle;
  final Color color;
  final VoidCallback? onTap;

  const ProgressItem({
    Key? key,
    required this.title,
    required this.subtitle,
    required this.color,
    this.onTap,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: onTap,
      borderRadius: BorderRadius.circular(10),
      child: Container(
        padding: const EdgeInsets.all(20),
        decoration: BoxDecoration(
          color: color,
          borderRadius: BorderRadius.circular(10),
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              title,
              style: GoogleFonts.poppins(
                fontSize: 18,
                fontWeight: FontWeight.bold,
                color: Colors.white,
              ),
            ),
            const SizedBox(height: 5),
            Text(
              subtitle,
              style: const TextStyle(fontSize: 14, color: Colors.white),
            ),
          ],
        ),
      ),
    );
  }
}
