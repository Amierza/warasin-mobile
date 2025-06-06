import 'package:frontend/model/psycholog.dart';
import 'package:frontend/model/user_detail_response.dart';

class AvailableSlot {
  final String slotId;
  final String slotStart;
  final String slotEnd;
  final bool isBooked;

  AvailableSlot({
    required this.slotId,
    required this.slotStart,
    required this.slotEnd,
    required this.isBooked,
  });

  factory AvailableSlot.fromJson(Map<String, dynamic> json) {
    return AvailableSlot(
      slotId: json['slot_id'],
      slotStart: json['slot_start'],
      slotEnd: json['slot_end'],
      isBooked: json['slot_is_booked'],
    );
  }
}

class Schedule {
  final String schedId;
  final String schedDay;
  final String schedOpen;
  final String schedClose;

  Schedule({
    required this.schedId,
    required this.schedDay,
    required this.schedOpen,
    required this.schedClose,
  });

  factory Schedule.fromJson(Map<String, dynamic> json) {
    return Schedule(
      schedId: json['prac_sched_id'],
      schedDay: json['prac_sched_day'],
      schedOpen: json['prac_sched_open'],
      schedClose: json['prac_sched_close'],
    );
  }
}

class Practice {
  final String pracId;
  final String pracType;
  final String pracName;
  final String pracAddress;
  final String pracPhoneNumber;
  final List<Schedule>? pracSchedule;

  Practice({
    required this.pracId,
    required this.pracType,
    required this.pracName,
    required this.pracAddress,
    required this.pracPhoneNumber,
    this.pracSchedule,
  });

  factory Practice.fromJson(Map<String, dynamic> json) {
    return Practice(
      pracId: json['prac_id'],
      pracType: json['prac_type'],
      pracName: json['prac_name'],
      pracAddress: json['prac_address'],
      pracPhoneNumber: json['prac_phone_number'],
      pracSchedule:
          json['practice_schedule'] != null
              ? (json['practice_schedule'] as List)
                  .map((item) => Schedule.fromJson(item))
                  .toList()
              : [],
    );
  }
}

class AvailableSlotResponse {
  final bool status;
  final String message;
  final String timestamp;
  final List<AvailableSlot> data;

  AvailableSlotResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory AvailableSlotResponse.fromJson(Map<String, dynamic> json) {
    return AvailableSlotResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data:
          (json['data'] as List)
              .map((slot) => AvailableSlot.fromJson(slot))
              .toList(),
    );
  }
}

class PracticeResponse {
  final bool status;
  final String message;
  final String timestamp;
  final List<Practice> data;

  PracticeResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory PracticeResponse.fromJson(Map<String, dynamic> json) {
    return PracticeResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data:
          (json['data'] as List)
              .map((prac) => Practice.fromJson(prac))
              .toList(),
    );
  }
}

class Consultation {
  final String consulId;
  final String consulDate;
  int? consulRate;
  String? consulComment;
  final int consulStatus;
  final Psycholog psycholog;
  final AvailableSlot availableSlot;
  final Practice practice;

  Consultation({
    required this.consulId,
    required this.consulDate,
    this.consulRate,
    this.consulComment,
    required this.consulStatus,
    required this.psycholog,
    required this.availableSlot,
    required this.practice,
  });

  factory Consultation.fromJson(Map<String, dynamic> json) {
    return Consultation(
      consulId: json['consul_id'],
      consulDate: json['consul_date'],
      consulRate: json['consul_rate'],
      consulComment: json['consul_comment'],
      consulStatus: json['consul_status'],
      psycholog: Psycholog.fromJson(json['psycholog']),
      availableSlot: AvailableSlot.fromJson(json['available_slot']),
      practice: Practice.fromJson(json['practice']),
    );
  }
}

class AllConsultation {
  final User user;
  final List<Consultation> consultation;

  AllConsultation({required this.user, required this.consultation});

  factory AllConsultation.fromJson(Map<String, dynamic> json) {
    return AllConsultation(
      user: User.fromJson(json['user']),
      consultation:
          json['consultation'] != null
              ? (json['consultation'] as List)
                  .map((consul) => Consultation.fromJson(consul))
                  .toList()
              : [],
    );
  }
}

class AllConsultationResponse {
  final bool status;
  final String message;
  final String timestamp;
  final AllConsultation data;

  AllConsultationResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory AllConsultationResponse.fromJson(Map<String, dynamic> json) {
    return AllConsultationResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data: AllConsultation.fromJson(json['data']),
    );
  }
}

class DetailConsultationResponse {
  final bool status;
  final String message;
  final String timestamp;
  final Consultation data;

  DetailConsultationResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory DetailConsultationResponse.fromJson(Map<String, dynamic> json) {
    return DetailConsultationResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data: Consultation.fromJson(json['data']),
    );
  }
}

class ConsultationResponse {
  final bool status;
  final String message;
  final String timestamp;
  final ConsultationData data;

  ConsultationResponse({
    required this.status,
    required this.message,
    required this.timestamp,
    required this.data,
  });

  factory ConsultationResponse.fromJson(Map<String, dynamic> json) {
    return ConsultationResponse(
      status: json['status'],
      message: json['message'],
      timestamp: json['timestamp'],
      data: ConsultationData.fromJson(json['data']),
    );
  }
}

class ConsultationData {
  final String consulId;
  final String consulDate;
  final int? consulRate;
  final String? consulComment;
  final int consulStatus;
  final User user;
  final AvailableSlot availableSlot;
  final Practice practice;

  ConsultationData({
    required this.consulId,
    required this.consulDate,
    this.consulRate,
    this.consulComment,
    required this.consulStatus,
    required this.user,
    required this.availableSlot,
    required this.practice,
  });

  factory ConsultationData.fromJson(Map<String, dynamic> json) {
    return ConsultationData(
      consulId: json['consul_id'],
      consulDate: json['consul_date'],
      consulRate: json['consul_rate'],
      consulComment: json['consul_comment'],
      consulStatus: json['consul_status'],
      user: User.fromJson(json['user']),
      availableSlot: AvailableSlot.fromJson(json['available_slot']),
      practice: Practice.fromJson(json['practice']),
    );
  }
}
