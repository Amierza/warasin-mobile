import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:google_fonts/google_fonts.dart';

class HistoryConsultationPage extends StatelessWidget {
  const HistoryConsultationPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(80),
        child: AppBar(
          leadingWidth: 60,
          leading: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: IconButton(
              icon: Icon(Icons.arrow_back_ios, color: backgroundColor),
              onPressed: () => Navigator.pop(context),
            ),
          ),
          title: Text(
            "Hisori Konsultasi",
            style: GoogleFonts.poppins(
              color: backgroundColor,
              fontWeight: bold,
            ),
          ),
          centerTitle: true,
          backgroundColor: primaryColor,
          elevation: 0,
        ),
      ),

      body: Padding(
        padding: const EdgeInsets.symmetric(vertical: 40, horizontal: 30),
        child: Column(
          children: [
            TextFormField(
              style: GoogleFonts.poppins(color: tertiaryTextColor),
              decoration: InputDecoration(
                hintText: "Bulan",
                hintStyle: GoogleFonts.poppins(color: tertiaryTextColor, fontSize: 20),
                contentPadding: const EdgeInsets.symmetric(
                  horizontal: 20,
                  vertical: 14,
                ),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20),
                  borderSide: BorderSide(width: 2, color: tertiaryTextColor),
                ),
                enabledBorder: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20),
                  borderSide: BorderSide(width: 2, color: tertiaryTextColor),
                ),
                focusedBorder: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20),
                  borderSide: BorderSide(width: 2, color: primaryTextColor),
                ),
              ),
            ),
            const SizedBox(height: 20),
            TextFormField(
              style: GoogleFonts.poppins(color: tertiaryTextColor),
              keyboardType: TextInputType.number ,
              decoration: InputDecoration(
                hintText: "Tahun",
                hintStyle: GoogleFonts.poppins(color: tertiaryTextColor, fontSize: 20),
                contentPadding: const EdgeInsets.symmetric(
                  horizontal: 20,
                  vertical: 14,
                ),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20),
                  borderSide: BorderSide(width: 2, color: tertiaryTextColor),
                ),
                enabledBorder: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20),
                  borderSide: BorderSide(width: 2, color: tertiaryTextColor),
                ),
                focusedBorder: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20),
                  borderSide: BorderSide(width: 2, color: primaryTextColor),
                ),
              ),
            ),
            const SizedBox(height: 40),
            ElevatedButton(
              onPressed: () {},
              style: ElevatedButton.styleFrom(
                backgroundColor: primaryColor,
                minimumSize: Size(double.infinity, 50),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(100),
                ),
              ),
              child: Text(
                "Cari...",
                style: GoogleFonts.poppins(
                  fontSize: 18,
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
