import 'package:flutter/material.dart';
import 'package:frontend/controller/news/create_news_detail_controller.dart';
import 'package:frontend/controller/news/get_detail_news_controller.dart';
import 'package:frontend/shared/theme.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:intl/intl.dart';

class NewsDetailPage extends StatelessWidget {
  const NewsDetailPage({super.key});

  String formatDate(String isoDate) {
    try {
      final dateTime = DateTime.parse(isoDate);
      return DateFormat('dd-MM-yyyy').format(dateTime);
    } catch (_) {
      return isoDate;
    }
  }

  @override
  Widget build(BuildContext context) {
    final controller = Get.put(GetDetailNewsController());
    final newsController = Get.put(CreateNewsDetailController());
    final newsId = Get.parameters['id'];

    if (controller.detailNews.value == null && newsId != null) {
      controller.fetchDetailNews(newsId);
    }

    return Scaffold(
      backgroundColor: backgroundColor,
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
          "Detail Berita",
          style: GoogleFonts.poppins(color: primaryTextColor, fontWeight: bold),
        ),
        centerTitle: true,
        backgroundColor: Colors.transparent,
        elevation: 0,
      ),
      body: Obx(() {
        if (controller.isLoading.value) {
          return const Center(child: CircularProgressIndicator());
        }

        final news = controller.detailNews.value;

        if (news == null) {
          return Center(
            child: Text(
              "Berita tidak ditemukan.",
              style: GoogleFonts.poppins(
                color: primaryTextColor,
                fontSize: 18,
                fontWeight: bold,
              ),
            ),
          );
        }

        return SingleChildScrollView(
          padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 10),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              ClipRRect(
                borderRadius: BorderRadius.circular(20),
                child:
                    news.data.newsImage.isNotEmpty
                        ? Image.network(
                          news.data.newsImage,
                          width: double.infinity,
                          height: 200,
                          fit: BoxFit.cover,
                          errorBuilder:
                              (context, error, stackTrace) => Image.asset(
                                "assets/default_image.jpg",
                                width: double.infinity,
                                height: 200,
                                fit: BoxFit.cover,
                              ),
                        )
                        : Image.asset(
                          "assets/default_image.jpg",
                          width: double.infinity,
                          height: 200,
                          fit: BoxFit.cover,
                        ),
              ),
              const SizedBox(height: 16),
              Text(
                news.data.newsTitle,
                style: GoogleFonts.poppins(
                  color: primaryTextColor,
                  fontSize: 18,
                  fontWeight: bold,
                ),
              ),
              const SizedBox(height: 8),
              Row(
                children: [
                  const CircleAvatar(
                    backgroundImage: AssetImage("assets/default_profile.png"),
                    radius: 30,
                    backgroundColor: Colors.transparent,
                  ),
                  const SizedBox(width: 8),
                  Text(
                    "By Admin Warasin",
                    style: GoogleFonts.poppins(
                      fontSize: 13,
                      fontWeight: light,
                      color: primaryTextColor,
                    ),
                  ),
                  const Spacer(),
                  Text(
                    formatDate(news.data.newsDate),
                    style: GoogleFonts.poppins(
                      fontSize: 11,
                      fontWeight: light,
                      color: primaryTextColor,
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 16),
              Text(
                news.data.newsBody,
                textAlign: TextAlign.justify,
                style: GoogleFonts.poppins(
                  fontSize: 16,
                  height: 1.5,
                  color: primaryTextColor,
                ),
              ),
              const SizedBox(height: 26),
              Obx(() {
                return SizedBox(
                  width: double.infinity,
                  height: 56,
                  child: ElevatedButton(
                    onPressed:
                        newsController.isLoading.value || newsId == null
                            ? null
                            : () =>
                                newsController.fetchCreateNewsDetail(newsId),
                    style: ElevatedButton.styleFrom(
                      backgroundColor: primaryColor,
                      foregroundColor: Colors.white,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(16),
                      ),
                      elevation: 2,
                      shadowColor: primaryColor.withOpacity(0.3),
                    ),
                    child:
                        newsController.isLoading.value
                            ? const SizedBox(
                              width: 24,
                              height: 24,
                              child: CircularProgressIndicator(
                                color: Colors.white,
                                strokeWidth: 2,
                              ),
                            )
                            : Text(
                              "Selesaikan membaca",
                              style: GoogleFonts.poppins(
                                fontSize: 16,
                                fontWeight: FontWeight.w600,
                              ),
                            ),
                  ),
                );
              }),
            ],
          ),
        );
      }),
    );
  }
}
