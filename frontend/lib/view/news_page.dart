import 'package:flutter/material.dart';
import 'package:frontend/controller/news/get_all_news_controller.dart';
import 'package:frontend/shared/theme.dart';
import 'package:frontend/widget/header.dart';
import 'package:frontend/widget/navigation_bar.dart';
import 'package:frontend/widget/news_card.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';

class NewsPage extends StatelessWidget {
  const NewsPage({super.key});

  @override
  Widget build(BuildContext context) {
    final controller = Get.put(GetAllNewsController());

    return Scaffold(
      extendBody: true,
      backgroundColor: primaryColor,
      body: Obx(() {
        if (controller.isLoading.value) {
          return Center(child: CircularProgressIndicator());
        }
        return SingleChildScrollView(
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
                        hintText: "Temukan Berita Menarik...",
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
                  child: Column(
                    children: [
                      Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text(
                            "Daftar Berita",
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
                      Column(
                        children:
                            controller.newsList.map((news) {
                              return NewsCard(
                                image: news.newsImage,
                                title: news.newsTitle,
                                desc: news.newsBody,
                                date: news.newsDate,
                              );
                            }).toList(),
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ),
        );
      }),
      bottomNavigationBar: const CustomNavigationBar(currentIndex: 3),
    );
  }
}
