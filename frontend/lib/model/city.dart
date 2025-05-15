import 'package:frontend/model/province.dart';

class City {
  final String? cityId;
  final String cityName;
  final String cityType;
  final Province province;

  City({
    this.cityId,
    required this.cityName,
    required this.cityType,
    required this.province,
  });

  factory City.fromJson(Map<String, dynamic> json) {
    return City(
      cityId: json['city_id'],
      cityName: json['city_name'],
      cityType: json['city_type'],
      province: Province.fromJson(json['province']),
    );
  }
}