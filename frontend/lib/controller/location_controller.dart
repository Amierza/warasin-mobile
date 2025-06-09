import 'package:get/get.dart';
import 'package:frontend/model/city.dart';
import 'package:frontend/model/province.dart';
import 'package:frontend/service/master_service.dart';

class LocationController extends GetxController {
  var provinces = <Province>[].obs;
  var cities = <City>[].obs;

  var selectedProvinceId = ''.obs;
  var selectedCityId = ''.obs;

  var isLoadingProvince = false.obs;
  var isLoadingCity = false.obs;

  @override
  void onInit() {
    fetchAllProvinces();
    super.onInit();
  }

  Future<void> fetchAllProvinces() async {
    isLoadingProvince.value = true;
    final result = await MasterService.getAllProvince();
    if (result != null && result is ProvinceResponse) {
      provinces.value = result.data;
    }
    isLoadingProvince.value = false;
  }

  Future<void> fetchCitiesByProvince(String provinceId) async {
    isLoadingCity.value = true;
    selectedProvinceId.value = provinceId;
    final result = await MasterService.getAllCity(provinceId);
    if (result != null && result is CityResponse) {
      cities.value = result.data;
    }
    isLoadingCity.value = false;
  }

  void selectProvince(String provinceId) {
    selectedProvinceId.value = provinceId;
    selectedCityId.value = '';
    fetchCitiesByProvince(provinceId);
  }

  void selectCity(String cityId) {
    selectedCityId.value = cityId;
  }
}
