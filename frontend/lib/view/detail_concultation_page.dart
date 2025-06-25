import 'package:flutter/material.dart';
import 'package:frontend/controller/consultation/create_consultation.dart';
import 'package:frontend/controller/consultation/get_all_available_slot.dart';
import 'package:frontend/controller/consultation/get_all_practice.dart';
import 'package:frontend/controller/consultation/get_detail_psycholog.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/dialog_auth.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:intl/intl.dart';

class DetailConcultationPage extends StatefulWidget {
  const DetailConcultationPage({super.key});

  @override
  State<DetailConcultationPage> createState() => _DetailConcultationPageState();
}

class _DetailConcultationPageState extends State<DetailConcultationPage> {
  int selectedDateIndex = 0;
  int? selectedConsultationIndex;
  String? selectedConsulId;
  int? selectedSlotIndex;
  String? selectedSlotId;

  List<DateTime> availableDates = [];

  @override
  void initState() {
    super.initState();
    generateDates();
  }

  void generateDates() {
    final now = DateTime.now();
    availableDates = List.generate(
      14,
      (index) => now.add(Duration(days: index)),
    );
  }

  String _getMonthName(int month) {
    const months = [
      'Januari',
      'Februari',
      'Maret',
      'April',
      'Mei',
      'Juni',
      'Juli',
      'Agustus',
      'September',
      'Oktober',
      'November',
      'Desember',
    ];
    return months[month - 1];
  }

  String getSelectedDateFormatted() {
    final selectedDate = availableDates[selectedDateIndex];
    return DateFormat('yyyy-MM-dd').format(selectedDate);
  }

  @override
  Widget build(BuildContext context) {
    final consultationController = Get.put(ConsultationController());
    final controller = Get.put(GetDetailPsycholog());
    final practiceController = Get.put(GetAllPractice());
    final slotController = Get.put(GetAllAvailableSlot());
    final psyId = Get.parameters['id'];

    if (controller.detailPsycholog.value == null && psyId != null) {
      controller.fetchDetailPsycholog(psyId);
    }

    if (practiceController.allPractice.value == null && psyId != null) {
      practiceController.fetchAllPractice(psyId);
    }

    if (slotController.allSlot.value == null && psyId != null) {
      slotController.fetchAllSlot(psyId);
    }

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
      body: Obx(() {
        if (controller.isLoading.value) {
          return Center(child: CircularProgressIndicator());
        }

        final psycholog = controller.detailPsycholog.value;
        final practice = practiceController.allPractice.value;
        final availableSlot = slotController.allSlot.value;

        if (psycholog == null) {
          return Center(
            child: Text(
              "Psycholog tidak ditemukan.",
              style: GoogleFonts.poppins(
                color: primaryTextColor,
                fontSize: 18,
                fontWeight: bold,
              ),
            ),
          );
        }

        if (practice == null) {
          return Center(
            child: Text(
              "Practice tidak ditemukan.",
              style: GoogleFonts.poppins(
                color: primaryTextColor,
                fontSize: 18,
                fontWeight: bold,
              ),
            ),
          );
        }

        return SingleChildScrollView(
          padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 20),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                children: [
                  ClipRRect(
                    borderRadius: BorderRadius.circular(10),
                    child:
                        (psycholog.data.psyImage != null &&
                                psycholog.data.psyImage!.startsWith('http'))
                            ? Image.network(
                              psycholog.data.psyImage!,
                              width: 100,
                              height: 100,
                              fit: BoxFit.cover,
                            )
                            : Image.asset(
                              'assets/default_profile.png',
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
                          psycholog.data.psyName,
                          style: GoogleFonts.poppins(
                            fontSize: 18,
                            fontWeight: bold,
                            color: primaryTextColor,
                          ),
                        ),
                        Column(
                          children:
                              practice.data
                                  .map(
                                    (prac) => Text(
                                      '${prac.pracName}\n${prac.pracType}',
                                      style: GoogleFonts.poppins(
                                        fontSize: 14,
                                        color: primaryTextColor,
                                      ),
                                    ),
                                  )
                                  .toList(),
                        ),
                        SizedBox(height: 8),
                      ],
                    ),
                  ),
                ],
              ),
              SizedBox(height: 40),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  _infoCard(
                    Icons.people,
                    psycholog.data.educations.length.toString(),
                    "Gelar",
                  ),
                  _infoCard(
                    Icons.bar_chart_outlined,
                    psycholog.data.psyWorkYear,
                    "Tahun Kerja",
                  ),
                  _infoCard(
                    Icons.star_outlined,
                    psycholog.data.specializations.length.toString(),
                    "Spesialisasi",
                  ),
                  _infoCard(
                    Icons.language,
                    psycholog.data.languages.length.toString(),
                    "Bahasa",
                  ),
                ],
              ),
              SizedBox(height: 40),
              Text(
                "Tentang Konselor",
                style: GoogleFonts.poppins(fontSize: 22, fontWeight: bold),
              ),
              Text(
                psycholog.data.psyDescription,
                style: GoogleFonts.poppins(
                  fontSize: 14,
                  color: primaryTextColor,
                ),
              ),
              SizedBox(height: 40),
              Text(
                "Mulai Konsultasi",
                style: GoogleFonts.poppins(fontSize: 22, fontWeight: bold),
              ),
              const SizedBox(height: 10),
              practice.data.isNotEmpty
                  ? Wrap(
                    spacing: 12,
                    runSpacing: 12,
                    children: List.generate(practice.data.length, (index) {
                      final consul = practice.data[index];
                      final isSelected = selectedConsultationIndex == index;

                      return GestureDetector(
                        onTap: () {
                          setState(() {
                            selectedConsultationIndex = index;
                            selectedConsulId =
                                consul
                                    .pracId; // ini yang benar untuk ID praktik
                          });
                        },
                        child: _consultationMethod(
                          consul.pracType,
                          isSelected,
                          consul.pracName,
                          consul.pracAddress,
                        ),
                      );
                    }),
                  )
                  : Text(
                    "Konsultasi tidak tersedia.",
                    style: GoogleFonts.poppins(
                      fontSize: 14,
                      color: primaryTextColor,
                    ),
                  ),
              const SizedBox(height: 10),
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
                scrollDirection: Axis.horizontal,
                child: Row(
                  children: List.generate(availableDates.length, (index) {
                    final date = availableDates[index];
                    final day = date.day.toString();
                    final month = _getMonthName(date.month);
                    return Padding(
                      padding: EdgeInsets.only(right: 16),
                      child: GestureDetector(
                        onTap: () {
                          setState(() {
                            selectedDateIndex = index;
                          });
                        },
                        child: _dateCard(
                          day,
                          month,
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
              availableSlot != null && availableSlot.data.isNotEmpty
                  ? Wrap(
                    spacing: 12,
                    runSpacing: 12,
                    children: List.generate(availableSlot.data.length, (index) {
                      final slot = availableSlot.data[index];
                      final isDisabled = slot.isBooked;

                      return GestureDetector(
                        onTap:
                            isDisabled
                                ? null
                                : () {
                                  setState(() {
                                    selectedSlotIndex = index;
                                    selectedSlotId = slot.slotId;
                                  });
                                },
                        child: Opacity(
                          opacity: isDisabled ? 0.4 : 1.0,
                          child: _timeCard(
                            "${slot.slotStart} - ${slot.slotEnd}",
                            selectedSlotIndex == index && !isDisabled,
                          ),
                        ),
                      );
                    }),
                  )
                  : Text(
                    "Tidak ada slot tersedia.",
                    style: GoogleFonts.poppins(
                      fontSize: 14,
                      color: primaryTextColor,
                    ),
                  ),

              const SizedBox(height: 50),
              Obx(
                () => ElevatedButton(
                  onPressed: () async {
                    var dateFormatted = getSelectedDateFormatted();
                    var selectedSlot = selectedSlotId;
                    var selectedPractice = selectedConsulId;

                    if (selectedSlot?.isNotEmpty == true &&
                        selectedPractice?.isNotEmpty == true) {
                      consultationController.consulDate.value = dateFormatted;
                      consultationController.slotId.value = selectedSlot!;
                      consultationController.pracId.value = selectedPractice!;

                      await consultationController.createConsultation();

                      if (consultationController.errorResult.value != null) {
                        showCustomDialog(
                          context: context,
                          icon: Icons.cancel,
                          title: 'Gagal Konsultasi',
                          message: 'Konsultasi gagal dipesan.',
                          onPressed: () {
                            Get.back();
                          },
                          backgroundColor: dangerColor,
                        );
                      }

                      if (consultationController.errorResult.value == null) {
                        showCustomDialog(
                          context: context,
                          icon: Icons.check_circle,
                          title: 'Pesan Konsultasi Berhasil',
                          message: 'Konsultasi berhasil dipesan.',
                          onPressed: () {
                            Get.back();
                            if (psyId != null) {
                              slotController.fetchAllSlot(psyId);
                              setState(() {
                                selectedSlotIndex = null;
                                selectedSlotId = null;
                              });
                            }
                          },
                          backgroundColor: successColor,
                        );
                      }
                    } else {
                      Get.snackbar(
                        "Gagal",
                        "Lengkapi semua pilihan terlebih dahulu",
                        snackPosition: SnackPosition.BOTTOM,
                      );
                    }
                  },
                  style: ElevatedButton.styleFrom(
                    backgroundColor: primaryColor,
                    minimumSize: Size(double.infinity, 50),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                  ),
                  child:
                      consultationController.isLoading.value
                          ? CircularProgressIndicator(color: Colors.white)
                          : Text(
                            "Pesan Konselor Sekarang",
                            style: GoogleFonts.poppins(
                              fontSize: 16,
                              color: Colors.white,
                              fontWeight: semiBold,
                            ),
                          ),
                ),
              ),
            ],
          ),
        );
      }),
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
    padding: const EdgeInsets.all(20),
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

Widget _consultationMethod(
  String type,
  bool isSelected,
  String name,
  String address,
) {
  return Container(
    padding: EdgeInsets.symmetric(vertical: 12, horizontal: 16),
    width: double.infinity,
    decoration: BoxDecoration(
      color: isSelected ? primaryColor : Colors.grey[200],
      borderRadius: BorderRadius.circular(10),
    ),
    child: Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(
          '$name - $type',
          style: GoogleFonts.poppins(
            fontSize: 14,
            fontWeight: semiBold,
            color: isSelected ? Colors.white : primaryTextColor,
          ),
        ),
        SizedBox(height: 10),
        Text(
          address,
          style: GoogleFonts.poppins(
            fontSize: 12,
            fontWeight: medium,
            color: isSelected ? Colors.white : primaryTextColor,
          ),
        ),
      ],
    ),
  );
}

Widget _timeCard(String time, bool isSelected, {bool isDisabled = false}) {
  return Container(
    padding: EdgeInsets.all(10),
    decoration: BoxDecoration(
      color:
          isDisabled
              ? Colors.grey[300]
              : isSelected
              ? primaryColor
              : Colors.grey[200],
      borderRadius: BorderRadius.circular(10),
      border: isDisabled ? Border.all(color: Colors.grey) : null,
    ),
    child: Text(
      time,
      style: GoogleFonts.poppins(
        fontSize: 14,
        fontWeight: semiBold,
        color:
            isDisabled
                ? Colors.grey
                : isSelected
                ? Colors.white
                : primaryTextColor,
      ),
    ),
  );
}
