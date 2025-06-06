import 'package:flutter/material.dart';
import 'package:frontend/controller/consultation/get_all_psycholog.dart';
import 'package:frontend/model/psycholog.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/header.dart';
import 'package:frontend/widget/navigation_bar.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class ConcultationPage extends StatelessWidget {
  const ConcultationPage({super.key});

  @override
  Widget build(BuildContext context) {
    final controller = Get.put(GetAllPsycholog());

    return Scaffold(
      extendBody: true,
      backgroundColor: primaryColor,
      body: SingleChildScrollView(
        padding: EdgeInsets.only(bottom: 36),
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
                      hintText: "Temukan Konselormu...",
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
                child: Obx(() {
                  if (controller.isLoading.value) {
                    return const Center(child: CircularProgressIndicator());
                  }

                  return Column(
                    children: [
                      Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text(
                            "Daftar Konselor",
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
                      const SizedBox(height: 14),
                      ...controller.psychologList.map(
                        (psy) => Padding(
                          padding: const EdgeInsets.only(bottom: 14),
                          child: CounselorCard(
                            psycholog: psy,
                            onTap: () {
                              Get.toNamed(
                                '/concultation_detail',
                              ); // will be handled later
                            },
                          ),
                        ),
                      ),
                    ],
                  );
                }),
              ),
            ),
          ],
        ),
      ),
      bottomNavigationBar: const CustomNavigationBar(currentIndex: 2),
    );
  }
}

class CounselorCard extends StatelessWidget {
  final Psycholog psycholog;
  final VoidCallback onTap;

  const CounselorCard({Key? key, required this.psycholog, required this.onTap})
    : super(key: key);

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: onTap,
      borderRadius: BorderRadius.circular(20),
      child: Card(
        color: backgroundColor,
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20)),
        child: Padding(
          padding: const EdgeInsets.all(14),
          child: Row(
            children: [
              ClipRRect(
                borderRadius: BorderRadius.circular(15),
                child: Image.network(
                  psycholog.psyImage ?? '',
                  width: 100,
                  height: 100,
                  fit: BoxFit.cover,
                  errorBuilder: (context, error, stackTrace) {
                    return Image.asset(
                      'assets/Alone.png',
                      width: 100,
                      height: 100,
                      fit: BoxFit.cover,
                    );
                  },
                ),
              ),

              const SizedBox(width: 16),
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(
                      psycholog.psyName,
                      style: GoogleFonts.poppins(
                        fontSize: 18,
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                    const Divider(),
                    Row(
                      children: [
                        Text(
                          psycholog.city.cityName,
                          style: GoogleFonts.poppins(fontSize: 14),
                        ),
                        Padding(
                          padding: const EdgeInsets.symmetric(horizontal: 8),
                          child: VerticalDivider(
                            thickness: 1,
                            color: Colors.black54,
                          ),
                        ),
                        Text(
                          "${psycholog.psyWorkYear} Tahun",
                          style: GoogleFonts.poppins(fontSize: 14),
                        ),
                      ],
                    ),
                    const SizedBox(height: 10),
                    Wrap(
                      spacing: 6,
                      runSpacing: 6,
                      children:
                          psycholog.specializations.map((e) {
                            return Container(
                              padding: const EdgeInsets.symmetric(
                                vertical: 4,
                                horizontal: 8,
                              ),
                              decoration: BoxDecoration(
                                color: Colors.blue.shade100,
                                borderRadius: BorderRadius.circular(12),
                              ),
                              child: Text(
                                e.speName,
                                style: GoogleFonts.poppins(
                                  fontSize: 12,
                                  color: Colors.blue.shade800,
                                ),
                              ),
                            );
                          }).toList(),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
