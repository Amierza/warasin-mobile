import 'package:flutter/material.dart';
import 'package:frontend/model/consultation.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/service/consultation_service.dart';
import 'package:get/get.dart';

class UpdateConsultationController extends GetxController {
  var isLoading = false.obs;
  var updateSuccess = false.obs;
  var errorMessage = ''.obs;
  var selectedPracticeId = ''.obs;
  var selectedSlotId = ''.obs;

  RxList<Practice> practiceList = <Practice>[].obs;
  RxList<AvailableSlot> slotList = <AvailableSlot>[].obs;

  late String psychologistId;

  @override
  void onInit() {
    super.onInit();
    // psychologistId perlu di-set dari luar sebelum fetch data
  }

  Future<void> fetchInitialData(String psyId) async {
    psychologistId = psyId;
    isLoading.value = true;
    await Future.wait([
      fetchPractices(psychologistId),
      fetchSlots(psychologistId),
    ]);
    isLoading.value = false;
  }

  Future<void> fetchPractices(String psyId) async {
    try {
      final result = await ConsultationService.getAllPractice(psyId);

      if (result is PracticeResponse) {
        practiceList.value = result.data;
      } else if (result is ErrorResponse) {
        errorMessage.value = result.message;
        Get.snackbar("Gagal", result.message);
      } else {
        errorMessage.value = "Gagal memuat data praktik";
        Get.snackbar("Gagal", errorMessage.value);
      }
    } catch (e) {
      errorMessage.value = "Terjadi kesalahan saat memuat praktik: $e";
      Get.snackbar("Error", errorMessage.value);
    }
  }

  Future<void> fetchSlots(String psyId) async {
    try {
      final result = await ConsultationService.getAllAvailableSlot(psyId);

      if (result is AvailableSlotResponse) {
        slotList.value = result.data;
      } else if (result is ErrorResponse) {
        errorMessage.value = result.message;
        Get.snackbar("Gagal", result.message);
      } else {
        errorMessage.value = "Gagal memuat data slot";
        Get.snackbar("Gagal", errorMessage.value);
      }
    } catch (e) {
      errorMessage.value = "Terjadi kesalahan saat memuat slot: $e";
      Get.snackbar("Error", errorMessage.value);
    }
  }

  void setInitialValues({required String practiceId, required String slotId}) {
    selectedPracticeId.value = practiceId;
    selectedSlotId.value = slotId;
  }

  String get selectedPracticeLabel {
    final prac = practiceList.firstWhereOrNull(
      (e) => e.pracId == selectedPracticeId.value,
    );
    return prac != null ? '${prac.pracType} - ${prac.pracName}' : '';
  }

  String get selectedSlotLabel {
    final slot = slotList.firstWhereOrNull(
      (e) => e.slotId == selectedSlotId.value,
    );
    return slot != null ? '${slot.slotStart} - ${slot.slotEnd}' : '';
  }

  Future<void> updateConsultation({
    required String consulId,
    required Map<String, dynamic> body,
  }) async {
    try {
      isLoading.value = true;
      errorMessage.value = '';
      updateSuccess.value = false;

      final result = await ConsultationService.updateConsultationService(
        consulId,
        body,
      );

      if (result == null) {
        errorMessage.value = "Tidak dapat terhubung ke server";
        Get.snackbar("Error", errorMessage.value);
      } else if (result is ConsultationResponse) {
        updateSuccess.value = true;
        Get.back(result: true); // kembali ke halaman sebelumnya
        Get.snackbar(
          "Sukses",
          "Konsultasi berhasil diperbarui",
          snackPosition: SnackPosition.BOTTOM,
          backgroundColor: Colors.green,
          colorText: Colors.white,
        );
      } else if (result is ErrorResponse) {
        errorMessage.value = result.message;
        Get.snackbar(result.error, result.message);
      }
    } catch (e) {
      errorMessage.value = "Terjadi kesalahan: ${e.toString()}";
      Get.snackbar("Error", errorMessage.value);
    } finally {
      isLoading.value = false;
    }
  }
}
