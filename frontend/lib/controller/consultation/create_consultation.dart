import 'package:frontend/model/consultation.dart';
import 'package:frontend/model/error.dart';
import 'package:frontend/service/consultation_service.dart';
import 'package:get/get.dart';

class ConsultationController extends GetxController {
  var consulDate = ''.obs;
  var slotId = ''.obs;
  var pracId = ''.obs;

  var isLoading = false.obs;
  var consultationResult = Rxn<ConsultationResponse>();
  var errorResult = Rxn<ErrorResponse>();

  Future<void> createConsultation() async {
    if (consulDate.value.isEmpty ||
        slotId.value.isEmpty ||
        pracId.value.isEmpty) {
      return;
    }

    try {
      isLoading.value = true;
      final result = await ConsultationService.createConsultation(
        consulDate.value,
        slotId.value,
        pracId.value,
      );

      if (result is ConsultationResponse) {
        consultationResult.value = result;
        errorResult.value = null;
      } 
      if (result is ErrorResponse) {
        errorResult.value = result;
      } 
    } catch(error) {
      print('Unexpected error or null result, $error');
    }
    isLoading.value = false;
  }
}
