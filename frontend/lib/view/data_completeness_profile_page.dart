import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/logo_widget.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class DataCompletenessProfilePage extends StatefulWidget {
  const DataCompletenessProfilePage({super.key});

  @override
  State<DataCompletenessProfilePage> createState() =>
      _DataCompletenessProfilePageState();
}

class _DataCompletenessProfilePageState
  extends State<DataCompletenessProfilePage> {
  DateTime? selectedDate;
  final TextEditingController dateController = TextEditingController();

  @override
  void dispose() {
    dateController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: Center(
        child: Padding(
          padding: const EdgeInsets.symmetric(vertical: 40, horizontal: 50),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              LogoWarasin(),
              const SizedBox(height: 24),
              Text(
                "Langkah terakhir, lengkapi biodata mu...",
                style: GoogleFonts.poppins(fontSize: 20, fontWeight: semiBold),
              ),
              const SizedBox(height: 24),
              Expanded(
                child: Column(
                  children: [
                    FormFields(
                      hintText: "Gender",
                      icon: Icons.male,
                      keyboardType: TextInputType.name,
                    ),
                    const SizedBox(height: 20),
                    FormFields(
                      hintText: "Tanggal Lahir",
                      icon: Icons.calendar_month_outlined,
                      keyboardType: TextInputType.none,
                      controller: dateController,
                      onTap: () async {
                        final picked = await showDatePicker(
                          context: context,
                          initialDate: selectedDate ?? DateTime.now(),
                          firstDate: DateTime(1945),
                          lastDate: DateTime.now(),
                        );
                        if (picked != null) {
                          setState(() {
                            selectedDate = picked;
                            dateController.text =
                                "${picked.day}/${picked.month}/${picked.year}";
                          });
                        }
                      },
                    ),
                    const SizedBox(height: 20),
                    FormFields(
                      hintText: "Negara",
                      icon: Icons.location_city_sharp,
                      keyboardType: TextInputType.name,
                    ),
                    const SizedBox(height: 20),
                    FormFields(
                      hintText: "Nomer Telepon",
                      icon: Icons.phone,
                      keyboardType: TextInputType.phone,
                    ),
                  ],
                ),
              ),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed: () {
                    Get.toNamed('/home');
                  },
                  style: ElevatedButton.styleFrom(
                    backgroundColor: primaryColor,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(20),
                    ),
                    padding: const EdgeInsets.symmetric(vertical: 12),
                  ),
                  child: Text(
                    "Mulai",
                    style: GoogleFonts.poppins(
                      fontSize: 18,
                      fontWeight: semiBold,
                      color: secondaryTextColor,
                    ),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class FormFields extends StatelessWidget {
  final String hintText;
  final IconData icon;
  final TextInputType keyboardType;
  final VoidCallback? onTap;
  final TextEditingController? controller;

  const FormFields({
    Key? key,
    required this.hintText,
    required this.icon,
    required this.keyboardType,
    this.onTap,
    this.controller,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      controller: controller,
      readOnly: onTap != null,
      onTap: onTap,
      keyboardType: keyboardType,
      decoration: InputDecoration(
        hintText: hintText,
        filled: true,
        fillColor: secondaryTextColor,
        contentPadding: const EdgeInsets.symmetric(
          horizontal: 25,
          vertical: 20,
        ),
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(20),
          borderSide: BorderSide.none,
        ),
        suffixIcon: Icon(icon, size: 30),
        suffixIconColor: Colors.white,
      ),
    );
  }
}
