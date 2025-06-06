import 'package:frontend/model/consultation.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/service/consultation_service.dart';
import 'package:get/get.dart';

class GetAllConsultation extends GetxController {
  var consultationList = <Consultation>[].obs;
  var isLoading = false.obs;
  var allConsultation = Rxn<AllConsultationResponse>();

  Future<void> fetchAllConsultation() async {
    isLoading.value = true;

    try {
      final response = await ConsultationService.getAllConsultation();

      if (response is AllConsultationResponse) {
        allConsultation.value = response;
        consultationList.value = response.data.consultation;
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
