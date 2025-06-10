import 'package:frontend/model/consultation.dart';
import 'package:frontend/service/consultation_service.dart';
import 'package:get/get.dart';

class GetDetailConsultationController extends GetxController {
  var isLoading = false.obs;
  var consultation = Rxn<GetDetailConsultationResponse>();
  var errorMessage = ''.obs;

  Future<void> fetchConsultationDetail(String consulId) async {
    try {
      isLoading.value = true;
      errorMessage.value = '';

      final result = await ConsultationService.getDetailConsultationService(
        consulId,
      );

      if (result is GetDetailConsultationResponse) {
        consultation.value = result;
      } else {
        errorMessage.value = 'Failed to fetch consultation detail.';
      }
    } catch (e) {
      errorMessage.value = 'An error occurred: $e';
    } finally {
      isLoading.value = false;
    }
  }
}
