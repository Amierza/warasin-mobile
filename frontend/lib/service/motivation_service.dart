import 'dart:convert';

import 'package:frontend/config/config.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/model/motivation.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;

final baseUrl = Config.apiKey;

class MotivationService {
  static Future<dynamic> getAllMotivationService() async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-all-motivation'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final motivationData = AllMotivationResponse.fromJson(responseBody);
          return motivationData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
      if (response.statusCode == 200) {}
    } catch (error) {
      print('Error Fetching get all motivation : $error');
      return null;
    }
  }

  static Future<dynamic> createMotivationService(
    String date,
    int reaction,
    String motivationid,
  ) async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.post(
        Uri.parse('$baseUrl/user/create-user-motivation'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
        body: jsonEncode({
          "user_mot_display_date": date,
          "user_mot_reaction": reaction,
          "mot_id": motivationid,
        }),
      );

      if (response.statusCode == 400) {
        final errorResponse = ErrorResponse.fromJson(jsonDecode(response.body));
        return errorResponse;
      }

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final userMotivationData = CreateUserMotivationResponse.fromJson(
            responseBody,
          );
          return userMotivationData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
      if (response.statusCode == 200) {}
    } catch (error) {
      print('Error Fetching create user motivation : $error');
      return null;
    }
  }

  static Future<dynamic> getAllUserMotivation() async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-all-user-motivation'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final motivationData = AllUserMotivationResponse.fromJson(
            responseBody,
          );
          return motivationData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
      if (response.statusCode == 200) {}
    } catch (error) {
      print('Error Fetching get all user motivation : $error');
      return null;
    }
  }
}
