import 'dart:convert';

import 'package:frontend/config/config.dart';
import 'package:frontend/model/chatbot.dart';
import 'package:frontend/model/error.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;

final baseUrl = Config.apiKey;

class ChatbotService {
  static Future<dynamic> createChatbotService(String message) async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.post(
        Uri.parse('$baseUrl/user/chat'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
        body: jsonEncode({"message": message}),
      );

      final responseBody = jsonDecode(response.body);
      print(responseBody);

      if (response.statusCode == 200) {
        final chatbotResponse = ChatbotResponse.fromJson(responseBody);
        return chatbotResponse;
      } else {
        final errorResponse = ErrorResponse.fromJson(responseBody);
        return errorResponse;
      }
    } catch (error) {
      print('Error Fetching create message chatbot : $error');
      return null;
    }
  }
}
