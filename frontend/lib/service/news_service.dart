import 'dart:convert';
import 'package:frontend/config/config.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/model/news.dart';
import 'package:frontend/model/news_detail.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;

final baseUrl = Config.apiKey;

class NewsService {
  static Future<dynamic> getAllNews() async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-all-news'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final newsData = GetAllNewsResponse.fromJson(responseBody);
          return newsData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get all news : $error');
      return null;
    }
  }

  static Future<dynamic> getDetailNews(String newsId) async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-detail-news/$newsId'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final newsData = GetDetailNewsResponse.fromJson(responseBody);
          return newsData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get detail news : $error');
      return null;
    }
  }

  static Future<dynamic> createNewsDetailService(
    String date,
    String newsId,
  ) async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.post(
        Uri.parse('$baseUrl/user/create-news-detail'),
        headers: {
          "Authorization": "Bearer $token",
          "Content-Type": "application/json",
        },
        body: jsonEncode({"news_detail_date": date, "news_id": newsId}),
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final newsDetailResponse = CreateNewsDetail.fromJson(responseBody);
          return newsDetailResponse;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching create news detail : $error');
      return null;
    }
  }
}
