import 'package:frontend/model/consultation.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/service/consultation_service.dart';
import 'package:get/get.dart';

class GetAllAvailableSlot extends GetxController {
  var slotList = <AvailableSlot>[].obs;
  var isLoading = false.obs;
  var allSlot = Rxn<AvailableSlotResponse>();

  Future<void> fetchAllSlot(String psyId) async {
    isLoading.value = true;

    try {
      final response = await ConsultationService.getAllAvailableSlot(psyId);
      if (response is AvailableSlotResponse) {
        allSlot.value = response;
        slotList.value = response.data;
      }

      if (response is ErrorResponse) {
        Get.snackbar(response.error, response.message);
      }
    } catch (error) {
      Get.snackbar('Something went wrong', error.toString());
    } finally {
      isLoading.value = false;
    }
  }
}
