import 'package:frontend/model/city.dart';
import 'package:frontend/model/role.dart';

class Psycholog {
  final String psyId;
  final String psyName;
  final String psyStrNumber;
  final String psyEmail;
  final String psyPassword;
  final String psyWorkYear;
  final String psyDescription;
  final String psyPhoneNumber;
  final String? psyImage;
  final City city;
  final Role role;
  final List<Language> languages;
  final List<Specialization> specializations;
  final List<Education> educations;

  Psycholog({
    required this.psyId,
    required this.psyName,
    required this.psyStrNumber,
    required this.psyEmail,
    required this.psyPassword,
    required this.psyWorkYear,
    required this.psyDescription,
    required this.psyPhoneNumber,
    required this.psyImage,
    required this.city,
    required this.role,
    required this.languages,
    required this.specializations,
    required this.educations,
  });

  factory Psycholog.fromJson(Map<String, dynamic> json) {
    return Psycholog(
      psyId: json['psy_id'],
      psyName: json['psy_name'],
      psyStrNumber: json['psy_str_number'],
      psyEmail: json['psy_email'],
      psyPassword: json['psy_password'],
      psyWorkYear: json['psy_work_year'],
      psyDescription: json['psy_description'],
      psyPhoneNumber: json['psy_phone_number'],
      psyImage: json['psy_image'],
      city: City.fromJson(json['city']),
      role: Role.fromJson(json['role']),
      languages:
          (json['language'] as List)
              .map((lang) => Language.fromJson(lang))
              .toList(),
      specializations:
          (json['specialization'] as List)
              .map((spe) => Specialization.fromJson(spe))
              .toList(),
      educations:
          (json['education'] as List)
              .map((edu) => Education.fromJson(edu))
              .toList(),
    );
  }
}

class Language {
  final String? langId;
  final String langName;

  Language({this.langId, required this.langName});

  factory Language.fromJson(Map<String, dynamic> json) {
    return Language(langId: json['lang_id'], langName: json['lang_name']);
  }
}

class Specialization {
  final String? speId;
  final String speName;
  final String speDesc;

  Specialization({this.speId, required this.speName, required this.speDesc});

  factory Specialization.fromJson(Map<String, dynamic> json) {
    return Specialization(
      speId: json['spe_id'],
      speName: json['spe_name'],
      speDesc: json['spe_desc'],
    );
  }
}

class Education {
  final String? eduId;
  final String eduDegree;
  final String eduMajor;
  final String eduInstitution;
  final String eduGraduationYear;

  Education({
    this.eduId,
    required this.eduDegree,
    required this.eduMajor,
    required this.eduInstitution,
    required this.eduGraduationYear,
  });

  factory Education.fromJson(Map<String, dynamic> json) {
    return Education(
      eduId: json['edu_id'],
      eduDegree: json['edu_degree'],
      eduMajor: json['edu_major'],
      eduInstitution: json['edu_institution'],
      eduGraduationYear: json['edu_graduation_year'],
    );
  }
}

class AllPsychologResponse {
  final bool status;
  final String message;
  final String timestamp;
  final List<Psycholog> data;

  AllPsychologResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory AllPsychologResponse.fromJson(Map<String, dynamic> json) {
    return AllPsychologResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data:
          (json['data'] as List).map((psy) => Psycholog.fromJson(psy)).toList(),
    );
  }
}

class PsychologResponse {
  final bool status;
  final String message;
  final String timestamp;
  final Psycholog data;

  PsychologResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory PsychologResponse.fromJson(Map<String, dynamic> json) {
    return PsychologResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data: Psycholog.fromJson(json['data']),
    );
  }
}
