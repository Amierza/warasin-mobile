import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:intl/intl.dart';

class NewsCard extends StatelessWidget {
  final String image;
  final String title;
  final String desc;
  final String date;
  final VoidCallback? onTap;

  const NewsCard({
    Key? key,
    required this.image,
    required this.title,
    required this.desc,
    required this.date,
    this.onTap,
  }) : super(key: key);

  String formatDate(String dateString) {
    try {
      final parsedDate = DateTime.parse(dateString);
      return DateFormat('dd-MM-yyyy').format(parsedDate);
    } catch (e) {
      return dateString;
    }
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap ?? () => Get.toNamed("/news/:id"),
      child: Container(
        margin: const EdgeInsets.symmetric(vertical: 8, horizontal: 16),
        padding: const EdgeInsets.all(12),
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(16),
          boxShadow: [
            BoxShadow(
              color: Colors.black.withOpacity(0.05),
              blurRadius: 6,
              spreadRadius: 2,
            ),
          ],
        ),
        child: Row(
          children: [
            ClipRRect(
              borderRadius: BorderRadius.circular(12),
              child: image.isNotEmpty
                  ? Image.network(
                      image,
                      width: 100,
                      height: 100,
                      fit: BoxFit.cover,
                      errorBuilder: (context, error, stackTrace) {
                        return Image.asset(
                          'assets/default_image.jpg',
                          width: 80,
                          height: 80,
                          fit: BoxFit.cover,
                        );
                      },
                    )
                  : Image.asset(
                      'assets/default_image.jpg',
                      width: 80,
                      height: 80,
                      fit: BoxFit.cover,
                    ),
            ),
            const SizedBox(width: 12),
            Expanded(
              child: SizedBox(
                height: 100,
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      formatDate(date),
                      style: GoogleFonts.poppins(
                        fontSize: 11,
                        color: Colors.grey,
                        fontWeight: light,
                      ),
                    ),
                    const SizedBox(height: 4),
                    Text(
                      title,
                      maxLines: 2,
                      overflow: TextOverflow.ellipsis,
                      style: GoogleFonts.poppins(
                        fontSize: 14,
                        fontWeight: bold,
                      ),
                    ),
                    const SizedBox(height: 4),
                    Expanded(
                      child: Text(
                        desc,
                        maxLines: 2,
                        overflow: TextOverflow.ellipsis,
                        style: GoogleFonts.poppins(
                          fontSize: 12,
                          color: Colors.grey,
                          fontWeight: light,
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
    );
  }
}
