import 'package:flutter/material.dart';
import 'package:flutter_svg/svg.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/navigation_bar.dart';
import 'package:google_fonts/google_fonts.dart';

class TerapiinPage extends StatelessWidget {
  const TerapiinPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SvgPicture.asset(
              'assets/404.svg',
              height: 160,
            ),
            const SizedBox(height: 20),
            Text(
              "404 Not Found",
              style: GoogleFonts.poppins(
                fontSize: 26,
                fontWeight: regular,
                color: primaryTextColor,
              ),
            ),
          ],
        ),
      ),
      bottomNavigationBar: const CustomNavigationBar(currentIndex: 1),
    );
  }
}
