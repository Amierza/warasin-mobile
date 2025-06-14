import 'package:flutter/material.dart';
import 'package:frontend/controller/motivation/get_all_user_motivation.dart';
import 'package:get/get.dart';
import 'package:frontend/model/motivation.dart';

class HistoryUserMotivationPage extends StatefulWidget {
  const HistoryUserMotivationPage({super.key});

  @override
  State<HistoryUserMotivationPage> createState() => _MotivationHomePageState();
}

class _MotivationHomePageState extends State<HistoryUserMotivationPage> {
  final GetAllUserMotivationController _controller = Get.put(
    GetAllUserMotivationController(),
  );

  String selectedCategory = "Semua";
  List<String> categories = ["Semua"];

  @override
  void initState() {
    super.initState();
    _loadData();
  }

  void _loadData() async {
    await _controller.fetchAllUserMotivation();
    if (_controller.userMotivationList.isNotEmpty) {
      final uniqueCategories =
          _controller.userMotivationList
              .map((m) => m.motivation.motivationCategory.categoryName)
              .toSet();
      setState(() {
        categories.addAll(uniqueCategories);
      });
    }
  }

  List<UserMotivation> get filteredMotivations {
    if (selectedCategory == "Semua") {
      return _controller.userMotivationList;
    }
    return _controller.userMotivationList
        .where(
          (m) =>
              m.motivation.motivationCategory.categoryName == selectedCategory,
        )
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
      backgroundColor: Colors.grey[50],
      appBar: AppBar(
        title: const Text(
          'Motivasi Harian',
          style: TextStyle(fontWeight: FontWeight.bold, color: Colors.white),
        ),
        backgroundColor: Colors.blue[600],
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
              padding: const EdgeInsets.symmetric(vertical: 8),
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
                        style: TextStyle(
                          color: isSelected ? Colors.white : Colors.grey[700],
                          fontSize: 12,
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
                      backgroundColor: Colors.grey[200],
                      selectedColor: Colors.blue[600],
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
                    motivation: motivation.motivation,
                    categoryColor: getCategoryColor(
                      motivation.motivation.motivationCategory.categoryName,
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
            colors: [Colors.white, categoryColor.withOpacity(0.05)],
          ),
        ),
        child: Padding(
          padding: const EdgeInsets.all(20),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
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
                  style: TextStyle(
                    color: categoryColor,
                    fontSize: 12,
                    fontWeight: FontWeight.w600,
                  ),
                ),
              ),
              const SizedBox(height: 16),
              Text(
                motivation.motivationContent,
                style: const TextStyle(
                  fontSize: 16,
                  fontWeight: FontWeight.w500,
                  height: 1.5,
                  color: Colors.black87,
                ),
              ),
              const SizedBox(height: 16),
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
                    style: TextStyle(
                      fontSize: 14,
                      fontWeight: FontWeight.w600,
                      color: Colors.grey[700],
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
