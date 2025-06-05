import 'dart:convert';
import 'package:http/http.dart' as http;
import '../base_service.dart' show BaseService;

class IncomeSourceService extends BaseService {
  Future<Map<String, dynamic>?> createIncomeSource(
    Map<String, dynamic> data,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.post(
        url('income-sources/create'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      final body = jsonDecode(response.body);
      return response.statusCode == 201
          ? body['income_source']
          : {
            'error': body['message'] ?? 'Failed to create income source',
            'statusCode': response.statusCode,
          };
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<List<Map<String, dynamic>>> getAllIncomeSources() async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('income-sources/'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['income_sources']);
      } else {
        throw Exception('Failed to fetch income sources');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> getIncomeSourceByID(int sourceID) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('income-sources/$sourceID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['income_source'];
      } else {
        return {
          'error': 'Failed to fetch income source',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<List<Map<String, dynamic>>> getIncomeSourcesByUserID(
    int userID,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.get(
        url('income-sources/user/$userID'),
        headers: headersWithToken(token),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return List<Map<String, dynamic>>.from(body['income_sources']);
      } else {
        throw Exception('Failed to fetch income sources for user');
      }
    } catch (e) {
      return [];
    }
  }

  Future<Map<String, dynamic>?> updateIncomeSource(
    int sourceID,
    Map<String, dynamic> data,
  ) async {
    try {
      final token = await getAuthToken();
      final response = await http.put(
        url('income-sources/$sourceID'),
        headers: headersWithToken(token),
        body: jsonEncode(data),
      );

      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        return body['income_source'];
      } else {
        return {
          'error': 'Failed to update income source',
          'statusCode': response.statusCode,
        };
      }
    } catch (e) {
      return {'error': e.toString(), 'statusCode': 500};
    }
  }

  Future<bool> deleteIncomeSource(int sourceID) async {
    try {
      final token = await getAuthToken();
      final response = await http.delete(
        url('income-sources/$sourceID'),
        headers: headersWithToken(token),
      );

      return response.statusCode == 200;
    } catch (e) {
      return false;
    }
  }
}
