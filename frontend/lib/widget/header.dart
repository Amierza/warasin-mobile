import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:google_fonts/google_fonts.dart';

class Header extends StatelessWidget {
  const Header({super.key});

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        ClipRRect(
          borderRadius: BorderRadius.circular(100),
          child: SizedBox(
            height: 70,
            width: 70,
            child: Image.asset(
              'assets/Alone.png',
              fit: BoxFit.cover,
            ),
          ),
        ),
        const SizedBox(width: 20),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                'Hi, Budi',
                style: GoogleFonts.poppins(
                  fontSize: 24,
                  color: backgroundColor,
                  fontWeight: semiBold
                ),
              ),
              const SizedBox(
                height: 2,
              ),
              Text(
                'Sekarang adalah waktu yang tepat untuk memperbaiki masalahmu',
                style: GoogleFonts.poppins(
                  fontSize: 12,
                  color: backgroundColor,
                  fontWeight: regular
                ),
              ),
            ],
          ),
        )
      ],
    );
  }
}