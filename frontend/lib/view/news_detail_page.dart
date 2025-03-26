import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:google_fonts/google_fonts.dart';

class NewsDetailPage extends StatelessWidget {
  const NewsDetailPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leadingWidth: 60,
        leading: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: IconButton(
            icon: Icon(
              Icons.arrow_back_ios, 
              color:primaryColor,
            ),
            onPressed: () => Navigator.pop(context),
          ),
        ),
        title: Text(
          "Detail Berita",
          style: GoogleFonts.poppins(
            color: primaryTextColor, 
            fontWeight: bold
          ),
        ),
        centerTitle: true,
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: SingleChildScrollView(
        padding: const EdgeInsets.symmetric(
          horizontal: 30, 
          vertical: 10
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            ClipRRect(
              borderRadius: BorderRadius.circular(20),
              child: Image.asset(
                "assets/berita1.png",
                width: double.infinity,
                height: 200,
                fit: BoxFit.cover,
              ),
            ),
            const SizedBox(height: 16),
            Text(
              "Skrining Kesehatan Mental dalam Program CKG Berlangsung Singkat",
              style: GoogleFonts.poppins(
                color: primaryTextColor,
                fontSize: 18,
                fontWeight: bold,
              ),
            ),
            const SizedBox(height: 8),
            Row(
              children: [
                const CircleAvatar(
                  backgroundImage: AssetImage("assets/Alone.png"),
                  radius: 24,
                ),
                const SizedBox(width: 8),
                Text(
                  "By Rudi Hartono",
                  style: GoogleFonts.poppins(
                    fontSize: 13,
                    fontWeight: light,
                    color: primaryTextColor,
                  ),
                ),
                const Spacer(),
                Text(
                  "10-02-2025 18:56:20",
                  style: GoogleFonts.poppins(
                    fontSize: 11,
                    fontWeight: light,
                    color: primaryTextColor,
                  ),
                ),
              ],
            ),
            const SizedBox(height: 16),
            Text(
              '''Pelaksanaan Cek Kesehatan Gratis (CKG) hari pertama di Puskesmas Beji, Depok, Jawa Barat Senin (10/2) turut mencakup pemeriksaan kesehatan mental selain kesehatan fisik peserta. Penanggungjawab CKG Puskesmas Beji, dr. Asmarini Ratnaningsih mengatakan pemeriksaan kesehatan mental itu menjadi satu rangkaian CKG.
                  
                  Ia menjelaskan pemeriksaan kesehatan mental itu sudah biasa dilakukan terhadap pasien yang melakukan pemeriksaan kesehatan di Puskesmas Beji.

                  "Kalaupun kami tidak CKG skrining kesehatan mental kita tetap lakukan jadi bukan hanya pasien CKG saja. Diluar pasien ckg pun kami melakukan pemeriksaan kesehatan mental," kata Asmarini di Puskesmas Beji.''',
              textAlign: TextAlign.justify,
              style: GoogleFonts.poppins(fontSize: 16, height: 1.5, color: primaryTextColor),
            ),
          ],
        ),
      ),
    );
  }
}
