import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:google_fonts/google_fonts.dart';

class DetailConcultationPage extends StatefulWidget {
  const DetailConcultationPage({super.key});

  @override
  State<DetailConcultationPage> createState() => _DetailConcultationPageState();
}

class _DetailConcultationPageState extends State<DetailConcultationPage> {
  int selectedDateIndex = 0;
  int selectedTimeIndex = 1;

  List<String> dates = ["23", "24", "25", "26"];
  List<String> months = ["Mei", "Mei", "Mei", "Mei"];
  List<String> times = ["10.00 - 11.00", "12.00 - 13.00", "14.00 - 15.00"];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leadingWidth: 60,
        leading: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: IconButton(
            icon: Icon(Icons.arrow_back_ios, color: primaryColor),
            onPressed: () => Navigator.pop(context),
          ),
        ),
        title: Text(
          "Detail Konselor",
          style: GoogleFonts.poppins(color: primaryTextColor, fontWeight: bold),
        ),
        centerTitle: true,
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 10),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                ClipRRect(
                  borderRadius: BorderRadius.circular(10),
                  child: Image.asset(
                    'assets/berita1.png',
                    width: 100,
                    height: 100,
                    fit: BoxFit.cover,
                  ),
                ),
                SizedBox(width: 16),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        "Dr. Adolfus Brustisus",
                        style: GoogleFonts.poppins(
                          fontSize: 18,
                          fontWeight: bold,
                          color: primaryTextColor,
                        ),
                      ),
                      Text(
                        "Konselor - RSUD Soetomo",
                        style: GoogleFonts.poppins(
                          fontSize: 14,
                          color: primaryTextColor,
                        ),
                      ),
                      SizedBox(height: 8),
                      Row(
                        children: [
                          Icon(Icons.star, color: Colors.amber, size: 26),
                          SizedBox(width: 5),
                          Text("4.9", style: GoogleFonts.poppins(fontSize: 16)),
                        ],
                      ),
                    ],
                  ),
                ),
              ],
            ),
            SizedBox(height: 40),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                _infoCard(Icons.people, "328+", "Pasien"),
                _infoCard(Icons.bar_chart_outlined, "8+", "Tahun Kerja"),
                _infoCard(Icons.star_outlined, "4.9", "Rating"),
                _infoCard(Icons.chat_rounded, "300+", "Ulasan"),
              ],
            ),
            SizedBox(height: 40),
            Text(
              "Tentang Konselor",
              style: GoogleFonts.poppins(fontSize: 22, fontWeight: bold),
            ),
            Text(
              "Saya seorang dokter berpengalaman dengan pengalaman praktik lebih dari",
              style: GoogleFonts.poppins(fontSize: 14, color: primaryTextColor),
            ),
            SizedBox(height: 40),
            Text(
              "Jadwal",
              style: GoogleFonts.poppins(
                fontSize: 22,
                fontWeight: bold,
                color: primaryTextColor,
              ),
            ),
            const SizedBox(height: 10),
            SingleChildScrollView(
              scrollDirection: Axis.horizontal, // Bisa digeser ke kanan/kiri
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: List.generate(dates.length, (index) {
                  return Padding(
                    padding: EdgeInsets.only(right: 16), // Jarak antar elemen
                    child: GestureDetector(
                      onTap: () {
                        setState(() {
                          selectedDateIndex = index;
                        });
                      },
                      child: _dateCard(
                        dates[index],
                        months[index],
                        selectedDateIndex == index,
                      ),
                    ),
                  );
                }),
              ),
            ),
            const SizedBox(height: 20),
            Text(
              "Waktu",
              style: GoogleFonts.poppins(fontSize: 22, fontWeight: bold),
            ),
            const SizedBox(height: 10),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: List.generate(times.length, (index) {
                return GestureDetector(
                  onTap: () {
                    setState(() {
                      selectedTimeIndex = index;
                    });
                  },
                  child: _timeCard(times[index], selectedTimeIndex == index),
                );
              }),
            ),
            const SizedBox(height: 60),
            ElevatedButton(
              onPressed: () {},
              style: ElevatedButton.styleFrom(
                backgroundColor: primaryColor,
                minimumSize: Size(double.infinity, 50),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(10),
                ),
              ),
              child: Text(
                "Pesan Konselor Sekarang",
                style: GoogleFonts.poppins(
                  fontSize: 16,
                  fontWeight: bold,
                  color: Colors.white,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

Widget _infoCard(IconData icon, String value, String label) {
  return Column(
    children: [
      Container(
        decoration: BoxDecoration(
          color: secondaryTextColor,
          borderRadius: BorderRadius.circular(100),
        ),
        child: Padding(
          padding: const EdgeInsets.all(6),
          child: Icon(icon, color: primaryColor, size: 30),
        ),
      ),
      const SizedBox(height: 8),
      Text(value, style: GoogleFonts.poppins(fontSize: 18, fontWeight: bold)),
      Text(
        label,
        style: GoogleFonts.poppins(fontSize: 14, color: primaryTextColor),
      ),
    ],
  );
}

Widget _dateCard(String day, String month, bool isSelected) {
  return Container(
    padding: EdgeInsets.all(20),
    decoration: BoxDecoration(
      color: isSelected ? primaryColor : Colors.grey[300],
      borderRadius: BorderRadius.circular(10),
    ),
    child: Column(
      children: [
        Text(
          day,
          style: GoogleFonts.poppins(
            fontSize: 20,
            fontWeight: FontWeight.bold,
            color: isSelected ? Colors.white : Colors.black,
          ),
        ),
        Text(
          month,
          style: GoogleFonts.poppins(
            fontSize: 18,
            color: isSelected ? Colors.white : primaryTextColor,
          ),
        ),
      ],
    ),
  );
}

Widget _timeCard(String time, bool isSelected) {
  return Container(
    padding: EdgeInsets.all(10),
    decoration: BoxDecoration(
      color: isSelected ? primaryColor : Colors.grey[300],
      borderRadius: BorderRadius.circular(10),
    ),
    child: Text(
      time,
      style: GoogleFonts.poppins(
        fontSize: 14,
        fontWeight: semiBold,
        color: isSelected ? Colors.white : primaryTextColor,
      ),
    ),
  );
}
