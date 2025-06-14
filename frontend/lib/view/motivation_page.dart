import 'package:flutter/material.dart';
import 'package:frontend/controller/motivation/get_all_motivation.dart';
import 'package:frontend/model/motivation.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class MotivationPage extends StatefulWidget {
  const MotivationPage({super.key});

  @override
  State<MotivationPage> createState() => _MotivationHomePageState();
}

class _MotivationHomePageState extends State<MotivationPage> {
  final GetAllMotivationController _controller = Get.put(
    GetAllMotivationController(),
  );
  String selectedCategory = "Semua";
  List<String> categories = ["Semua"];

  @override
  void initState() {
    super.initState();
    _controller.fetchAllMotivatio().then((_) {
      // Extract unique categories after data is loaded
      if (_controller.motivationList.isNotEmpty) {
        Set<String> uniqueCategories =
            _controller.motivationList
                .map((m) => m.motivationCategory.categoryName)
                .toSet();
        setState(() {
          categories.addAll(uniqueCategories.toList());
        });
      }
    });
  }

  List<Motivation> get filteredMotivations {
    if (selectedCategory == "Semua") {
      return _controller.motivationList;
    }
    return _controller.motivationList
        .where((m) => m.motivationCategory.categoryName == selectedCategory)
        .toList();
  }

  Color getCategoryColor(String categoryName) {
    switch (categoryName) {
      case "Strategi Menghadapi Tantangan":
        return Colors.orange;
      case "Cinta Diri":
        return Colors.pink;
      case "Mengatasi Kecemasan":
        return Colors.teal;
      case "Penyembuhan Emosional":
        return Colors.purple;
      case "Ketahanan Mental":
        return Colors.indigo;
      case "Pemikiran Positif":
        return Colors.amber;
      case "Rasa Syukur":
        return Colors.green;
      case "Pengurangan Stres":
        return Colors.blue;
      case "Kewaspadaan Penuh":
        return Colors.deepPurple;
      case "Perawatan Diri":
        return Colors.red;
      default:
        return Colors.grey;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      appBar: AppBar(
        title: Text(
          'Motivasi Harian',
          style: GoogleFonts.poppins(
            fontSize: 20,
            fontWeight: FontWeight.bold,
            color: Colors.white,
          ),
        ),
        backgroundColor: primaryColor,
        elevation: 0,
        centerTitle: true,
      ),
      body: Obx(() {
        if (_controller.isLoading.value) {
          return const Center(child: CircularProgressIndicator());
        }

        return Column(
          children: [
            // Category Filter
            Container(
              height: 60,
              padding: const EdgeInsets.symmetric(vertical: 8, horizontal: 12),
              child: ListView.builder(
                scrollDirection: Axis.horizontal,
                itemCount: categories.length,
                itemBuilder: (context, index) {
                  final category = categories[index];
                  final isSelected = category == selectedCategory;
                  return Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 4),
                    child: FilterChip(
                      label: Text(
                        category,
                        style: GoogleFonts.poppins(
                          color: isSelected ? Colors.white : tertiaryTextColor,
                          fontSize: 14,
                          fontWeight:
                              isSelected ? FontWeight.bold : FontWeight.normal,
                        ),
                      ),
                      selected: isSelected,
                      onSelected: (selected) {
                        setState(() {
                          selectedCategory = category;
                        });
                      },
                      backgroundColor: Colors.grey[100],
                      selectedColor: primaryColor,
                      checkmarkColor: Colors.white,
                    ),
                  );
                },
              ),
            ),
            // Motivations List
            Expanded(
              child: ListView.builder(
                padding: const EdgeInsets.all(16),
                itemCount: filteredMotivations.length,
                itemBuilder: (context, index) {
                  final motivation = filteredMotivations[index];
                  return MotivationCard(
                    motivation: motivation,
                    categoryColor: getCategoryColor(
                      motivation.motivationCategory.categoryName,
                    ),
                  );
                },
              ),
            ),
          ],
        );
      }),
    );
  }
}

class MotivationCard extends StatelessWidget {
  final Motivation motivation;
  final Color categoryColor;

  const MotivationCard({
    super.key,
    required this.motivation,
    required this.categoryColor,
  });

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: const EdgeInsets.only(bottom: 16),
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      elevation: 4,
      shadowColor: Colors.black26,
      child: Container(
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(16),
          gradient: LinearGradient(
            begin: Alignment.topLeft,
            end: Alignment.bottomRight,
            colors: [
              categoryColor.withOpacity(0.2),
              categoryColor.withOpacity(0.7),
            ],
          ),
        ),
        child: Padding(
          padding: const EdgeInsets.all(20),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Category Badge
              Container(
                padding: const EdgeInsets.symmetric(
                  horizontal: 12,
                  vertical: 6,
                ),
                decoration: BoxDecoration(
                  color: categoryColor.withOpacity(0.1),
                  borderRadius: BorderRadius.circular(20),
                  border: Border.all(
                    color: categoryColor.withOpacity(0.3),
                    width: 1,
                  ),
                ),
                child: Text(
                  motivation.motivationCategory.categoryName,
                  style: GoogleFonts.poppins(
                    color: categoryColor,
                    fontSize: 12,
                    fontWeight: semiBold,
                  ),
                ),
              ),
              const SizedBox(height: 16),
              // Quote Content
              Text(
                motivation.motivationContent,
                style: GoogleFonts.poppins(
                  fontSize: 14,
                  fontWeight: medium,
                  height: 1.5,
                  color: primaryTextColor,
                ),
              ),
              const SizedBox(height: 16),
              // Author
              Row(
                children: [
                  Container(
                    width: 4,
                    height: 20,
                    decoration: BoxDecoration(
                      color: categoryColor,
                      borderRadius: BorderRadius.circular(2),
                    ),
                  ),
                  const SizedBox(width: 12),
                  Text(
                    "— ${motivation.motivationAuthor}",
                    style: GoogleFonts.poppins(
                      fontSize: 14,
                      fontWeight: FontWeight.w600,
                      color: primaryTextColor,
                      fontStyle: FontStyle.italic,
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
