import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:google_fonts/google_fonts.dart';

class RequirementResetPassword extends StatelessWidget {
  final IconData icon;
  final Color color;
  final String text;

  const RequirementResetPassword({
    super.key,
    required this.icon,
    required this.color,
    required this.text,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Icon(
          icon,
          color: color,
        ),
        const SizedBox(width: 10),
        Text(
          text,
          style: GoogleFonts.poppins(
            fontSize: 12,
            fontWeight: regular,
            color: primaryTextColor,
          ),
        ),
      ],
    );
  }
}