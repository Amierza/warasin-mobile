import 'package:frontend/model/city.dart';
import 'package:frontend/model/role.dart';

class UserDetailResponse {
  final bool status;
  final String message;
  final User data;
  final String timestamp;

  UserDetailResponse({
    required this.status,
    required this.message,
    required this.data,
    required this.timestamp,
  });

  factory UserDetailResponse.fromJson(Map<String, dynamic> json) {
    return UserDetailResponse(
      status: json['status'],
      message: json['message'],
      data: User.fromJson(json['data']),
      timestamp: json['timestamp'],
    );
  }
}

class User {
  final String userId;
  final String userName;
  final String userEmail;
  final String userPassword;
  final String userImage;
  final bool? userGender;
  final String userBirthDate;
  final String userPhoneNumber;
  final int userData01;
  final int userData02;
  final int userData03;
  final bool? isVerified;
  final City city;
  final Role role;

  User({
    required this.userId,
    required this.userName,
    required this.userEmail,
    required this.userPassword,
    required this.userImage,
    required this.userGender,
    required this.userBirthDate,
    required this.userPhoneNumber,
    required this.userData01,
    required this.userData02,
    required this.userData03,
    required this.isVerified,
    required this.city,
    required this.role,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      userId: json['user_id'],
      userName: json['user_name'],
      userEmail: json['user_email'],
      userPassword: json['user_password'],
      userImage: json['user_image'],
      userGender: json['user_gender'],
      userBirthDate: json['user_birth_date'],
      userPhoneNumber: json['user_phone_number'],
      userData01: json['user_data01'],
      userData02: json['user_data02'],
      userData03: json['user_data03'],
      isVerified: json['is_verified'],
      city: City.fromJson(json['city']),
      role: Role.fromJson(json['role']),
    );
  }
}
