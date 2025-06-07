import 'package:frontend/model/province.dart';

class City {
  String? cityId;
  final String cityName;
  final String cityType;
  final Province? province;

  City({
    this.cityId,
    required this.cityName,
    required this.cityType,
    this.province,
  });

  factory City.fromJson(Map<String, dynamic> json) {
    return City(
      cityId: json['city_id'],
      cityName: json['city_name'],
      cityType: json['city_type'],
      province:
          json['province'] != null ? Province.fromJson(json['province']) : null,
    );
  }
}

class CityResponse {
  final bool status;
  final String message;
  final String timestamp;
  final List<City> data;

  CityResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory CityResponse.fromJson(Map<String, dynamic> json) {
    return CityResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data: (json['data'] as List).map((city) => City.fromJson(city)).toList(),
    );
  }
}
