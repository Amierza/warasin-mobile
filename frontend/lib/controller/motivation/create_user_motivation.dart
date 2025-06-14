import 'package:frontend/model/error.dart';
import 'package:frontend/service/motivation_service.dart';
import 'package:get/get.dart';
import 'package:intl/intl.dart';

class CreateUserMotivationController extends GetxController {
  var selectedRating = 0.obs;
  var isLoading = false.obs;

  void setRating(int rating) {
    selectedRating.value = rating;
  }

  Future<void> submitRating(String motivationId) async {
    if (selectedRating.value == 0) {
      Get.snackbar("Peringatan", "Silakan beri rating terlebih dahulu.");
      return;
    }

    isLoading.value = true;

    final date = DateFormat('yyyy-MM-dd').format(DateTime.now());
    final response = await MotivationService.createMotivationService(
      date,
      selectedRating.value,
      motivationId,
    );

    isLoading.value = false;

    if (response is ErrorResponse) {
      Get.snackbar(response.error, response.message);
    }

    if (response != null && response.status == true) {
      Get.snackbar("Berhasil", "Rating berhasil dikirim.");
    } else {
      Get.snackbar("Gagal", "Gagal mengirim rating.");
    }
  }
}
