import 'package:flutter/material.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/logo_widget.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

List<Map<String, dynamic>> form = [
  {"hintText": "Gender", "icon": Icons.male},
  {"hintText": "Tanggal Lahir", "icon": Icons.calendar_month_outlined},
  {"hintText": "Negara", "icon": Icons.location_city_sharp},
  {"hintText": "Nomer Telepon", "icon": Icons.phone},
];

class DataCompletenessProfilePage extends StatelessWidget {
  const DataCompletenessProfilePage({super.key});

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
                child: ListView.separated(
                  itemCount: form.length,
                  separatorBuilder:
                      (context, index) => const SizedBox(height: 20),
                  itemBuilder: (context, index) {
                    return GestureDetector(
                      child: FormFields(
                        hintText: form[index]['hintText']!,
                        icon: form[index]['icon']!,
                      ),
                    );
                  },
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
                    padding: const EdgeInsets.symmetric(vertical: 16),
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

  const FormFields({Key? key, required this.hintText, required this.icon})
    : super(key: key);

  @override
  Widget build(BuildContext context) {
    return TextFormField(
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
