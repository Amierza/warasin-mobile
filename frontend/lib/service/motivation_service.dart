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
          final psychologData = AllMotivationResponse.fromJson(responseBody);
          return psychologData;
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
}
