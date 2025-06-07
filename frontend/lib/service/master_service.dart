import 'dart:convert';

import 'package:frontend/config/config.dart';
import 'package:frontend/model/city.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/model/province.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;

final baseUrl = Config.apiKey;

class MasterService {
  static Future<dynamic> getAllProvince() async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/get-all-province'),
        headers: {'Content-Type': 'application/json'},
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final provinceData = ProvinceResponse.fromJson(responseBody);
          return provinceData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get all province : $error');
      return null;
    }
  }

  static Future<dynamic> getAllCity(String provinceId) async {
    try {
      final response = await http.get(
        Uri.parse('$baseUrl/get-all-city?province_id=$provinceId'),
        headers: {'Content-Type': 'application/json'},
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final cityData = CityResponse.fromJson(responseBody);
          return cityData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get all cities : $error');
      return null;
    }
  }
}
