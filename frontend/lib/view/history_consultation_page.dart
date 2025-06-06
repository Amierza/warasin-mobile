import 'package:flutter/material.dart';
import 'package:frontend/controller/consultation/get_all_consultation.dart';
import 'package:frontend/model/consultation.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:intl/intl.dart';

class HistoryConsultationPage extends StatelessWidget {
  const HistoryConsultationPage({super.key});

  @override
  Widget build(BuildContext context) {
    final consultationController = Get.put(GetAllConsultation());
    consultationController.fetchAllConsultation(); // Fetch data when page loads

    return Scaffold(
      backgroundColor: backgroundColor,
      appBar: PreferredSize(
        preferredSize: const Size.fromHeight(80),
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
            "Histori Konsultasi",
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
            Row(
              children: [
                Expanded(
                  child: TextFormField(
                    style: GoogleFonts.poppins(color: tertiaryTextColor),
                    decoration: InputDecoration(
                      hintText: "Bulan",
                      hintStyle: GoogleFonts.poppins(
                        color: tertiaryTextColor,
                        fontSize: 20,
                      ),
                      contentPadding: const EdgeInsets.symmetric(
                        horizontal: 20,
                        vertical: 14,
                      ),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: BorderSide(
                          width: 2,
                          color: tertiaryTextColor,
                        ),
                      ),
                      enabledBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: BorderSide(
                          width: 2,
                          color: tertiaryTextColor,
                        ),
                      ),
                      focusedBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: BorderSide(
                          width: 2,
                          color: primaryTextColor,
                        ),
                      ),
                    ),
                  ),
                ),
                const SizedBox(width: 16),
                Expanded(
                  child: TextFormField(
                    style: GoogleFonts.poppins(color: tertiaryTextColor),
                    keyboardType: TextInputType.number,
                    decoration: InputDecoration(
                      hintText: "Tahun",
                      hintStyle: GoogleFonts.poppins(
                        color: tertiaryTextColor,
                        fontSize: 20,
                      ),
                      contentPadding: const EdgeInsets.symmetric(
                        horizontal: 20,
                        vertical: 14,
                      ),
                      border: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: BorderSide(
                          width: 2,
                          color: tertiaryTextColor,
                        ),
                      ),
                      enabledBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: BorderSide(
                          width: 2,
                          color: tertiaryTextColor,
                        ),
                      ),
                      focusedBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(20),
                        borderSide: BorderSide(
                          width: 2,
                          color: primaryTextColor,
                        ),
                      ),
                    ),
                  ),
                ),
              ],
            ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                // Add search functionality here
                consultationController.fetchAllConsultation();
              },
              style: ElevatedButton.styleFrom(
                backgroundColor: primaryColor,
                minimumSize: const Size(double.infinity, 50),
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
            const SizedBox(height: 20),
            Expanded(
              child: Obx(() {
                if (consultationController.isLoading.value) {
                  return const Center(child: CircularProgressIndicator());
                }

                final consultations = consultationController.consultationList;
                if (consultations.isEmpty) {
                  return Center(
                    child: Text(
                      "Belum ada histori konsultasi",
                      style: GoogleFonts.poppins(fontSize: 16),
                    ),
                  );
                }

                return ListView.builder(
                  itemCount: consultations.length,
                  itemBuilder: (context, index) {
                    final consultation = consultations[index];
                    return HistoryConsultationCard(
                      consultation: consultation,
                      onTap: () {
                        // Navigate to detail page if needed
                      },
                    );
                  },
                );
              }),
            ),
          ],
        ),
      ),
    );
  }
}

class HistoryConsultationCard extends StatelessWidget {
  final Consultation consultation;
  final VoidCallback onTap;

  const HistoryConsultationCard({
    super.key,
    required this.consultation,
    required this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    final dateFormat = DateFormat('dd MMMM yyyy');
    final formattedDate = dateFormat.format(
      DateTime.parse(consultation.consulDate),
    );

    Color statusColor;
    String statusText;

    switch (consultation.consulStatus) {
      case 0:
        statusColor = Colors.orange.shade300;
        statusText = "Upcoming";
        break;
      case 1:
        statusColor = Colors.red.shade300;
        statusText = "Cancelled";
        break;
      case 2:
        statusColor = Colors.green.shade300;
        statusText = "Done";
        break;
      default:
        statusColor = Colors.grey.shade300;
        statusText = "Unknown";
    }

    return InkWell(
      onTap: onTap,
      borderRadius: BorderRadius.circular(20),
      child: Card(
        margin: const EdgeInsets.only(bottom: 16),
        color: backgroundColor,
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20)),
        child: Padding(
          padding: const EdgeInsets.all(14),
          child: Row(
            children: [
              ClipRRect(
                borderRadius: BorderRadius.circular(15),
                child: Image.network(
                  consultation.psycholog.psyImage ?? '',
                  width: 100,
                  height: 100,
                  fit: BoxFit.cover,
                  errorBuilder: (context, error, stackTrace) {
                    return Image.asset(
                      'assets/default_profile.png',
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
                      consultation.psycholog.psyName,
                      style: GoogleFonts.poppins(
                        fontSize: 18,
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                    const Divider(),
                    Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      mainAxisAlignment: MainAxisAlignment.start,
                      children: [
                        Text(
                          consultation.practice.pracName,
                          style: GoogleFonts.poppins(fontSize: 14),
                        ),
                        Text(
                          "${consultation.availableSlot.slotStart} - ${consultation.availableSlot.slotEnd}",
                          style: GoogleFonts.poppins(fontSize: 14),
                        ),
                      ],
                    ),
                    const SizedBox(height: 10),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Container(
                          padding: const EdgeInsets.symmetric(
                            vertical: 4,
                            horizontal: 8,
                          ),
                          decoration: BoxDecoration(
                            color: statusColor,
                            borderRadius: BorderRadius.circular(12),
                          ),
                          child: Text(
                            statusText,
                            style: GoogleFonts.poppins(
                              fontSize: 12,
                              color: Colors.blue.shade800,
                            ),
                          ),
                        ),
                        Text(
                          formattedDate,
                          style: GoogleFonts.poppins(fontSize: 14),
                        ),
                      ],
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
