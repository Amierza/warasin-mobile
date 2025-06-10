import 'dart:convert';

import 'package:frontend/config/config.dart';
import 'package:frontend/model/consultation.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/model/psycholog.dart';
import 'package:get_storage/get_storage.dart';
import 'package:http/http.dart' as http;

final baseUrl = Config.apiKey;

class ConsultationService {
  static Future<dynamic> getAllPsycholog() async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-all-psycholog'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final psychologData = AllPsychologResponse.fromJson(responseBody);
          return psychologData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get all psycholog : $error');
      return null;
    }
  }

  static Future<dynamic> getDetailPsycholog(String psyId) async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-detail-psycholog/$psyId'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final psychologData = PsychologResponse.fromJson(responseBody);
          return psychologData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get detail psycholog : $error');
      return null;
    }
  }

  static Future<dynamic> createConsultation(
    String consulDate,
    String slotId,
    String pracId,
  ) async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.post(
        Uri.parse('$baseUrl/user/create-consultation'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
        body: jsonEncode({
          "consul_date": consulDate,
          "slot_id": slotId,
          "prac_id": pracId,
        }),
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final consultationData = ConsultationResponse.fromJson(responseBody);
          return consultationData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching create consultation : $error');
      return null;
    }
  }

  static Future<dynamic> getAllPractice(String psyId) async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-all-practice/$psyId'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final practiceData = PracticeResponse.fromJson(responseBody);
          return practiceData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get all psycholog : $error');
      return null;
    }
  }

  static Future<dynamic> getAllAvailableSlot(String psyId) async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-all-available-slot/$psyId'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final slotData = AvailableSlotResponse.fromJson(responseBody);
          return slotData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get all slot : $error');
      return null;
    }
  }

  static Future<dynamic> getAllConsultation() async {
    final box = GetStorage();
    final token = box.read('access_token');

    if (token == null) return null;

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-all-consultation'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final allConsultationData = AllConsultationResponse.fromJson(
            responseBody,
          );
          return allConsultationData;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get all consultation : $error');
      return null;
    }
  }

  static Future<dynamic> updateConsultationService(
    String consulId,
    Map<String, dynamic> data,
  ) async {
    final box = GetStorage();
    final token = box.read('access_token');

    try {
      final response = await http.patch(
        Uri.parse('$baseUrl/user/update-consultation/$consulId'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
        body: jsonEncode(data),
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final consultationResponse = GetDetailConsultationResponse.fromJson(
            responseBody,
          );
          print('consultation response : $consultationResponse');
          return responseBody;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching update consultation : $error');
      return null;
    }
  }

  static Future<dynamic> getDetailConsultationService(String consulId) async {
    final box = GetStorage();
    final token = box.read('access_token');

    try {
      final response = await http.get(
        Uri.parse('$baseUrl/user/get-detail-consultation/$consulId'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final responseBody = jsonDecode(response.body);
        if (responseBody['status'] == true) {
          final consultationResponse = GetDetailConsultationResponse.fromJson(
            responseBody,
          );
          return consultationResponse;
        } else {
          final errorResponse = ErrorResponse.fromJson(responseBody);
          return errorResponse;
        }
      }
    } catch (error) {
      print('Error Fetching get detail consultation : $error');
      return null;
    }
  }
}
