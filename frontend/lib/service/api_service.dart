import 'dart:convert';
import 'package:frontend/config/config.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/model/token_payload.dart';
import 'package:frontend/model/user_detail_response.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;
import 'package:jwt_decoder/jwt_decoder.dart';

final baseUrl = Config.apiKey;

class ApiService {
  TokenPayload? getUserFromToken(String token) {
    try {
      final decoded = JwtDecoder.decode(token);
      return TokenPayload.fromJson(decoded);
    } catch (e) {
      print("Failed to decode token: $e");
      return null;
    }
  }

  static Future<http.Response> register(
    String name,
    String email,
    String password,
  ) async {
    final response = await http.post(
      Uri.parse('$baseUrl/user/register'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'name': name, 'email': email, 'password': password}),
    );
    return response;
  }

  static Future<http.Response> login(String email, String password) async {
    final response = await http.post(
      Uri.parse('$baseUrl/user/login'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'email': email, 'password': password}),
    );
    return response;
  }

  static Future<http.Response> checkEmail(String email) async {
    final response = await http.post(
      Uri.parse('${baseUrl}/user/send-forgot-password-email'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'email': email}),
    );
    return response;
  }

  static Future<dynamic> getUserDetailService() async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-detail-user'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      final responseBody = jsonDecode(response.body);
      if (response.statusCode == 200) {
        if (responseBody['status'] == true) {
          return UserDetailResponse.fromJson(responseBody);
        } else {
          box.remove('access_token');
          return ErrorResponse.fromJson(responseBody);
        }
      } else if (response.statusCode == 401) {
        box.remove('access_token');
        return ErrorResponse.fromJson(responseBody);
      } else {
        return null;
      }
    } catch (e) {
      print("Error fetching user detail: $e");
      return null;
    }
  }
}
