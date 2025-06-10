import 'package:flutter/material.dart';
import 'package:frontend/controller/consultation/get_detail_consultation.dart';
import 'package:frontend/controller/consultation/update_consultation.dart';
import 'package:get/get.dart';
import 'package:intl/intl.dart';

class UpdateConsultationPage extends StatefulWidget {
  const UpdateConsultationPage({super.key});

  @override
  State<UpdateConsultationPage> createState() => _UpdateConsultationPageState();
}

class _UpdateConsultationPageState extends State<UpdateConsultationPage> {
  final _formKey = GlobalKey<FormState>();
  final updateController = Get.put(UpdateConsultationController());
  final detailController = Get.put(GetDetailConsultationController());

  final TextEditingController dateController = TextEditingController();
  final TextEditingController commentController = TextEditingController();

  final RxnInt selectedRate = RxnInt();
  final RxnInt selectedStatus = RxnInt();

  DateTime? selectedDate;
  String dateForStorage = '';

  late final String consultationId;

  @override
  void initState() {
    super.initState();
    consultationId = Get.parameters['id'] ?? '';
    _loadConsultationData();
  }

  Future<void> _loadConsultationData() async {
    await detailController.fetchConsultationDetail(consultationId);
    if (detailController.consultation.value != null) {
      final data = detailController.consultation.value!.data;
      final date = DateTime.parse(data.consulDate);
      selectedDate = date;
      dateForStorage = DateFormat('yyyy-MM-dd').format(date);
      dateController.text = _formatDateIndonesian(date);

      // Set initial values for rate and status
      selectedRate.value =
          _isValidRate(data.consulRate) ? data.consulRate : null;
      selectedStatus.value =
          _isValidStatus(data.consulStatus) ? data.consulStatus : null;

      commentController.text = data.consulComment ?? '';

      // Load practice and slot data
      await updateController.fetchInitialData(data.psycholog.psyId);

      // Set initial values for practice and slot
      updateController.setInitialValues(
        practiceId: data.practice.pracId,
        slotId: data.availableSlot.slotId,
      );
    }
  }

  bool _isValidRate(int? rate) {
    return rate != null && rate >= 1 && rate <= 5;
  }

  bool _isValidStatus(int? status) {
    return status != null && (status == 0 || status == 1);
  }

  Future<void> _selectDate(BuildContext context) async {
    final DateTime? picked = await showDatePicker(
      context: context,
      initialDate: selectedDate ?? DateTime.now(),
      firstDate: DateTime(2020),
      lastDate: DateTime(2030),
      builder: (context, child) {
        return Theme(
          data: Theme.of(context).copyWith(
            colorScheme: ColorScheme.light(
              primary: Colors.blue[600]!,
              onPrimary: Colors.white,
              surface: Colors.white,
              onSurface: Colors.black,
            ),
          ),
          child: child!,
        );
      },
    );
    if (picked != null && picked != selectedDate) {
      setState(() {
        selectedDate = picked;
        dateForStorage = DateFormat('yyyy-MM-dd').format(picked);
        dateController.text = _formatDateIndonesian(picked);
      });
    }
  }

  String _getIndonesianDayName(DateTime date) {
    const days = [
      'Senin',
      'Selasa',
      'Rabu',
      'Kamis',
      'Jumat',
      'Sabtu',
      'Minggu',
    ];
    return days[date.weekday - 1];
  }

  String _getIndonesianMonthName(DateTime date) {
    const months = [
      'Januari',
      'Februari',
      'Maret',
      'April',
      'Mei',
      'Juni',
      'Juli',
      'Agustus',
      'September',
      'Oktober',
      'November',
      'Desember',
    ];
    return months[date.month - 1];
  }

  String _formatDateIndonesian(DateTime date) {
    return '${_getIndonesianDayName(date)}, ${date.day} ${_getIndonesianMonthName(date)} ${date.year}';
  }

  Future<void> _submitUpdate() async {
    if (_formKey.currentState!.validate()) {
      final updateBody = {
        "date": dateForStorage,
        "rate": selectedRate.value,
        "status": selectedStatus.value,
        "comment": commentController.text,
        "practice_id": updateController.selectedPracticeId.value,
        "slot_id": updateController.selectedSlotId.value,
      };
      updateBody.removeWhere((key, value) => value == null || value == '');

      await updateController.updateConsultation(
        consulId: consultationId,
        body: updateBody,
      );

      if (updateController.updateSuccess.value) {
        Get.back(result: true);
        Get.snackbar(
          'Berhasil',
          'Konsultasi berhasil diperbarui',
          backgroundColor: Colors.green,
          colorText: Colors.white,
        );
      }
    }
  }

  Widget _buildTextFormField({
    required TextEditingController controller,
    required String label,
    required String hint,
    required IconData icon,
    int maxLines = 1,
    bool readOnly = false,
    String? Function(String?)? validator,
    Widget? suffixIcon,
  }) {
    return TextFormField(
      controller: controller,
      readOnly: readOnly,
      validator: validator,
      maxLines: maxLines,
      decoration: InputDecoration(
        labelText: label,
        hintText: hint,
        prefixIcon: Icon(icon),
        suffixIcon: suffixIcon,
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(10)),
      ),
    );
  }

  Widget _buildSectionHeader(String title, IconData icon) {
    return Row(
      children: [
        Icon(icon, size: 20, color: Colors.blue),
        const SizedBox(width: 8),
        Text(
          title,
          style: const TextStyle(fontSize: 16, fontWeight: FontWeight.w600),
        ),
      ],
    );
  }

  Widget _buildDropdownField<T>({
    required T? value,
    required String label,
    required IconData icon,
    required List<DropdownMenuItem<T>> items,
    required void Function(T?) onChanged,
    String? Function(T?)? validator,
    String? hint,
    bool isRequired =
        false, // Tambahkan parameter untuk menentukan apakah wajib diisi
  }) {
    return DropdownButtonFormField<T>(
      value: value,
      onChanged: onChanged,
      items: items,
      validator:
          isRequired
              ? (value) => value == null ? '$label wajib dipilih' : null
              : null,
      decoration: InputDecoration(
        labelText: label,
        hintText: hint,
        prefixIcon: Icon(icon),
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(10)),
      ),
      isExpanded: true,
    );
  }

  Widget _buildLoadingIndicator(String message) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const CircularProgressIndicator(strokeWidth: 3),
          const SizedBox(height: 16),
          Text(
            message,
            style: const TextStyle(color: Colors.grey, fontSize: 16),
          ),
        ],
      ),
    );
  }

  Widget _buildErrorView(String message) {
    return Center(
      child: Padding(
        padding: const EdgeInsets.all(20),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Icon(Icons.error_outline, size: 48, color: Colors.red),
            const SizedBox(height: 16),
            Text(
              message,
              style: const TextStyle(color: Colors.red, fontSize: 16),
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: _loadConsultationData,
              child: const Text('Coba Lagi'),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildFormContent() {
    return SingleChildScrollView(
      child: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Form(
          key: _formKey,
          child: Column(
            children: [
              // Header
              Container(
                padding: const EdgeInsets.all(20),
                decoration: BoxDecoration(
                  color: Colors.blue[50],
                  borderRadius: BorderRadius.circular(12),
                  border: Border.all(color: Colors.blue[100]!),
                ),
                child: Column(
                  children: [
                    Icon(
                      Icons.medical_services_outlined,
                      size: 48,
                      color: Colors.blue[600],
                    ),
                    const SizedBox(height: 12),
                    Text(
                      'Form Update Konsultasi',
                      style: TextStyle(
                        fontSize: 18,
                        fontWeight: FontWeight.w600,
                        color: Colors.blue[800],
                      ),
                    ),
                    const SizedBox(height: 4),
                    Text(
                      'Perbarui informasi konsultasi Anda',
                      style: TextStyle(fontSize: 14, color: Colors.blue[600]),
                    ),
                  ],
                ),
              ),

              const SizedBox(height: 24),

              // Form
              Container(
                padding: const EdgeInsets.all(20),
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(12),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.grey.withOpacity(0.1),
                      blurRadius: 10,
                    ),
                  ],
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    _buildSectionHeader('Informasi Waktu', Icons.schedule),
                    const SizedBox(height: 12),
                    GestureDetector(
                      onTap: () => _selectDate(context),
                      child: AbsorbPointer(
                        child: _buildTextFormField(
                          controller: dateController,
                          label: 'Tanggal Konsultasi',
                          hint: 'Pilih tanggal konsultasi',
                          icon: Icons.calendar_today,
                          validator:
                              (value) =>
                                  value!.isEmpty ? 'Tanggal wajib diisi' : null,
                          suffixIcon: Icon(
                            Icons.arrow_drop_down,
                            color: Colors.grey[600],
                          ),
                        ),
                      ),
                    ),
                    const SizedBox(height: 24),

                    _buildSectionHeader('Penilaian', Icons.star),
                    const SizedBox(height: 12),
                    Obx(() {
                      return _buildDropdownField<int>(
                        value: selectedRate.value,
                        label: 'Rating',
                        icon: Icons.star,
                        hint: 'Pilih rating konsultasi',
                        validator:
                            (value) =>
                                value == null ? 'Rating wajib dipilih' : null,
                        items: List.generate(5, (index) {
                          int val = index + 1;
                          return DropdownMenuItem<int>(
                            value: val,
                            child: Row(
                              mainAxisSize: MainAxisSize.min,
                              children: [
                                Row(
                                  mainAxisSize: MainAxisSize.min,
                                  children: List.generate(
                                    val,
                                    (_) => const Icon(
                                      Icons.star,
                                      color: Colors.amber,
                                      size: 16,
                                    ),
                                  ),
                                ),
                                const SizedBox(width: 8),
                                Text('$val Bintang'),
                              ],
                            ),
                          );
                        }),
                        onChanged: (val) {
                          selectedRate.value = val;
                        },
                      );
                    }),

                    const SizedBox(height: 24),

                    _buildSectionHeader('Status Konsultasi', Icons.assignment),
                    const SizedBox(height: 12),
                    // Status konsultasi: 0 = UpComing, 1 = Done (tidak ditampilkan untuk user), 2 = Cancel
                    Obx(() {
                      return _buildDropdownField<int>(
                        value: selectedStatus.value,
                        label: 'Status',
                        icon: Icons.info,
                        hint: 'Pilih status konsultasi',
                        validator:
                            (value) =>
                                value == null ? 'Status wajib dipilih' : null,
                        items: const [
                          DropdownMenuItem<int>(
                            value: 0, // UpComing
                            child: Row(
                              children: [
                                Icon(
                                  Icons.pending,
                                  color: Colors.orange,
                                  size: 16,
                                ),
                                SizedBox(width: 8),
                                Text("Belum Selesai"),
                              ],
                            ),
                          ),
                          DropdownMenuItem<int>(
                            value: 1,
                            child: Row(
                              children: [
                                Icon(Icons.cancel, color: Colors.red, size: 16),
                                SizedBox(width: 8),
                                Text("Batal Pesanan"),
                              ],
                            ),
                          ),
                        ],
                        onChanged: (val) {
                          selectedStatus.value = val;
                        },
                      );
                    }),

                    const SizedBox(height: 24),

                    _buildSectionHeader('Komentar', Icons.comment),
                    const SizedBox(height: 12),
                    _buildTextFormField(
                      controller: commentController,
                      label: 'Komentar',
                      hint: 'Masukkan komentar konsultasi',
                      icon: Icons.note,
                      maxLines: 4,
                    ),
                    const SizedBox(height: 24),

                    _buildSectionHeader('Identifikasi', Icons.badge),
                    const SizedBox(height: 12),
                    Obx(() {
                      if (updateController.isLoading.value) {
                        return const CircularProgressIndicator();
                      }
                      return _buildDropdownField<String>(
                        value:
                            updateController.selectedPracticeId.value.isEmpty
                                ? null
                                : updateController.selectedPracticeId.value,
                        label: 'Metode Konsultasi',
                        icon: Icons.medical_services,
                        hint: 'Pilih metode konsultasi',
                        validator:
                            (value) =>
                                value == null ? 'Metode wajib dipilih' : null,
                        items:
                            updateController.practiceList.map((practice) {
                              return DropdownMenuItem<String>(
                                value: practice.pracId,
                                child: Text(
                                  '${practice.pracType} - ${practice.pracName}',
                                ),
                              );
                            }).toList(),
                        onChanged: (value) {
                          if (value != null) {
                            updateController.selectedPracticeId.value = value;
                          }
                        },
                      );
                    }),
                    const SizedBox(height: 16),
                    Obx(() {
                      if (updateController.isLoading.value) {
                        return const CircularProgressIndicator();
                      }
                      return _buildDropdownField<String>(
                        value:
                            updateController.selectedSlotId.value.isEmpty
                                ? null
                                : updateController.selectedSlotId.value,
                        label: 'Slot Waktu',
                        icon: Icons.access_time,
                        hint: 'Pilih slot waktu',
                        validator:
                            (value) =>
                                value == null ? 'Slot wajib dipilih' : null,
                        items:
                            updateController.slotList.map((slot) {
                              return DropdownMenuItem<String>(
                                value: slot.slotId,
                                child: Text(
                                  '${slot.slotStart} - ${slot.slotEnd}',
                                ),
                              );
                            }).toList(),
                        onChanged: (value) {
                          if (value != null) {
                            updateController.selectedSlotId.value = value;
                          }
                        },
                      );
                    }),
                  ],
                ),
              ),

              const SizedBox(height: 32),
              SizedBox(
                width: double.infinity,
                height: 56,
                child: Obx(() {
                  return ElevatedButton(
                    onPressed:
                        updateController.isLoading.value ? null : _submitUpdate,
                    style: ElevatedButton.styleFrom(
                      backgroundColor: Colors.blue[600],
                      foregroundColor: Colors.white,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12),
                      ),
                    ),
                    child:
                        updateController.isLoading.value
                            ? const Row(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                SizedBox(
                                  width: 20,
                                  height: 20,
                                  child: CircularProgressIndicator(
                                    strokeWidth: 2,
                                    valueColor: AlwaysStoppedAnimation<Color>(
                                      Colors.white,
                                    ),
                                  ),
                                ),
                                SizedBox(width: 12),
                                Text("Menyimpan..."),
                              ],
                            )
                            : const Text(
                              "Simpan Perubahan",
                              style: TextStyle(
                                fontSize: 16,
                                fontWeight: FontWeight.w600,
                              ),
                            ),
                  );
                }),
              ),
            ],
          ),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[50],
      appBar: AppBar(
        title: const Text(
          "Update Konsultasi",
          style: TextStyle(fontWeight: FontWeight.w600, fontSize: 20),
        ),
        backgroundColor: Colors.white,
        foregroundColor: Colors.black87,
        elevation: 0,
        centerTitle: true,
        bottom: PreferredSize(
          preferredSize: const Size.fromHeight(1),
          child: Container(height: 1, color: Colors.grey[200]),
        ),
      ),
      body: Obx(() {
        if (detailController.isLoading.value) {
          return _buildLoadingIndicator('Memuat data konsultasi...');
        }
        if (detailController.errorMessage.value.isNotEmpty) {
          return _buildErrorView(detailController.errorMessage.value);
        }
        if (detailController.consultation.value == null) {
          return _buildErrorView('Data konsultasi tidak ditemukan');
        }
        return _buildFormContent();
      }),
    );
  }

  @override
  void dispose() {
    dateController.dispose();
    commentController.dispose();
    super.dispose();
  }
}
