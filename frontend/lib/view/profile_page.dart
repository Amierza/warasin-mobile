import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/header.dart';
import 'package:frontend/widget/navigation_bar.dart';
import 'package:get/get.dart';
import 'package:get_storage/get_storage.dart';
import 'package:google_fonts/google_fonts.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

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
              width: double.infinity,
              height: double.maxFinite,
              decoration: BoxDecoration(color: backgroundColor),
              child: Padding(
                padding: const EdgeInsets.symmetric(
                  horizontal: 20,
                  vertical: 30,
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      'Riwayat',
                      style: GoogleFonts.poppins(
                        fontSize: 22,
                        fontWeight: semiBold,
                        color: primaryTextColor,
                      ),
                    ),
                    MenuButton(
                      icon: Icons.history,
                      name: "Riwayat Konsultasi",
                      onPressed: () {
                        Get.toNamed('/history_consultation');
                      },
                    ),
                    MenuButton(
                      icon: Icons.people_outline,
                      name: "Riwayat Berita",
                      onPressed: () {
                        Get.toNamed('/news_history');
                      },
                    ),

                    const SizedBox(height: 60),
                    Text(
                      'Pengaturan',
                      style: GoogleFonts.poppins(
                        fontSize: 22,
                        fontWeight: semiBold,
                        color: primaryTextColor,
                      ),
                    ),
                    MenuButton(
                      icon: Icons.settings,
                      name: "Pengaturan Pribadi",
                      onPressed: () {
                        Get.toNamed('/edit_profile');
                      },
                    ),
                    MenuButton(icon: Icons.note_outlined, name: "Catatan Anda"),
                    MenuButton(
                      icon: Icons.favorite_border_outlined,
                      name: "Daftar Favorit",
                    ),
                    const SizedBox(height: 30),
                    Center(
                      child: ElevatedButton(
                        onPressed: () {
                          final box = GetStorage();
                          box.erase();
                          Get.offAllNamed('/login');
                        },
                        style: ElevatedButton.styleFrom(
                          backgroundColor: Colors.redAccent,
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(12),
                          ),
                          padding: const EdgeInsets.symmetric(
                            horizontal: 24,
                            vertical: 14,
                          ),
                        ),
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.center,
                          crossAxisAlignment: CrossAxisAlignment.center,
                          mainAxisSize: MainAxisSize.min,
                          children: [
                            const Icon(Icons.logout, color: Colors.white),
                            const SizedBox(width: 8),
                            Text(
                              'Logout',
                              style: GoogleFonts.poppins(
                                color: Colors.white,
                                fontSize: 16,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ],
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
      bottomNavigationBar: const CustomNavigationBar(currentIndex: 4),
    );
  }
}

class MenuButton extends StatelessWidget {
  final IconData icon;
  final String name;
  final VoidCallback? onPressed;

  const MenuButton({
    super.key,
    required this.icon,
    required this.name,
    this.onPressed,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onPressed,
      behavior: HitTestBehavior.opaque, // Agar area kosong juga merespon tap
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 16, horizontal: 16),
        decoration: BoxDecoration(
          border: Border(bottom: BorderSide(color: tertiaryTextColor)),
        ),
        child: Row(
          children: [
            Icon(icon, color: primaryTextColor, size: 30),
            const SizedBox(width: 16),
            Expanded(
              child: Text(
                name,
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: medium,
                  color: primaryTextColor,
                ),
              ),
            ),
            Icon(Icons.chevron_right, color: tertiaryTextColor, size: 36),
            const SizedBox(height: 20),
          ],
        ),
      ),
    );
  }
}
