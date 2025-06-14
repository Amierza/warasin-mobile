import 'package:flutter/material.dart';
import 'package:frontend/controller/motivation/get_all_motivation.dart';
import 'package:frontend/controller/motivation/create_user_motivation.dart';
import 'package:frontend/controller/motivation/get_all_user_motivation.dart';
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
  final CreateUserMotivationController _ratingController = Get.put(
    CreateUserMotivationController(),
  );
  final GetAllUserMotivationController _userMotivationController = Get.put(
    GetAllUserMotivationController(),
  );
  String selectedCategory = "Semua";
  List<String> categories = ["Semua"];
  List<String> ratedMotivationIds = [];
  Map<String, int> motivationRatings = {};

  @override
  void initState() {
    super.initState();
    _controller.fetchAllMotivatio().then((_) {
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

    _userMotivationController.fetchAllUserMotivation().then((_) {
      setState(() {
        ratedMotivationIds =
            _userMotivationController.userMotivationList
                .map((e) => e.motivation.motivationId)
                .toList();

        // Store ratings for each motivation
        for (var userMotivation
            in _userMotivationController.userMotivationList) {
          motivationRatings[userMotivation.motivation.motivationId] =
              userMotivation.userMotivationReaction;
        }
      });
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
                    onRatingSubmitted: (rating) {
                      _ratingController
                          .submitRating(motivation.motivationId)
                          .then((_) {
                            // After successful rating, update the local state
                            setState(() {
                              ratedMotivationIds.add(motivation.motivationId);
                              motivationRatings[motivation.motivationId] =
                                  rating;
                            });
                          });
                    },
                    isAlreadyRated: ratedMotivationIds.contains(
                      motivation.motivationId,
                    ),
                    userRating: motivationRatings[motivation.motivationId] ?? 0,
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

class MotivationCard extends StatefulWidget {
  final Motivation motivation;
  final Color categoryColor;
  final Function(int) onRatingSubmitted;
  final bool isAlreadyRated;
  final int userRating;

  const MotivationCard({
    super.key,
    required this.motivation,
    required this.categoryColor,
    required this.onRatingSubmitted,
    required this.isAlreadyRated,
    required this.userRating,
  });

  @override
  State<MotivationCard> createState() => _MotivationCardState();
}

class _MotivationCardState extends State<MotivationCard> {
  int selectedRating = 0;
  bool showRatingBar = false;
  final CreateUserMotivationController _ratingController = Get.find();

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
              widget.categoryColor.withOpacity(0.2),
              widget.categoryColor.withOpacity(0.7),
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
                  color: widget.categoryColor.withOpacity(0.1),
                  borderRadius: BorderRadius.circular(20),
                  border: Border.all(
                    color: widget.categoryColor.withOpacity(0.3),
                    width: 1,
                  ),
                ),
                child: Text(
                  widget.motivation.motivationCategory.categoryName,
                  style: GoogleFonts.poppins(
                    color: widget.categoryColor,
                    fontSize: 12,
                    fontWeight: semiBold,
                  ),
                ),
              ),
              const SizedBox(height: 16),
              // Quote Content
              Text(
                widget.motivation.motivationContent,
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
                      color: widget.categoryColor,
                      borderRadius: BorderRadius.circular(2),
                    ),
                  ),
                  const SizedBox(width: 12),
                  Text(
                    "— ${widget.motivation.motivationAuthor}",
                    style: GoogleFonts.poppins(
                      fontSize: 14,
                      fontWeight: FontWeight.w600,
                      color: primaryTextColor,
                      fontStyle: FontStyle.italic,
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 16),
              // Rating Section
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  if (widget.isAlreadyRated)
                    Column(
                      crossAxisAlignment: CrossAxisAlignment.start,
                      children: [
                        Text(
                          'Rating Anda:',
                          style: GoogleFonts.poppins(
                            fontSize: 12,
                            color: primaryTextColor,
                          ),
                        ),
                        const SizedBox(height: 4),
                        Row(
                          children: List.generate(5, (index) {
                            return Icon(
                              index < widget.userRating
                                  ? Icons.star
                                  : Icons.star_border,
                              color:
                                  index < widget.userRating
                                      ? Colors.amber
                                      : Colors.grey,
                              size: 24,
                            );
                          }),
                        ),
                      ],
                    )
                  else if (!showRatingBar)
                    Align(
                      alignment: Alignment.centerRight,
                      child: TextButton(
                        onPressed: () {
                          setState(() {
                            showRatingBar = true;
                          });
                        },
                        child: Text(
                          'Beri Rating',
                          style: GoogleFonts.poppins(
                            color: widget.categoryColor,
                            fontWeight: semiBold,
                          ),
                        ),
                      ),
                    ),
                  if (!widget.isAlreadyRated && showRatingBar) ...[
                    Text(
                      'Berikan rating untuk motivasi ini:',
                      style: GoogleFonts.poppins(
                        fontSize: 12,
                        color: primaryTextColor,
                      ),
                    ),
                    const SizedBox(height: 8),
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: List.generate(5, (index) {
                        return GestureDetector(
                          onTap: () {
                            setState(() {
                              selectedRating = index + 1;
                              _ratingController.setRating(selectedRating);
                            });
                          },
                          child: Icon(
                            index < selectedRating
                                ? Icons.star
                                : Icons.star_border,
                            color:
                                index < selectedRating
                                    ? Colors.amber
                                    : Colors.grey,
                            size: 30,
                          ),
                        );
                      }),
                    ),
                    const SizedBox(height: 8),
                    Obx(() {
                      if (_ratingController.isLoading.value) {
                        return const Center(child: CircularProgressIndicator());
                      }
                      return Row(
                        mainAxisAlignment: MainAxisAlignment.end,
                        children: [
                          TextButton(
                            onPressed: () {
                              setState(() {
                                showRatingBar = false;
                                selectedRating = 0;
                              });
                            },
                            child: Text(
                              'Batal',
                              style: GoogleFonts.poppins(color: Colors.grey),
                            ),
                          ),
                          const SizedBox(width: 8),
                          ElevatedButton(
                            onPressed:
                                selectedRating > 0
                                    ? () {
                                      widget.onRatingSubmitted(selectedRating);
                                      setState(() {
                                        showRatingBar = false;
                                        selectedRating = 0;
                                      });
                                    }
                                    : null,
                            style: ElevatedButton.styleFrom(
                              backgroundColor: widget.categoryColor,
                              shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(8),
                              ),
                            ),
                            child: Text(
                              'Kirim',
                              style: GoogleFonts.poppins(color: Colors.white),
                            ),
                          ),
                        ],
                      );
                    }),
                  ],
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
