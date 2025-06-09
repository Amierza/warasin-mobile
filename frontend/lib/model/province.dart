class Province {
  final String? provinceId;
  final String provinceName;

  Province({this.provinceId, required this.provinceName});

  factory Province.fromJson(Map<String, dynamic> json) {
    return Province(
      provinceId: json['province_id'],
      provinceName: json['province_name'],
    );
  }

  @override
  String toString() {
    return 'Province(provinceId: $provinceId, provinceName: $provinceName)';
  }
}

class ProvinceResponse {
  final bool status;
  final String message;
  final String timestamp;
  final List<Province> data;

  ProvinceResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory ProvinceResponse.fromJson(Map<String, dynamic> json) {
    return ProvinceResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data:
          (json['data'] as List)
              .map((province) => Province.fromJson(province))
              .toList(),
    );
  }
}
