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

  // Tambahkan ini:
  @override
  bool operator ==(Object other) {
    return identical(this, other) ||
        other is Province &&
            runtimeType == other.runtimeType &&
            provinceId == other.provinceId;
  }

  @override
  int get hashCode => provinceId.hashCode;

  @override
  String toString() => provinceName;
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
